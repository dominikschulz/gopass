package ondisk

import (
	"context"
	"os"
	"path/filepath"

	"github.com/gopasspw/gopass/pkg/out"
)

func (o *OnDisk) Fsck(ctx context.Context) error {
	// build a list of existing files
	files := make(map[string]struct{}, len(o.idx.Entries)+1)
	files[idxFile] = struct{}{}
	for _, v := range o.idx.Entries {
		for _, r := range v.Revisions {
			files[r.Filename] = struct{}{}
		}
	}

	return filepath.Walk(o.dir, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() && len(fi.Name()) != 2 {
			out.Print(ctx, "Skipping unknown dir: %s", path)
			return filepath.SkipDir
		}
		// TODO check dir
		return nil
	})
}
