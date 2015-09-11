package shuffler

import (
  "math/rand"
  "testing"
)

type Tester struct {
  val int
}

var complex64Sl = []complex64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var complex128Sl = []complex128{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var float32Sl = []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var float64Sl = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var intSl = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var int8Sl = []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var int16Sl = []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var int32Sl = []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var uintSl = []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var stringSl = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

func TestShuffleInterface(t *testing.T) {
  rand.Seed(0)
  test := make([]interface{}, 10)
  for i := 0; i < 10; i++ {
    test[i] = Tester{i}
  }
  // make it into a slice of interface
  ShuffleInterface(test)
  expected := []int{8,2,3,0,5,7,1,6,9,4}
  for i, v := range test {
    if expected[i] != v.(Tester).val {
      t.Errorf("Expected %d got %d", expected[i], v.(Tester).val)
    }
  }
}

func TestShuffleComplex64(t *testing.T) {
  rand.Seed(0)
  // copy the original
  test := make([]complex64, len(complex64Sl))
  n := copy(test, complex64Sl)
  if n != len(complex64Sl) {
    t.Errorf("short copy: expected %d to be copied, %d were", len(complex64Sl), n)
    return
  }
  ShuffleComplex64(test)
  expected := []complex64{8,2,3,0,5,7,1,6,9,4}
  for i, v := range test {
    if expected[i] != v {
      t.Errorf("Expected %d got %d", expected[i], v)
    }
  }
}

func TestShuffleComplex128(t *testing.T) {
  rand.Seed(0)
  // copy the original
  test := make([]complex128, len(complex128Sl))
  n := copy(test, complex128Sl)
  if n != len(complex128Sl) {
    t.Errorf("short copy: expected %d to be copied, %d were", len(complex128Sl), n)
    return
  }
  ShuffleComplex128(test)
  expected := []complex128{8,2,3,0,5,7,1,6,9,4}
  for i, v := range test {
    if expected[i] != v {
      t.Errorf("Expected %d got %d", expected[i], v)
    }
  }
}

func TestShuffleFloat32(t *testing.T) {
  rand.Seed(0)
  // copy the original
  test := make([]float32, len(float32Sl))
  n := copy(test, float32Sl)
  if n != len(float32Sl) {
    t.Errorf("short copy: expected %d to be copied, %d were", len(float32Sl), n)
    return
  }
  ShuffleFloat32(test)
  expected := []float32{8,2,3,0,5,7,1,6,9,4}
  for i, v := range test {
    if expected[i] != v {
      t.Errorf("Expected %d got %d", expected[i], v)
    }
  }
}

func TestShuffleFloat64(t *testing.T) {
  rand.Seed(0)
  // copy the original
  test := make([]float64, len(float64Sl))
  n := copy(test, float64Sl)
  if n != len(float64Sl) {
    t.Errorf("short copy: expected %d to be copied, %d were", len(float64Sl), n)
    return
  }
  ShuffleFloat64(test)
  expected := []float64{8,2,3,0,5,7,1,6,9,4}
  for i, v := range test {
    if expected[i] != v {
      t.Errorf("Expected %d got %d", expected[i], v)
    }
  }
}

func TestShuffleInt(t *testing.T) {
  rand.Seed(0)
  // copy the original
  test := make([]int, len(intSl))
  n := copy(test, intSl)
  if n != len(intSl) {
    t.Errorf("short copy: expected %d to be copied, %d were", len(intSl), n)
    return
  }
  ShuffleInt(test)
  expected := []int{8,2,3,0,5,7,1,6,9,4}
  for i, v := range test {
    if expected[i] != v {
      t.Errorf("Expected %d got %d", expected[i], v)
    }
  }
}

func TestShuffleInt8(t *testing.T) {
  rand.Seed(0)
  // copy the original
  test := make([]int8, len(int8Sl))
  n := copy(test, int8Sl)
  if n != len(int8Sl) {
    t.Errorf("short copy: expected %d to be copied, %d were", len(int8Sl), n)
    return
  }
  ShuffleInt8(test)
  expected := []int8{8,2,3,0,5,7,1,6,9,4}
  for i, v := range test {
    if expected[i] != v {
      t.Errorf("Expected %d got %d", expected[i], v)
    }
  }
}

func TestShuffleInt16(t *testing.T) {
  rand.Seed(0)
  // copy the original
  test := make([]int16, len(int16Sl))
  n := copy(test, int16Sl)
  if n != len(int16Sl) {
    t.Errorf("short copy: expected %d to be copied, %d were", len(int16Sl), n)
    return
  }
  ShuffleInt16(test)
  expected := []int16{8,2,3,0,5,7,1,6,9,4}
  for i, v := range test {
    if expected[i] != v {
      t.Errorf("Expected %d got %d", expected[i], v)
    }
  }
}

func TestShuffleInt32(t *testing.T) {
  rand.Seed(0)
  // copy the original
  test := make([]int32, len(int32Sl))
  n := copy(test, int32Sl)
  if n != len(int32Sl) {
    t.Errorf("short copy: expected %d to be copied, %d were", len(int32Sl), n)
    return
  }
  ShuffleInt32(test)
  expected := []int32{8,2,3,0,5,7,1,6,9,4}
  for i, v := range test {
    if expected[i] != v {
      t.Errorf("Expected %d got %d", expected[i], v)
    }
  }
}

func TestShuffleUint(t *testing.T) {
  rand.Seed(0)
  // copy the original
  test := make([]uint, len(uintSl))
  n := copy(test, uintSl)
  if n != len(uintSl) {
    t.Errorf("short copy: expected %d to be copied, %d were", len(uintSl), n)
    return
  }
  ShuffleUint(test)
  expected := []uint{8,2,3,0,5,7,1,6,9,4}
  for i, v := range test {
    if expected[i] != v {
      t.Errorf("Expected %d got %d", expected[i], v)
    }
  }
}

func TestShuffleString(t *testing.T) {
  rand.Seed(0)
  // copy the original
  test := make([]string, len(stringSl))
  n := copy(test, stringSl)
  if n != len(stringSl) {
    t.Errorf("short copy: expected %d to be copied, %d were", len(stringSl), n)
    return
  }
  ShuffleString(test)
  expected := []string{"i", "c", "d", "a", "f", "h", "b", "g", "j", "e"}
  for i, v := range test {
    if expected[i] != v {
      t.Errorf("Expected %q got %q", expected[i], v)
    }
  }
}
