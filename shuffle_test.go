package shuffler

import (
  "math/rand"
  "testing"
)

type Tester struct {
  val int
}

var float32Sl = []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var intSl = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
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

func TestShuffleFloat32(t *testing.T) {
  rand.Seed(0)
  // copy the original
  test := make([]float32, len(intSl))
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

func TestShuffleUint(t *testing.T) {
  rand.Seed(0)
  // copy the original
  test := make([]uint, len(intSl))
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
