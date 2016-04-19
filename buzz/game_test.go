package buzz_test

import (
	"testing"

	"github.com/derailed/go-fizz/buzz"
)

func TestPlay(t *testing.T) {
	var (
		expected = []string{"1", "2", "Fizz", "4", "Buzz"}
		g        = buzz.NewGame(5)
		res      = g.Play()
	)

	for i, e := range expected {
		if res[i] != e {
			t.Fatalf("Mismatch expected `%s got `%s", e, res[i])
		}
	}
}
