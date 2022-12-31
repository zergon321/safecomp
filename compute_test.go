package safecomp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zergon321/safecomp"
)

func TestMulFloat64(t *testing.T) {
	assert.Equal(t, 17999.205, safecomp.Mul(2090.5, 8.61))
	assert.Equal(t, 0.3, safecomp.Mul(3, 0.1))
}

func TestSubFloat64(t *testing.T) {
	assert.Equal(t, 0.2, safecomp.Sub(10, 9.8))
	assert.Equal(t, 0.15, safecomp.Sub(0.2, 0.05))
	assert.Equal(t, 0.199995, safecomp.Sub(0.2, 0.000005))
	assert.Equal(t, 0.199999995, safecomp.Sub(0.2, 0.000000005))
}

func TestDivFloat64(t *testing.T) {
	assert.Equal(t, 25_000.0, safecomp.
		Div(25, 0.001))
	assert.Equal(t, 0.025, safecomp.
		Div(25, 1000))
	assert.Equal(t, 9.54, safecomp.
		Div(28.62, 3))
	assert.Equal(t, 3.0, safecomp.
		Div(40.83, 13.61))
	assert.Equal(t, 0.5, safecomp.
		Div(13.61, 27.22))
	assert.Equal(t, 0.05, safecomp.
		Div(13.61, 272.2))
	assert.Equal(t, 0.05, safecomp.
		Div(0.1361, 2.722))
}

func BenchmarkSafeMulFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		safecomp.Mul(2090.5, 8.61)
	}
}

func BenchmarkRegularMulFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = 2090.5 * 8.61
	}
}
