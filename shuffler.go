// Copyright 2014, Joel Scoble, all rights reserved.
// Licensed under The MIT License. Please view the LICENSE file for more
// information.
//
package shuffler

import (
	"math/rand"
)

// ShuffleInterfaces randomizes a slice of interfaces. Since everything is
// done in place, the slice header is not modified: nothing is returned.
func ShuffleInterfaces(set []interface{}) {
	l := len(set)
	for i := 0; i < l; i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleInts randomizes a slice of ints. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleInts(set []int) {
	l := len(set)
	for i := 0; i < l; i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}

// ShuffleStrings randomizes a slice of ints. Since everything is done in
// place, the slice header is not modified: nothing is returned.
func ShuffleStrings(set []string) {
	l := len(set)
	for i := 0; i < l; i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
}
