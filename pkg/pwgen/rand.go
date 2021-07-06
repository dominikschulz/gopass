package pwgen

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"time"
)

func init() {
	// seed math/rand in case we have to fall back to using it
	rand.Seed(time.Now().Unix() + int64(os.Getpid()+os.Getppid()))
}

var (
	cReader = crand.Reader
	warn    = true
)

func randomInteger(max int) int {
	i, err := crand.Int(cReader, big.NewInt(int64(max)))
	if err == nil {
		return int(i.Int64())
	}
	if warn {
		fmt.Fprintln(os.Stderr, "WARNING: No crypto/rand available. Falling back to PRNG")
	}
	return rand.Intn(max)
}
