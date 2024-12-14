package age

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseRecipients(t *testing.T) {
	ctx := context.Background()
	td := t.TempDir()
	t.Setenv("GOPASS_HOMEDIR", td)
	a, err := New(ctx)
	require.NoError(t, err)

	tests := []struct {
		name       string
		recipients []string
		expectErr  bool
	}{
		{
			name: "Valid age key",
			// generated wtih `age-keygen`
			recipients: []string{"age19x9xn602fpv4xhrphduzmmwvgdc89udgf38u5t2ztugh8kyykspscaz4qw"},
			expectErr:  false,
		},
		{
			name: "Valid Ed25519 SSH key",
			// generated with `ssh-keygen -t ed25519 -o /tmp/id_ed25519`
			recipients: []string{"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOHoA1b+LfsjbvcqiTwHjemXWc9BQw2V6fPHhY07UU9C"},
			expectErr:  false,
		},
		{
			name: "Valid RSA SSH key",
			// generated with `ssh-keygen -t rsa -o /tmp/id_rsa`
			recipients: []string{"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDAiMItKovYdiaaEefTC8YJdcMTgsvNkcrtr8qB7qqZDCSSBYKSdc5CFKfieTF6y2TCE/yqPzOLepG/D5xPwZb9Lnc9nYW6Slb+OyBwxQOODWqJB/xRZZFwYzjtljX9G8OB7fkvZg8Ujn1UWlYpkuDlIZKt7n1Edqm80bsFKAQIXOiD2lnq+iIMXARrBSbN7witQkVm2bmPc3ruaLGgziCzVStecIh6nu6TJ3Hs5Eo6gitFbkXPvnhzF8Aa/LPZHnlwp3ilK5S6ObxUMfH12UvHK2TDh6lBLZyDKegz5QcbU9rxtKtG8EOlixd5XjjfR0RGNZJBjaUhJMVLCdaofhREAZbtKoNGPRMfQBbpC87TVx0/XgP0ILDGPJMLHPFBQw9v8gbeZufzEp6P8qL3RK8W92FfZS5dLmEc3pDflVk1PCl5978wcKcYSJ2+2vAKgWMkgerUyRWlTjDa60TcPdbddBi7ZliApEFlmyemJoAV0/euCt7spzqAwzpw7C+8lUs="},
			expectErr:  false,
		},
		{
			name:       "Invalid age key",
			recipients: []string{"age1invalidkey"},
			expectErr:  true,
		},
		{
			name:       "Invalid SSH key",
			recipients: []string{"ssh-invalidkey"},
			expectErr:  true,
		},
		{
			name:       "Unknown key",
			recipients: []string{"unknown-key"},
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := a.parseRecipients(ctx, tt.recipients)
			if tt.expectErr {
				//assert.Error(t, err)
				assert.Len(t, ret, 0)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, ret)
		})
	}
}
