// Copyright 2014, Joel Scoble, all rights reserved.
// Licensed under The MIT License. Please view the LICENSE file for more
// information.
//
package shuffler

import (
	"crypto/rand"
	"math/big"
)

// ShuffleInterface randomizes an interface slice
func ShuffleInterface(c []interface{}) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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

// ShuffleByte randomizes a byte slice.
func ShuffleByte(c []byte) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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

// ShuffleComplex64 randomizes a complex64 slice.
func ShuffleComplex64(c []complex64) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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

// ShuffleComplex129 randomizes a complex128 slice.
func ShuffleComplex128(c []complex128) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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

// ShuffleFloat32 randomizes a float32 slice.
func ShuffleFloat32(c []float32) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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

// ShuffleFloat64 randomizes a float64 slice.
func ShuffleFloat64(c []float64) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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

// ShuffleInt randomizes an int slice.
func ShuffleInt(c []int) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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

// ShuffleInt8 randomizes an int8 slice.
func ShuffleInt8(c []int8) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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

// ShuffleInt16 randomizes an int16 slice.
func ShuffleInt16(c []int16) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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

// ShuffleInt32 randomizes an int32 slice.
func ShuffleInt32(c []int32) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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

// ShuffleInt64 randomizes an int64 slice.
func ShuffleInt64(c []int64) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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

// ShuffleString randomizes a strings slice.
func ShuffleString(c []string) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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

// ShuffleUint randomizes an uint slice.
func ShuffleUint(c []uint) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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

// ShuffleUint8 randomizes an uint8 slice.
func ShuffleUint8(c []uint8) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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

// ShuffleUint16 randomizes an uint16 slice.
func ShuffleUint16(c []uint16) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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

// ShuffleUint32 randomizes an uint32 slice.
func ShuffleUint32(c []uint32) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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

// ShuffleUint64 randomizes an uint64 slice.
func ShuffleUint64(c []uint64) error {
	b := new(big.Int)
	for i := len(c) - 1; i >= 0; i-- {
		b = b.SetInt64(int64(i+1))
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
