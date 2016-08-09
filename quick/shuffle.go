// Copyright 2014, Joel Scoble, all rights reserved.
// Licensed under The MIT License. Please view the LICENSE file for more
// information.

// Shuffle provides math/rand based shuffling (randomization) of collections
// using the Fisher-Yates (Knuth) shuffling algorithm.  Functions for
// shuffling slices of non-composite types are provided, or you can implement
// the shuffle.Shuffler interface and shuffle using the Shuffle func.
//
// Shuffling is performed on the received slice; nothing is returned and no
// additional allocations are made.
//
// If a CSPRG based shuffle is needed, use the github.com/mohae/shuffle
// package.
package quick

import (
	"crypto/rand"
	"fmt"
	"math/big"

	pcg "github.com/dgryski/go-pcgr"
	"github.com/mohae/shuffle"
)

var rng pcg.Rand

func init() {
	err := Seed()
	if err != nil {
		panic(fmt.Sprintf("entropy read error: %s\n", err))
	}
}

// Seed seeds the rng by getting a new seed value from crypto/rand.  If
// the attempt to get a new seed value results in an error, that will be
// returned.
func Seed() error {
	bi := big.NewInt(1<<63 - 1)
	r, err := rand.Int(rand.Reader, bi)
	if err != nil {
		return err
	}
	rng.Seed(r.Int64())
	return nil
}

// Shuffle randomizes collections. A nil error will be returned.
func Shuffle(c shuffle.Shuffler) error {
	l := c.Len()
	for i := l - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c.Swap(i, j)
		}
	}
	return nil
}

// ShuffleByte randomizes a byte slice. A nil error will be returned.
func ShuffleByte(c []byte) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleComplex64 randomizes a complex64 slice. A nil error will be returned.
func ShuffleComplex64(c []complex64) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleComplex129 randomizes a complex128 slice. A nil error will be
// returned.
func ShuffleComplex128(c []complex128) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleFloat32 randomizes a float32 slice. A nil error will be returned.
func ShuffleFloat32(c []float32) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleFloat64 randomizes a float64 slice. A nil error will be returned.
func ShuffleFloat64(c []float64) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleInt randomizes an int slice. A nil error will be returned.
func ShuffleInt(c []int) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleInt8 randomizes an int8 slice. A nil error will be returned.
func ShuffleInt8(c []int8) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleInt16 randomizes an int16 slice. A nil error will be returned.
func ShuffleInt16(c []int16) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleInt32 randomizes an int32 slice. A nil error will be returned.
func ShuffleInt32(c []int32) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleInt64 randomizes an int64 slice. A nil error will be returned.
func ShuffleInt64(c []int64) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleString randomizes a strings slice. A nil error will be returned.
func ShuffleString(c []string) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleUint randomizes an uint slice. A nil error will be returned.
func ShuffleUint(c []uint) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleUint8 randomizes an uint8 slice. A nil error will be returned.
func ShuffleUint8(c []uint8) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleUint16 randomizes an uint16 slice. A nil error will be returned.
func ShuffleUint16(c []uint16) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleUint32 randomizes an uint32 slice. A nil error will be returned.
func ShuffleUint32(c []uint32) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}

// ShuffleUint64 randomizes an uint64 slice. A nil error will be returned.
func ShuffleUint64(c []uint64) error {
	for i := len(c) - 1; i >= 0; i-- {
		j := int(rng.Bound(uint32(i + 1)))
		if i != j {
			c[j], c[i] = c[i], c[j]
		}
	}
	return nil
}
