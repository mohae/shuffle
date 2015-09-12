// Copyright 2014, Joel Scoble, all rights reserved.
// Licensed under The MIT License. Please view the LICENSE file for more
// information.
//
package shuffler

import (
	"math/rand"
)

// ShuffleInterface randomizes a slice of interfaces. Since everything is
// done in place, the slice header is not modified: nothing is returned.
func ShuffleInterface(p []interface{}) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}

// ShuffleByte randomizes a slice of bytes. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleByte(p []byte) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}

// ShuffleComplex64 randomizes a slice of complex64s. Since everything is
// done in place, the slice header is not modified: nothing is returned.
func ShuffleComplex64(p []complex64) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}

// ShuffleComplex129 randomizes a slice of complex128s. Since everything is
// done in place, the slice header is not modified: nothing is returned.
func ShuffleComplex128(p []complex128) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}

// ShuffleFloat32 randomizes a slice of float32s. Since everything is done
// in place, the slice header is not modified: nothing is returned.
func ShuffleFloat32(p []float32) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}

// ShuffleFloat64 randomizes a slice of float64s. Since everything is done
// in place, the slice header is not modified: nothing is returned.
func ShuffleFloat64(p []float64) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}

// ShuffleInt randomizes a slice of ints. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleInt(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}

// ShuffleInt8 randomizes a slice of int8s. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleInt8(p []int8) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}

// ShuffleInt16 randomizes a slice of int16s. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleInt16(p []int16) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}

// ShuffleInt32 randomizes a slice of int32s. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleInt32(p []int32) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}

// ShuffleInt64 randomizes a slice of int64s. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleInt64(p []int64) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}

// ShuffleUint randomizes a slice of uints. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleUint(p []uint) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}

// ShuffleUint8 randomizes a slice of uint8s. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleUint8(p []uint8) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}

// ShuffleUint16 randomizes a slice of uint16s. Since everything is done
// in place, the slice header is not modified: nothing is returned.
func ShuffleUint16(p []uint16) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}

// ShuffleUint32 randomizes a slice of uint32s. Since everything is done
// in place, the slice header is not modified: nothing is returned.
func ShuffleUint32(p []uint32) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}

// ShuffleUint64 randomizes a slice of uint64s. Since everything is done
// in place, the slice header is not modified: nothing is returned.
func ShuffleUint64(p []uint64) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}

// ShuffleString randomizes a slice of strings. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleString(p []string) {
	for i := len(p) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if i != r {
			p[r], p[i] = p[i], p[r]
		}
	}
}
