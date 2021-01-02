package main

import (
	"flag"
	"fmt"
)

func main() {
	var count int
	flag.IntVar(&count, "c", 1, "number of facts")
	flag.Parse()
	for i := 0; i < count; i++ {
		var fact CatFact = getFact()
		fmt.Println(fact.Text + "\n")
	}
}
