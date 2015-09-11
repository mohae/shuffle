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
func ShuffleInterface(set []interface{}) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleByte randomizes a slice of bytes. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleByte(set []byte) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleComplex64 randomizes a slice of complex64s. Since everything is
// done in place, the slice header is not modified: nothing is returned.
func ShuffleComplex64(set []complex64) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleComplex129 randomizes a slice of complex128s. Since everything is
// done in place, the slice header is not modified: nothing is returned.
func ShuffleComplex128(set []complex128) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleFloat32 randomizes a slice of float32s. Since everything is done
// in place, the slice header is not modified: nothing is returned.
func ShuffleFloat32(set []float32) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleFloat64 randomizes a slice of float64s. Since everything is done
// in place, the slice header is not modified: nothing is returned.
func ShuffleFloat64(set []float64) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleInt randomizes a slice of ints. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleInt(set []int) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleInt8 randomizes a slice of int8s. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleInt8(set []int8) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleInt16 randomizes a slice of int16s. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleInt16(set []int16) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleInt32 randomizes a slice of int32s. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleInt32(set []int32) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleInt64 randomizes a slice of int64s. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleInt64(set []int64) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleUint randomizes a slice of uints. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleUint(set []uint) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleUint8 randomizes a slice of uint8s. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleUint8(set []uint8) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleUint16 randomizes a slice of uint16s. Since everything is done
// in place, the slice header is not modified: nothing is returned.
func ShuffleUint16(set []uint16) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleUint32 randomizes a slice of uint32s. Since everything is done
// in place, the slice header is not modified: nothing is returned.
func ShuffleUint32(set []uint32) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleUint64 randomizes a slice of uint64s. Since everything is done
// in place, the slice header is not modified: nothing is returned.
func ShuffleUint64(set []uint64) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleString randomizes a slice of strings. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleString(set []string) {
	for i := 0; i < len(set); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}
