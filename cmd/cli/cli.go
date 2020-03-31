package main

import (
	"fmt"

	"github.com/AnthonyHewins/sec-company-valuation/pkg/fetch"
)

func main() {
	fmt.Println(fetch.SecURL(2010, 1))
}
