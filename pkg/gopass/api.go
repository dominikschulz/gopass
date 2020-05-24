package gopass

import (
	"context"
	"fmt"
	"time"
)

type Header interface {
	Get(string) string
	Set(string, string)
	Keys() []string
}

type Byter interface {
	Bytes() ([]byte, error)
}

type Secret interface {
	fmt.Stringer
	Byter

	Header() Header
	//Password() Password
}

type Revision interface {
	Created() time.Time
	Message() string
	Unwrap() Secret
}

type Entry interface {
	Revisions() ([]string, error)
	Latest() (Revision, error)
	Revision(string) (Revision, error)
}

// Store is a secret store.
type Store interface {
	fmt.Stringer

	// List all secrets
	List(context.Context) ([]string, error)
	// Get an encrypted secret. Check existence with Get; err != nil
	Get(ctx context.Context, name string) (Entry, error)
	// Set (add) a new revision of an secret
	Set(ctx context.Context, name string, sec Byter) error
	// Remove a single secret
	Remove(ctx context.Context, name string) error
	// RemoveAll secrets with a common prefix
	RemoveAll(ctx context.Context, prefix string) error
	// Rename a path (secret of prefix) without decrypting
	Rename(ctx context.Context, src, dest string) error
	// Sync with a remote (if configured)
	// WARNING: This might be dropped if we decide to always auto-sync!
	Sync(ctx context.Context) error
}
