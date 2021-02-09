package protocols

import "lucky-sena/domain"

type CSVConverter interface {
	Convert([]domain.Bet) error
}
