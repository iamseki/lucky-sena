package generator

import (
	"fmt"
	"testing"
)

func TestIncludeMergedBet(t *testing.T) {
	bets := []GenaretedBet{
		{
			Numbers: []int{1, 2, 3, 4, 5, 6},
		},
		{
			Numbers: []int{8, 9, 11, 12, 13, 6},
		},
		{
			Numbers: []int{20, 19, 18, 17, 5, 6},
		},
	}

	m := includeMergedBet(bets, 3)

	fmt.Println(m)
}
