package shuffler

import (
  "math/rand"
  "testing"
)

type Tester struct {
  val int
}

func TestShuffleInterfaces(t *testing.T) {
    rand.Seed(0)
    itesters := make([]interface{}, 10)
    for i := 0; i < 10; i++ {
      itesters[i] = Tester{i}
    }
    // make it into a slice of interface
    itesters = ShuffleIface(itesters)
    expected := []int{8,2,3,0,5,7,1,6,9,4}
    for i, v := range itesters {
      if expected[i] != v.(Tester).val {
        t.Errorf("Expected %d got %d", expected[i], v.(Tester).val)
      }
    }
}
