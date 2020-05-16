package ondisk

import (
	"context"
	"fmt"

	"github.com/gopasspw/gopass/pkg/backend"
)

func (o *OnDisk) Add(ctx context.Context, args ...string) error {
	return nil
}

func (o *OnDisk) Commit(ctx context.Context, msg string) error {
	return nil
}

func (o *OnDisk) Push(ctx context.Context, remote, location string) error {
	return fmt.Errorf("not yet implemented")
}

func (o *OnDisk) Pull(ctx context.Context, remote, location string) error {
	return fmt.Errorf("not yet implemented")
}

func (o *OnDisk) InitConfig(ctx context.Context, name, email string) error {
	return nil
}

func (o *OnDisk) AddRemote(ctx context.Context, remote, location string) error {
	return fmt.Errorf("not yet implemented")
}

func (o *OnDisk) RemoveRemote(ctx context.Context, remote string) error {
	return fmt.Errorf("not yet implemented")
}

func (o *OnDisk) Revisions(ctx context.Context, name string) ([]backend.Revision, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (o *OnDisk) GetRevision(ctx context.Context, name, revision string) ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (o *OnDisk) Status(ctx context.Context) ([]byte, error) {
	return nil, nil
}
