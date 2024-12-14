package age

import (
	"context"
	"testing"
)

func TestWithOnlyNative(t *testing.T) {
	ctx := context.Background()
	ctx = WithOnlyNative(ctx, true)

	if !IsOnlyNative(ctx) {
		t.Errorf("expected true, got false")
	}

	ctx = WithOnlyNative(ctx, false)

	if IsOnlyNative(ctx) {
		t.Errorf("expected false, got true")
	}
}

func TestIsOnlyNative_Default(t *testing.T) {
	ctx := context.Background()

	if IsOnlyNative(ctx) {
		t.Errorf("expected false, got true")
	}
}
