package shuffle

import (
	"fmt"
	"testing"
)

// benchmark stuff
var bfloat []float64
var bint []int
var bstring []string

type card struct {
	suit  string
	value int
}

type Deck struct {
	cards []card
}

func (d Deck) Len() int { return len(d.cards) }
func (d Deck) Swap(i, j int) {
	d.cards[i], d.cards[j] = d.cards[j], d.cards[i]

}

var deck Deck

// this setup is for benchmarking
func init() {
	max := 100
	alpha := "abcdefghijklmnopqrstuvwxyz"
	bfloat = make([]float64, max)
	bint = make([]int, max)
	bstring = make([]string, max)
	var l, s int
	for i := 0; i < max; i++ {
		var xyz string
		bfloat[i] = float64(i)
		bint[i] = i
		l = i / 26
		l = i - 26*l
		switch {
		case l < 20:
			s = 7
		default:
			s = 26 - l

		}
		if s < 7 {
			xyz = fmt.Sprintf("%s%s", alpha[l:l+s], alpha[:7-s])

		} else {
			xyz = alpha[l : l+s]

		}
		bstring[i] = xyz

	}
	var val, j int
	suit := []string{"club", "diamond", "heart", "spade"}
	for i := 0; i < 52; i++ {
		if val == 13 {
			j++
			val = 0

		}
		deck.cards = append(deck.cards, card{suit[j], val + 1})
		val++

	}

}

var result error

func BenchmarkShuffleInt(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		err = Int(bint)
	}
	result = err
}

func BenchmarkShuffleFloat64(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		err = Float64(bfloat)
	}
	result = err
}

var stringResult []string

func BenchmarkShuffleString(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		err = String(bstring)
	}
	result = err
}

// shuffle 52
func BenchmarkShuffleDeck(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		err = Interface(deck)
	}
	result = err
}
