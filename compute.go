package safecomp

import (
	"math"

	"github.com/zergon321/safecomp/alt"
)

func Sum(a, b float64) float64 {
	_, aFracLen := alt.GetPartsLength(a, 'f', -1, 64)
	_, bFracLen := alt.GetPartsLength(b, 'f', -1, 64)

	aModifier := math.Pow10(aFracLen)
	bModifier := math.Pow10(bFracLen)
	modifier := math.Max(aModifier, bModifier)

	a *= modifier
	b *= modifier

	return (a + b) / modifier
}

func Sub(a, b float64) float64 {
	_, aFracLen := alt.GetPartsLength(a, 'f', -1, 64)
	_, bFracLen := alt.GetPartsLength(b, 'f', -1, 64)

	aModifier := math.Pow10(aFracLen)
	bModifier := math.Pow10(bFracLen)
	modifier := math.Max(aModifier, bModifier)

	a *= modifier
	b *= modifier

	return (a - b) / modifier
}

func Mul(a, b float64) float64 {
	_, aFracLen := alt.GetPartsLength(a, 'f', -1, 64)
	_, bFracLen := alt.GetPartsLength(b, 'f', -1, 64)

	aModifier := math.Pow10(aFracLen)
	bModifier := math.Pow10(bFracLen)
	modifier := aModifier * bModifier

	a *= aModifier
	b *= bModifier

	return a * b / modifier
}

func Div(a, b float64) float64 {
	_, aFracLen := alt.GetPartsLength(a, 'f', -1, 64)
	_, bFracLen := alt.GetPartsLength(b, 'f', -1, 64)

	aModifier := math.Pow10(aFracLen)
	bModifier := math.Pow10(bFracLen)
	minModifier := math.Min(aModifier, bModifier)
	maxModifier := math.Max(aModifier, bModifier)
	modifier := maxModifier / minModifier

	a *= aModifier
	b *= bModifier
	result := a / b

	if aModifier > bModifier {
		return result / modifier
	}

	if bModifier > aModifier {
		return result * modifier
	}

	return result
}
