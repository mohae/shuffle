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

func TestShuffleInterfaces(t *testing.T) {
  rand.Seed(0)
  itesters := make([]interface{}, 10)
  for i := 0; i < 10; i++ {
    itesters[i] = Tester{i}
  }
  // make it into a slice of interface
  ShuffleInterfaces(itesters)
  expected := []int{8,2,3,0,5,7,1,6,9,4}
  for i, v := range itesters {
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

func TestShuffleInts(t *testing.T) {
  rand.Seed(0)
  // copy the original
  testInts := make([]int, len(intSl))
  n := copy(testInts, intSl)
  if n != len(intSl) {
    t.Errorf("short copy: expected %d to be copied, %d were", len(intSl), n)
    return
  }
  ShuffleInts(testInts)
  expected := []int{8,2,3,0,5,7,1,6,9,4}
  for i, v := range testInts {
    if expected[i] != v {
      t.Errorf("Expected %d got %d", expected[i], v)
    }
  }
}

func TestShuffleUints(t *testing.T) {
  rand.Seed(0)
  // copy the original
  test := make([]uint, len(intSl))
  n := copy(test, uintSl)
  if n != len(uintSl) {
    t.Errorf("short copy: expected %d to be copied, %d were", len(uintSl), n)
    return
  }
  ShuffleUints(test)
  expected := []uint{8,2,3,0,5,7,1,6,9,4}
  for i, v := range test {
    if expected[i] != v {
      t.Errorf("Expected %d got %d", expected[i], v)
    }
  }
}

func TestShuffleStrings(t *testing.T) {
  rand.Seed(0)
  // copy the original
  testStrings := make([]string, len(stringSl))
  n := copy(testStrings, stringSl)
  if n != len(stringSl) {
    t.Errorf("short copy: expected %d to be copied, %d were", len(stringSl), n)
    return
  }

  ShuffleStrings(testStrings)
  expected := []string{"i", "c", "d", "a", "f", "h", "b", "g", "j", "e"}
  for i, v := range testStrings {
    if expected[i] != v {
      t.Errorf("Expected %q got %q", expected[i], v)
    }
  }
}
