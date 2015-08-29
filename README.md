shuffler
========

shuffler has shuffling implementations, well probably only one, but it could have others...at some point

**Implements Fisher-Yates**

## Usage
Import the library:

    import github.com/mohae/shuffler

Import `math/rand`:

    import math/rand

Seed `math/rand`:

    rand.Seed(time.Now().UTC().UnixNanao())

* While `time.Now().UTC().UnixNano()` is a popular method for seeding rand, one should strongly consider using a different method as this is vulnerable to timing. *

Shuffle/randomize:

    []mySet = shuffler.FisherYates([]mySet)

## Notes:
Currently, `shuffler` uses `math/rand` which is not a CSPRNG.

It is the caller's responsibility to seed the PRNG.


Copyright 2014 Joel Scoble, all rights reserved.

Licensed under The MIT License. Please see the included license file for more details.
