package buzz

import (
	"testing"
)

func TestFizz(t *testing.T) {
	expected := map[int]interface{}{
		0:  "",
		1:  "",
		3:  "Fizz",
		10: "Buzz",
		30: "FizzBuzz",
	}

	for k, e := range expected {
		if ans, _ := toFizz(k); ans != e {
			t.Fatalf("[%d] Mismatch expected %s got %s", k, e, ans)
		}
	}
}
