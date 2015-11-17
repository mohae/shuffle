shuffle
========
[![Build Status](https://travis-ci.org/mohae/shuffle.png)](https://travis-ci.org/mohae/shuffle)

Shuffle randomizes data within slices useing the Fisher-Yates shuffle algorithm.

Shuffle has implementations using both `math/rand` and `crypto/rand`. Use the one that fits your use case. The `math/rand` implementations can be found in `shuffle/quick`; their function names are slightly different than shuffle's.

## String and Numeric types
Shuffles are implemented for slices of Go string and numeric types. These shuffles modify the recieved slice; no additional slice related allocations will occur.

## Shuffler interface
Other types can be shuffled by implementing the `shuffler` interface and calling either `shuffle.Interface(c Shuffler)`, for `crypto/rand` based shuffles, or `quick.Shuffle(c shuffle.Shuffler)` for `math/rand` based shuffle.

## Using `shuffle`
Shuffles using `crypto/rand`. While the slices are modified in place, `shuffle` funcs return an error as getting a random number from `crypto/rand` may result in an error.

Import the library:

    import github.com/mohae/shuffle

Randomize:

    err := shuffle.Int(x)

## Using `shuffle/quick`
Quick shuffles allocate minimal memory and are much faster than `crypto/rand` shuffles.  The slices are modified in place and nothing is returned.  Rand should be properly seeded before use; that is left up to the user of the package.

Import the library:

    import github.com/mohae/shuffle/quick


Import `math/rand`:

    import math/rand

Seed `math/rand`:

    rand.Seed(time.Now().UTC().UnixNanao())

* While `time.Now().UTC().UnixNano()` is a popular method for seeding rand, one should strongly consider using a different method as this is vulnerable to timing. *

Randomize:

    quick.ShuffleInt(x)

## License

Copyright 2015 Joel Scoble, all rights reserved.

Licensed under The MIT License. Please see the included license file for more details.
