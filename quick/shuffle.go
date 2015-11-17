// Copyright 2014, Joel Scoble, all rights reserved.
// Licensed under The MIT License. Please view the LICENSE file for more
// information.

// Shuffle provides math/rand based shuffling (randomization) of collections
// using the Fisher-Yates (Knuth) shuffling algorithm.  Functions for
// shuffling slices of non-composite types are provided, or you can implement
// the Shuffler interface and shuffle using the shuffle.Interface() func.
//
// Shuffling is performed on the received slice; nothing is returned and no
// additional allocations are made.
//
// If a CSPRG based shuffle is needed, use the github.com/mohae/shuffle
// package.
package quick

import (
	"math/rand"

	"github.com/mohae/shuffle"
)

// Shuffle randomizes collections
func Shuffle(c shuffle.Shuffler) error {
	l := c.Len()
	for i := l - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		c.Swap(i, j)
	}
	return nil
}

// ShuffleByte randomizes a byte slice.
func ShuffleByte(c []byte) {
	for i := len(c) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
}

// ShuffleComplex64 randomizes a complex64 slice.
func ShuffleComplex64(c []complex64) {
	for i := len(c) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
}

// ShuffleComplex129 randomizes a complex128 slice.
func ShuffleComplex128(c []complex128) {
	for i := len(c) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
}

// ShuffleFloat32 randomizes a float32 slice.
func ShuffleFloat32(c []float32) {
	for i := len(c) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
}

// ShuffleFloat64 randomizes a float64 slice.
func ShuffleFloat64(c []float64) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleInt randomizes an int slice.
func ShuffleInt(c []int) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleInt8 randomizes an int8 slice.
func ShuffleInt8(c []int8) {
	for i := len(c) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
}

// ShuffleInt16 randomizes an int16 slice.
func ShuffleInt16(c []int16) {
	for i := len(c) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
}

// ShuffleInt32 randomizes an int32 slice.
func ShuffleInt32(c []int32) {
	for i := len(c) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		c[j], c[i] = c[i], c[j]
	}
}

// ShuffleInt64 randomizes an int64 slice.
func ShuffleInt64(c []int64) {
	for i := len(c) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
}

// ShuffleString randomizes a strings slice.
func ShuffleString(c []string) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleUint randomizes an uint slice.
func ShuffleUint(c []uint) {
	for i := len(c) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
}

// ShuffleUint8 randomizes an uint8 slice.
func ShuffleUint8(c []uint8) {
	for i := len(c) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
}

// ShuffleUint16 randomizes an uint16 slice.
func ShuffleUint16(c []uint16) {
	for i := len(c) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
}

// ShuffleUint32 randomizes an uint32 slice.
func ShuffleUint32(c []uint32) {
	for i := len(c) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
}

// ShuffleUint64 randomizes an uint64 slice.
func ShuffleUint64(c []uint64) {
	for i := len(c) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
}
