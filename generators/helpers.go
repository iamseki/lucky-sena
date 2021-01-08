package generators

func containValue(numbers []int, value int) bool {
	for _, n := range numbers {
		if n == value {
			return true
		}
	}
	return false
}
