package protocols

import "lucky-sena/domain/bet"

type CSVConverter interface {
	Convert([]bet.Bet) error
}
