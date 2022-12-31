package alt

import "math"

func GetPartsLength(val float64, fmt byte, prec, bitSize int) (int, int) {
	var bits uint64
	var flt *floatInfo
	switch bitSize {
	case 32:
		bits = uint64(math.Float32bits(float32(val)))
		flt = &float32info
	case 64:
		bits = math.Float64bits(val)
		flt = &float64info
	default:
		panic("strconv: illegal AppendFloat/FormatFloat bitSize")
	}

	exp := int(bits>>flt.mantbits) & (1<<flt.expbits - 1)
	mant := bits & (uint64(1)<<flt.mantbits - 1)

	switch exp {
	case 1<<flt.expbits - 1:
		// Inf, NaN
		return -1, -1

	case 0:
		// denormalized
		exp++

	default:
		// add implicit top bit
		mant |= uint64(1) << flt.mantbits
	}
	exp += flt.bias

	var digs decimalSlice
	// Negative precision means "only as much as needed to be exact."
	shortest := prec < 0
	if shortest {
		// Use Ryu algorithm.
		var buf [32]byte
		digs.d = buf[:]
		ryuFtoaShortest(&digs, mant, exp-int(flt.mantbits), flt)
	} else if fmt != 'f' {
		// Fixed number of digits.
		digits := prec
		switch fmt {
		case 'e', 'E':
			digits++
		case 'g', 'G':
			if prec == 0 {
				prec = 1
			}
			digits = prec
		default:
			// Invalid mode.
			digits = 1
		}
		var buf [24]byte
		if bitSize == 32 && digits <= 9 {
			digs.d = buf[:]
			ryuFtoaFixed32(&digs, uint32(mant), exp-int(flt.mantbits), digits)
		} else if digits <= 18 {
			digs.d = buf[:]
			ryuFtoaFixed64(&digs, mant, exp-int(flt.mantbits), digits)
		}
	}

	return digs.dp, digs.nd - digs.dp
}
