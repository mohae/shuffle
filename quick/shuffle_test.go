package quick

import "testing"

type Tester struct {
	val int
}

var byteSl = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
var complex64Sl = []complex64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var complex128Sl = []complex128{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var float32Sl = []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var float64Sl = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var intSl = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var int8Sl = []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var int16Sl = []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var int32Sl = []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var int64Sl = []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var uintSl = []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var uint8Sl = []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var uint16Sl = []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var uint32Sl = []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var uint64Sl = []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var stringSl = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

func TestShuffleByte(t *testing.T) {
	rng.Seed(0)
	// copy the original
	test := make([]byte, len(byteSl))
	n := copy(test, byteSl)
	if n != len(byteSl) {
		t.Errorf("short copy: expected %d to be copied, %d were", len(byteSl), n)
		return
	}
	ShuffleByte(test)
	expected := []byte{'f', 'e', 'j', 'b', 'c', 'g', 'a', 'd', 'h', 'i'}
	for i, v := range test {
		if expected[i] != v {
			t.Errorf("Expected %s got %s", string(expected[i]), string(v))
		}
	}
}

func TestShuffleComplex64(t *testing.T) {
	rng.Seed(0)
	// copy the original
	test := make([]complex64, len(complex64Sl))
	n := copy(test, complex64Sl)
	if n != len(complex64Sl) {
		t.Errorf("short copy: expected %d to be copied, %d were", len(complex64Sl), n)
		return
	}
	ShuffleComplex64(test)
	expected := []complex64{(5 + 0i), (4 + 0i), (9 + 0i), (1 + 0i), (2 + 0i), (6 + 0i), (0 + 0i), (3 + 0i), (7 + 0i), (8 + 0i)}
	for i, v := range test {
		if expected[i] != v {
			t.Errorf("Expected %v got %v", expected[i], v)
		}
	}
}

func TestShuffleComplex128(t *testing.T) {
	rng.Seed(0)
	// copy the original
	test := make([]complex128, len(complex128Sl))
	n := copy(test, complex128Sl)
	if n != len(complex128Sl) {
		t.Errorf("short copy: expected %d to be copied, %d were", len(complex128Sl), n)
		return
	}
	ShuffleComplex128(test)
	expected := []complex128{(5 + 0i), (4 + 0i), (9 + 0i), (1 + 0i), (2 + 0i), (6 + 0i), (0 + 0i), (3 + 0i), (7 + 0i), (8 + 0i)}
	for i, v := range test {
		if expected[i] != v {
			t.Errorf("Expected %v got %v", expected[i], v)
		}
	}
}

func TestShuffleFloat32(t *testing.T) {
	rng.Seed(0)
	// copy the original
	test := make([]float32, len(float32Sl))
	n := copy(test, float32Sl)
	if n != len(float32Sl) {
		t.Errorf("short copy: expected %d to be copied, %d were", len(float32Sl), n)
		return
	}
	ShuffleFloat32(test)
	expected := []float32{5, 4, 9, 1, 2, 6, 0, 3, 7, 8}
	for i, v := range test {
		if expected[i] != v {
			t.Errorf("Expected %v got %v", expected[i], v)
		}
	}
}

func TestShuffleFloat64(t *testing.T) {
	rng.Seed(0)
	// copy the original
	test := make([]float64, len(float64Sl))
	n := copy(test, float64Sl)
	if n != len(float64Sl) {
		t.Errorf("short copy: expected %d to be copied, %d were", len(float64Sl), n)
		return
	}
	ShuffleFloat64(test)
	expected := []float64{5, 4, 9, 1, 2, 6, 0, 3, 7, 8}
	for i, v := range test {
		if expected[i] != v {
			t.Errorf("Expected %v got %v", expected[i], v)
		}
	}
}

func TestShuffleInt(t *testing.T) {
	rng.Seed(0)
	// copy the original
	test := make([]int, len(intSl))
	n := copy(test, intSl)
	if n != len(intSl) {
		t.Errorf("short copy: expected %d to be copied, %d were", len(intSl), n)
		return
	}
	ShuffleInt(test)
	expected := []int{5, 4, 9, 1, 2, 6, 0, 3, 7, 8}
	for i, v := range test {
		if expected[i] != v {
			t.Errorf("Expected %v got %v", expected[i], v)
		}
	}
}

func TestShuffleInt8(t *testing.T) {
	rng.Seed(0)
	// copy the original
	test := make([]int8, len(int8Sl))
	n := copy(test, int8Sl)
	if n != len(int8Sl) {
		t.Errorf("short copy: expected %d to be copied, %d were", len(int8Sl), n)
		return
	}
	ShuffleInt8(test)
	expected := []int8{5, 4, 9, 1, 2, 6, 0, 3, 7, 8}
	for i, v := range test {
		if expected[i] != v {
			t.Errorf("Expected %d got %d", expected[i], v)
		}
	}
}

func TestShuffleInt16(t *testing.T) {
	rng.Seed(0)
	// copy the original
	test := make([]int16, len(int16Sl))
	n := copy(test, int16Sl)
	if n != len(int16Sl) {
		t.Errorf("short copy: expected %d to be copied, %d were", len(int16Sl), n)
		return
	}
	ShuffleInt16(test)
	expected := []int16{5, 4, 9, 1, 2, 6, 0, 3, 7, 8}
	for i, v := range test {
		if expected[i] != v {
			t.Errorf("Expected %d got %d", expected[i], v)
		}
	}
}

func TestShuffleInt32(t *testing.T) {
	rng.Seed(0)
	// copy the original
	test := make([]int32, len(int32Sl))
	n := copy(test, int32Sl)
	if n != len(int32Sl) {
		t.Errorf("short copy: expected %d to be copied, %d were", len(int32Sl), n)
		return
	}
	ShuffleInt32(test)
	expected := []int32{5, 4, 9, 1, 2, 6, 0, 3, 7, 8}
	for i, v := range test {
		if expected[i] != v {
			t.Errorf("Expected %d got %d", expected[i], v)
		}
	}
}

func TestShuffleInt64(t *testing.T) {
	rng.Seed(0)
	// copy the original
	test := make([]int64, len(int64Sl))
	n := copy(test, int64Sl)
	if n != len(int64Sl) {
		t.Errorf("short copy: expected %d to be copied, %d were", len(int64Sl), n)
		return
	}
	ShuffleInt64(test)
	expected := []int64{5, 4, 9, 1, 2, 6, 0, 3, 7, 8}
	for i, v := range test {
		if expected[i] != v {
			t.Errorf("Expected %d got %d", expected[i], v)
		}
	}
}

func TestShuffleUint(t *testing.T) {
	rng.Seed(0)
	// copy the original
	test := make([]uint, len(uintSl))
	n := copy(test, uintSl)
	if n != len(uintSl) {
		t.Errorf("short copy: expected %d to be copied, %d were", len(uintSl), n)
		return
	}
	ShuffleUint(test)
	expected := []uint{5, 4, 9, 1, 2, 6, 0, 3, 7, 8}
	for i, v := range test {
		if expected[i] != v {
			t.Errorf("Expected %d got %d", expected[i], v)
		}
	}
}

func TestShuffleUint8(t *testing.T) {
	rng.Seed(0)
	// copy the original
	test := make([]uint8, len(uint8Sl))
	n := copy(test, uint8Sl)
	if n != len(uint8Sl) {
		t.Errorf("short copy: expected %d to be copied, %d were", len(uint8Sl), n)
		return
	}
	ShuffleUint8(test)
	expected := []uint8{5, 4, 9, 1, 2, 6, 0, 3, 7, 8}
	for i, v := range test {
		if expected[i] != v {
			t.Errorf("Expected %d got %d", expected[i], v)
		}
	}
}

func TestShuffleUint16(t *testing.T) {
	rng.Seed(0)
	// copy the original
	test := make([]uint16, len(uint16Sl))
	n := copy(test, uint16Sl)
	if n != len(uint16Sl) {
		t.Errorf("short copy: expected %d to be copied, %d were", len(uint16Sl), n)
		return
	}
	ShuffleUint16(test)
	expected := []uint16{5, 4, 9, 1, 2, 6, 0, 3, 7, 8}
	for i, v := range test {
		if expected[i] != v {
			t.Errorf("Expected %d got %d", expected[i], v)
		}
	}
}

func TestShuffleUint32(t *testing.T) {
	rng.Seed(0)
	// copy the original
	test := make([]uint32, len(uint32Sl))
	n := copy(test, uint32Sl)
	if n != len(uint32Sl) {
		t.Errorf("short copy: expected %d to be copied, %d were", len(uint32Sl), n)
		return
	}
	ShuffleUint32(test)
	expected := []uint32{5, 4, 9, 1, 2, 6, 0, 3, 7, 8}
	for i, v := range test {
		if expected[i] != v {
			t.Errorf("Expected %d got %d", expected[i], v)
		}
	}
}

func TestShuffleUint64(t *testing.T) {
	rng.Seed(0)
	// copy the original
	test := make([]uint64, len(uint64Sl))
	n := copy(test, uint64Sl)
	if n != len(uint64Sl) {
		t.Errorf("short copy: expected %d to be copied, %d were", len(uint64Sl), n)
		return
	}
	ShuffleUint64(test)
	expected := []uint64{5, 4, 9, 1, 2, 6, 0, 3, 7, 8}
	for i, v := range test {
		if expected[i] != v {
			t.Errorf("Expected %d got %d", expected[i], v)
		}
	}
}

func TestShuffleString(t *testing.T) {
	rng.Seed(0)
	// copy the original
	test := make([]string, len(stringSl))
	n := copy(test, stringSl)
	if n != len(stringSl) {
		t.Errorf("short copy: expected %d to be copied, %d were", len(stringSl), n)
		return
	}
	ShuffleString(test)
	expected := []string{"f", "e", "j", "b", "c", "g", "a", "d", "h", "i"}
	for i, v := range test {
		if expected[i] != v {
			t.Errorf("Expected %q got %q", expected[i], v)
		}
	}
}
