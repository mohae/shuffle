shuffler
========

Shuffler randomizes data within slices.

Shuffler uses the Fisher-Yates shuffle algorithm for randomizing slices.  The shuffle functions receive a slice and randomize the data within the slice.  Nothing is returned by the slice because the randomization is done by swapping elements: no reallocations are done.

Shuffle can be called by either data type, e.g. `ShuffleInt8(int8Slice)`, or by calling `Shuffle()`.  If  `Shuffle()` is called, the slice is passed as an `interface{}`. The type of the slice will be determine which shuffle function to call; if the type is not supported an error will be returned.

If the slice is a slice of a custom type, it will need to be passed as `[]interface{}`.

## Usage
Import the library:

    import github.com/mohae/shuffler

Import `math/rand`:

    import math/rand

Seed `math/rand`:

    rand.Seed(time.Now().UTC().UnixNanao())

* While `time.Now().UTC().UnixNano()` is a popular method for seeding rand, one should strongly consider using a different method as this is vulnerable to timing. *

Randomize:


## Notes:
Currently, `shuffler` uses `math/rand` which is not a CSPRNG.  This means the randomization is not truly random and the larger the set the worse the "randomization".

It is the caller's responsibility to seed the PRNG.

## License

Copyright 2014 Joel Scoble, all rights reserved.

Licensed under The MIT License. Please see the included license file for more details.
