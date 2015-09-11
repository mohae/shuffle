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
