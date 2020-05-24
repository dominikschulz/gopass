package secret

import (
	"fmt"

	"github.com/gopasspw/gopass/internal/store/secret/legacy"
)

func Parse(buf []byte) (*legacy.Secret, error) {
	if s, err := legacy.Parse(buf); err == nil {
		return s, nil
	}
	return nil, fmt.Errorf("not supported")
}
