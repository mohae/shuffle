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
	l := len(set)
	for i := 0; i < l; i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleFloat32 randomizes a slice of float32s. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleFloat32(set []float32) {
	l := len(set)
	for i := 0; i < l; i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleInt randomizes a slice of ints. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleInt(set []int) {
	l := len(set)
	for i := 0; i < l; i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleUint randomizes a slice of uints. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleUint(set []uint) {
	l := len(set)
	for i := 0; i < l; i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleString randomizes a slice of strings. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleString(set []string) {
	l := len(set)
	for i := 0; i < l; i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}
