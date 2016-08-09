shuffle  
========  
[![GoDoc](https://godoc.org/github.com/mohae/json2go?status.svg)](https://godoc.org/github.com/mohae/shuffle)[![Build Status](https://travis-ci.org/mohae/shuffle.png)](https://travis-ci.org/mohae/shuffle)

Shuffle randomizes data within slices using the Fisher-Yates shuffle algorithm.

Shuffle has implementations using both a [PCG](http://www.pcg-random.org/) PRNG and `crypto/rand`. Use the one that fits your use case. The PRNG implementations can be found in `shuffle/quick`.

## String and Numeric types  
Shuffles are implemented for slices of Go bytes, strings, and numeric types. These shuffles modify the received slice; no additional slice related allocations will occur.

## Shuffler interface  
Other types can be shuffled by implementing the `shuffler` interface and calling either `shuffle.Interface(c Shuffler)`, for `crypto/rand` based shuffles, or `quick.Shuffle(c shuffle.Shuffler)` for PRNG based shuffle.

See `bench_test.go`'s `Deck` for an example implementation.

## `shuffle`  
Shuffles the elements using `crypto/rand`.  If there is an error obtaining a random value, a panic will occur.

## `shuffle/quick`  
Quick shuffles allocate minimal memory, if any, and are much faster than `crypto/rand` shuffles.  The slices are modified in place and nothing is returned.  The PRNG is seeded using a value obtained from `crypto/rand` during initialization; if this operation results in an error, a panic will occur.

## License

Copyright 2015-16 Joel Scoble, all rights reserved.

Licensed under The MIT License. Please see the included license file for more details.
