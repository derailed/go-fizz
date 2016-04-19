package buzz

import (
	"errors"
	"fmt"
)

var errNoAnswer = errors.New("No answer")

const (
	// FizzBuzz divisible by 3 and 5
	FizzBuzz = "FizzBuzz"
	// Fizz divisible by 3
	Fizz = "Fizz"
	// Buzz divisible by 5
	Buzz = "Buzz"
)

// Game of fizzbuzz
type Game struct {
	total int
}

// NewGame of FizzBuzz
func NewGame(total int) Game {
	return Game{total: total}
}

// Play a game of fizzbuzz
func (g Game) Play() []string {
	res := make([]string, g.total)
	for i, j := 0, 1; i < g.total; i, j = i+1, j+1 {
		a, err := toFizz(j)
		if err == errNoAnswer {
			a = fmt.Sprintf("%d", j)
		}
		res[i] = a
	}
	return res
}

func toFizz(num int) (s string, err error) {
	if num <= 0 {
		return s, errNoAnswer
	}

	switch {
	case num%3 == 0 && num%5 == 0:
		s = FizzBuzz
	case num%3 == 0:
		s = Fizz
	case num%5 == 0:
		s = Buzz
	default:
		err = errNoAnswer
	}
	return
}
