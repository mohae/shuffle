// Copyright 2014, Joel Scoble, all rights reserved.
// Licensed under The MIT License. Please view the LICENSE file for more
// information.
//
package shuffler

import (
	"math/rand"
)

// Takes a 'set', a slice of interfaces, and shuffles the set, returning a
// slice of interfaces.
func ShuffleIface(set []interface{}) []interface{} {
	l := len(set)
	for i := 0; i < l; i++ {
		r := rand.Intn(i + 1)
		if i != r {
			set[r], set[i] = set[i], set[r]
		}
	}
	return set
}
