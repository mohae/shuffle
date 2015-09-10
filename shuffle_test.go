package shuffler

import (
  "math/rand"
  "testing"
)

type Tester struct {
  val int
}

var ints = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var strings = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

func TestShuffleInterfaces(t *testing.T) {
  rand.Seed(0)
  itesters := make([]interface{}, 10)
  for i := 0; i < 10; i++ {
    itesters[i] = Tester{i}
  }
  // make it into a slice of interface
  itesters = ShuffleInterfaces(itesters)
  expected := []int{8,2,3,0,5,7,1,6,9,4}
  for i, v := range itesters {
    if expected[i] != v.(Tester).val {
      t.Errorf("Expected %d got %d", expected[i], v.(Tester).val)
    }
  }
}

func TestShuffleInts(t *testing.T) {
  rand.Seed(0)
  // make it into a slice of interface
  ShuffleInts(ints)
  expected := []int{8,2,3,0,5,7,1,6,9,4}
  for i, v := range ints {
    if expected[i] != v {
      t.Errorf("Expected %d got %d", expected[i], v)
    }
  }
}

func TestShuffleStrings(t *testing.T) {
  rand.Seed(0)
  // make it into a slice of interface
  ShuffleStrings(strings)
  expected := []string{"i", "c", "d", "a", "f", "h", "b", "g", "j", "e"}
  for i, v := range strings {
    if expected[i] != v {
      t.Errorf("Expected %q got %q", expected[i], v)
    }
  }
}
