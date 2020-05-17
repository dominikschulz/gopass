package ondisk

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/blang/semver"
	proto "github.com/golang/protobuf/proto"
	"github.com/gopasspw/gopass/pkg/backend/storage/kv/ondisk/gpb"
	"github.com/gopasspw/gopass/pkg/ctxutil"
	"github.com/gopasspw/gopass/pkg/out"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	idxFile = "index.pb"
)

type OnDisk struct {
	dir string
	idx *gpb.Store
}

func New(baseDir string) (*OnDisk, error) {
	idx, err := loadOrCreate(baseDir)
	if err != nil {
		return nil, err
	}
	return &OnDisk{
		dir: baseDir,
		idx: idx,
	}, nil
}

func loadOrCreate(path string) (*gpb.Store, error) {
	path = filepath.Join(path, idxFile)
	buf, err := ioutil.ReadFile(path)
	if os.IsNotExist(err) {
		return &gpb.Store{
			Name:    filepath.Base(path),
			Entries: make(map[string]*gpb.Entry),
		}, nil
	}
	idx := &gpb.Store{}
	err = proto.Unmarshal(buf, idx)
	return idx, err
}

func (o *OnDisk) saveIndex() error {
	buf, err := proto.Marshal(o.idx)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(o.dir, idxFile), buf, 0600)
}

func (o *OnDisk) Get(ctx context.Context, name string) ([]byte, error) {
	em := o.idx.GetEntries()
	if em == nil {
		return nil, fmt.Errorf("not found")
	}
	e, found := em[name]
	if !found {
		return nil, fmt.Errorf("not found")
	}
	r := e.Latest()
	if r == nil {
		return nil, fmt.Errorf("not found")
	}
	path := filepath.Join(o.dir, r.GetFilename())
	out.Debug(ctx, "Get(%s) - Reading from %s", name, path)
	return ioutil.ReadFile(path)
}

func filename(buf []byte) string {
	sum := fmt.Sprintf("%x", sha256.Sum256(buf))
	return filepath.Join(sum[0:2], sum[2:])
}

func (o *OnDisk) Set(ctx context.Context, name string, value []byte) error {
	fn := filename(value)
	fp := filepath.Join(o.dir, filename(value))
	if err := os.MkdirAll(filepath.Dir(fp), 0700); err != nil {
		return err
	}
	if err := ioutil.WriteFile(fp, value, 0600); err != nil {
		return err
	}
	out.Debug(ctx, "Set(%s) - Wrote to %s", name, fp)
	e := o.getEntry(ctx, name)
	msg := "Updated " + fn
	if cm := ctxutil.GetCommitMessage(ctx); cm != "" {
		msg = cm
	}
	e.Revisions = append(e.Revisions, &gpb.Revision{
		Created: &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
		},
		Message:  msg,
		Filename: fn,
	})
	out.Debug(ctx, "Set(%s) - Added Revision", name)
	o.idx.Entries[name] = e
	return o.saveIndex()
}

func (o *OnDisk) getEntry(ctx context.Context, name string) *gpb.Entry {
	if e, found := o.idx.Entries[name]; found && e != nil {
		return e
	}
	out.Debug(ctx, "getEntry(%s) - Created new Entry", name)
	return &gpb.Entry{
		Name:      name,
		Revisions: make([]*gpb.Revision, 0, 1),
	}
}

func (o *OnDisk) Delete(ctx context.Context, name string) error {
	if !o.Exists(ctx, name) {
		out.Debug(ctx, "Delete(%s) - Not adding tombstone for non-existing entry", name)
		return nil
	}
	// add tombstone
	e := o.getEntry(ctx, name)
	e.Delete(ctxutil.GetCommitMessage(ctx))
	o.idx.Entries[name] = e

	out.Debug(ctx, "Delete(%s) - Added tombstone")
	return o.saveIndex()
}

func (o *OnDisk) Exists(ctx context.Context, name string) bool {
	_, found := o.idx.Entries[name]
	out.Debug(ctx, "Exists(%s): %t", name, found)
	return found
}

func (o *OnDisk) List(ctx context.Context, prefix string) ([]string, error) {
	res := make([]string, 0, len(o.idx.Entries))
	for k, v := range o.idx.Entries {
		if v.IsDeleted() {
			continue
		}
		if strings.HasPrefix(k, prefix) {
			res = append(res, k)
		}
	}
	return res, nil
}

func (o *OnDisk) IsDir(ctx context.Context, name string) bool {
	return false
}

func (o *OnDisk) Prune(ctx context.Context, prefix string) error {
	l, _ := o.List(ctx, name)
	for _, e := range l {
		if err := o.Delete(ctx, e); err != nil {
			return err
		}
	}
	return nil
}

func (o *OnDisk) Name() string {
	return name
}

func (o *OnDisk) Version(context.Context) semver.Version {
	return semver.Version{Major: 1}
}

func (o *OnDisk) String() string {
	return fmt.Sprintf("%s(path: %s)", name, o.dir)
}

func (o *OnDisk) Available(ctx context.Context) error {
	return nil
}

// Compact will prune all deleted entries and truncate every other entry
// to the last 10 revisions.
func (o *OnDisk) Compact(ctx context.Context) error {
	for k, v := range o.idx.Entries {
		if v.IsDeleted() {
			delete(o.idx.Entries, k)
			continue
		}
		if len(o.idx.Entries[k].Revisions) > 10 {
			o.idx.Entries[k].Revisions = o.idx.Entries[k].Revisions[0:10]
		}
	}
	return o.saveIndex()
}
