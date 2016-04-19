// Slices and map + coverage
package main

import (
	"flag"
	"fmt"

	"github.com/derailed/go-fizz/buzz"
)

func main() {
	var total = flag.Int("t", 20, "Bizzbuzz iterations ")
	flag.Parse()

	res := buzz.NewGame(*total).Play()
	for i, r := range res {
		fmt.Printf("%02d\t%s\n", i, r)
	}
}
