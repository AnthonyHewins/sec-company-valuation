package main

import (
	"os"
	"fmt"
	"flag"
	"time"

	"github.com/AnthonyHewins/sec-company-valuation/pkg/fetch"
)

func main() {
	secYear       := flag.Int(     "y", time.Now().Year(), "Using DERA, fetch financial statements directly from the SEC's website")
	secQuarter    := flag.Int(     "q",                 1, "Using DERA, fetch financial statements directly from the SEC's website")
	thirdPartyUrl := flag.String("url",                "", "Fetch the financial statements elsewhere (still must match schema)")
	localZip      := flag.String("zip",                "", "Use a local ZIP to load the data (most favorable option)")

	flag.Parse()

	switch {
	case localZip != nil:
		fmt.Println(fetch.Zip(*localZip))
	case thirdPartyUrl != nil:
		fmt.Println(fetch.Url(*thirdPartyUrl))
	case secYear != nil || secQuarter != nil:
		fmt.Println(fetch.SecUrl(*secYear, *secQuarter))
	default:
		fmt.Println("Missing arguments. Exiting.")
		os.Exit(1)
	}
}
