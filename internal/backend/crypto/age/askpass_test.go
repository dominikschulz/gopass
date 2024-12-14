package age

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAskPass(t *testing.T) {
	ctx := context.Background()
	ap := newAskPass(ctx)
	assert.NotNil(t, ap)
	assert.NotNil(t, ap.cache)
}

func TestAskPass_Passphrase(t *testing.T) {
	ctx := context.Background()
	ap := newAskPass(ctx)
	ap.testing = true

	// Test setting and getting a passphrase
	key := "test-key"
	reason := "test-reason"
	repeat := false
	passphrase := "test-passphrase"

	ap.cache.Set(key, passphrase)
	retrievedPassphrase, err := ap.Passphrase(key, reason, repeat)
	assert.NoError(t, err)
	assert.Equal(t, passphrase, retrievedPassphrase)

	// Test getting a passphrase that is not in the cache
	ap.cache.Remove(key)
	got, err := ap.Passphrase(key, reason, repeat)
	assert.NoError(t, err)
	assert.Empty(t, got)
}

func TestAskPass_Remove(t *testing.T) {
	ctx := context.Background()
	ap := newAskPass(ctx)
	ap.testing = true

	key := "test-key"
	passphrase := "test-passphrase"

	ap.cache.Set(key, passphrase)
	ap.Remove(key)
	_, found := ap.cache.Get(key)
	assert.False(t, found)
}

func TestAskPass_Purge(t *testing.T) {
	ctx := context.Background()
	ap := newAskPass(ctx)
	ap.testing = true

	key := "test-key"
	passphrase := "test-passphrase"

	ap.cache.Set(key, passphrase)
	ap.cache.Set("another-key", "another-passphrase")
	ap.cache.Purge()

	_, found := ap.cache.Get(key)
	assert.False(t, found)
	_, found = ap.cache.Get("another-key")
	assert.False(t, found)
}
