package main

import (
	"flag"
	"lucky-sena/main/factories"
)

func main() {
	var collection string
	flag.StringVar(&collection, "collection", "results", "collection to be read")
	flag.StringVar(&collection, "c", "results", "shortly flag to what collection to be read")

	c := factories.NewResultsBetToCsvConverter()
	c.ConvertBetsToCsv()
}
