// Copyright 2014, Joel Scoble, all rights reserved.
// Licensed under The MIT License. Please view the LICENSE file for more
// information.

// Shuffle provides crypto/rand based shuffling (randomization) of collections
// using the Fisher-Yates (Knuth) shuffling algorithm.  Functions for shuffling
// slices of non-composite types are provided, or you can implement the
// Shuffler interface and shuffle using the shuffle.Interface() func.
//
// Shuffling is performed on the received slice; nothing is returned and no
// additional allocations, aside from those caused by using crypto/rand.
//
// If using a math/rand based shuffle is good enough, see the
// github.com/mohae/shuffle/quick package, which provides a quick,
// pseudo-random shuffle.
package shuffle

import (
	"crypto/rand"
	"math/big"
)

// Shuffler interface.
type Shuffler interface {
	Len() int      // number of elements in the p
	Swap(i, j int) // swaps elements
}

// Shuffle randomizes collections: use this for Shuffler implementations.
func Interface(c Shuffler) error {
	b := new(big.Int)
	l := c.Len()
	for i := l - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		c.Swap(i, int(j.Int64()))
	}
	return nil
}

// Instead of using an interface, these funcs shuffle Go types by accepting
// a slice of T and shuffling.  The Shuffle func is named after the Type that
// it shuffles.  Since everything is done in place, the slice header is not
// modified: nothing is returned.

// Byte randomizes a byte slice.
func Byte(c []byte) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		if i != int(j.Int64()) {
			c[int(j.Int64())], c[i] = c[i], c[int(j.Int64())]
		}
	}
	return nil
}

// Complex64 randomizes a complex64 slice.
func Complex64(c []complex64) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		if i != int(j.Int64()) {
			c[int(j.Int64())], c[i] = c[i], c[int(j.Int64())]
		}
	}
	return nil
}

// Complex129 randomizes a complex128 slice.
func Complex128(c []complex128) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		if i != int(j.Int64()) {
			c[int(j.Int64())], c[i] = c[i], c[int(j.Int64())]
		}
	}
	return nil
}

// Float32 randomizes a float32 slice.
func Float32(c []float32) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		if i != int(j.Int64()) {
			c[int(j.Int64())], c[i] = c[i], c[int(j.Int64())]
		}
	}
	return nil
}

// Float64 randomizes a float64 slice.
func Float64(c []float64) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		if i != int(j.Int64()) {
			c[int(j.Int64())], c[i] = c[i], c[int(j.Int64())]
		}
	}
	return nil
}

// Int randomizes an int slice.
func Int(c []int) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		if i != int(j.Int64()) {
			c[int(j.Int64())], c[i] = c[i], c[int(j.Int64())]
		}
	}
	return nil
}

// Int8 randomizes an int8 slice.
func Int8(c []int8) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		if i != int(j.Int64()) {
			c[int(j.Int64())], c[i] = c[i], c[int(j.Int64())]
		}
	}
	return nil
}

// Int16 randomizes an int16 slice.
func Int16(c []int16) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		if i != int(j.Int64()) {
			c[int(j.Int64())], c[i] = c[i], c[int(j.Int64())]
		}
	}
	return nil
}

// Int32 randomizes an int32 slice.
func Int32(c []int32) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		if i != int(j.Int64()) {
			c[int(j.Int64())], c[i] = c[i], c[int(j.Int64())]
		}
	}
	return nil
}

// Int64 randomizes an int64 slice.
func Int64(c []int64) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		if i != int(j.Int64()) {
			c[int(j.Int64())], c[i] = c[i], c[int(j.Int64())]
		}
	}
	return nil
}

// String randomizes a strings slice.
func String(c []string) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		if i != int(j.Int64()) {
			c[int(j.Int64())], c[i] = c[i], c[int(j.Int64())]
		}
	}
	return nil
}

// Uint randomizes an uint slice.
func Uint(c []uint) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		if i != int(j.Int64()) {
			c[int(j.Int64())], c[i] = c[i], c[int(j.Int64())]
		}
	}
	return nil
}

// Uint8 randomizes an uint8 slice.
func Uint8(c []uint8) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		if i != int(j.Int64()) {
			c[int(j.Int64())], c[i] = c[i], c[int(j.Int64())]
		}
	}
	return nil
}

// Uint16 randomizes an uint16 slice.
func Uint16(c []uint16) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		if i != int(j.Int64()) {
			c[int(j.Int64())], c[i] = c[i], c[int(j.Int64())]
		}
	}
	return nil
}

// Uint32 randomizes an uint32 slice.
func Uint32(c []uint32) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		if i != int(j.Int64()) {
			c[int(j.Int64())], c[i] = c[i], c[int(j.Int64())]
		}
	}
	return nil
}

// Uint64 randomizes an uint64 slice.
func Uint64(c []uint64) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i + 1))
		j, err := rand.Int(rand.Reader, b)
		if err != nil {
			return err
		}
		if i != int(j.Int64()) {
			c[int(j.Int64())], c[i] = c[i], c[int(j.Int64())]
		}
	}
	return nil
}
