package main

import (
	"fmt"

	"github.com/piquette/finance-go/quote"
)

func main() {
	q, err := quote.Get("AAPL")
	if err != nil {
		panic(err)
	}

	fmt.Println(q)
}
