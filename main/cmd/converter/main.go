package main

import "lucky-sena/main/factories"

func main() {
	c := factories.NewBetToCsvConverter()
	c.ConvertBetsToCsv()
}
