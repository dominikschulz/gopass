package pwgen

import (
	"fmt"
	mrand "math/rand"
	"strings"
	"testing"
)

type fixedReader struct{}

func (f *fixedReader) Read(p []byte) (int, error) {
	for i := 0; i < len(p); i++ {
		p[i] = 1
	}
	return len(p), nil
}

type brokenReader struct{}

func (b *brokenReader) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("b0rked")
}

func TestRandomInteger(t *testing.T) {
	//cReader = &fixedReader{}
	//cReader = &brokenReader{}
	mrand.Seed(1)
	warn = false

	lim := 100
	t.Logf("%d", mrand.Intn(lim))
	res := make(map[int]int, lim)
	for i := 0; i < 1000000; i++ {
		res[randomInteger(lim)]++
	}

	count := 0
	sum := 0
	min := 1000000
	max := 0
	for i := 0; i < lim; i++ {
		v := res[i]
		count++
		sum += v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		t.Logf("[% 5d] %s", i, strings.Repeat(".", v/10))
	}
	avg := sum / count

	t.Logf("count: %d - sum: %d - avg: %d - min: %d - max: %d", count, sum, avg, min, max)
}
