package ondisk

import (
	"context"

	"github.com/gopasspw/gopass/pkg/backend"
	"github.com/gopasspw/gopass/pkg/out"
)

const (
	name = "ondisk"
)

func init() {
	backend.RegisterStorage(backend.OnDisk, name, &loader{})
	backend.RegisterRCS(backend.OnDiskRCS, name, &loader{})
}

type loader struct{}

func (l loader) New(ctx context.Context, url *backend.URL) (backend.Storage, error) {
	be, err := New(url.Path)
	out.Debug(ctx, "Using Storage Backend: %s", be.String())
	return be, err
}

func (l loader) Open(ctx context.Context, path string) (backend.RCS, error) {
	be, err := New(path)
	out.Debug(ctx, "Using RCS Backend: %s", be.String())
	return be, err
}

func (l loader) Clone(ctx context.Context, repo, path string) (backend.RCS, error) {
	be, err := New(path)
	out.Debug(ctx, "Using RCS Backend: %s", be.String())
	return be, err
}

func (l loader) Init(ctx context.Context, path, username, email string) (backend.RCS, error) {
	be, err := New(path)
	out.Debug(ctx, "Using RCS Backend: %s", be.String())
	return be, err
}

func (l loader) String() string {
	return name
}
