package safecomp_test

import (
	"testing"

	"github.com/zergon321/safecomp/alt"
)

func BenchmarkGetPartsLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		alt.GetPartsLength(-19367.5889, 'f', -1, 64)
	}
}
