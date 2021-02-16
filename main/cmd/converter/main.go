package main

import "lucky-sena/main/factories"

func main() {
	c := factories.NewResultsBetToCsvConverter()
	c.ConvertBetsToCsv()
}
