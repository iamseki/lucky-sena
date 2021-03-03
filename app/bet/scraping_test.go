package betusecases

import (
	"lucky-sena/domain"
	"testing"
	"time"
)

type fakeLastBetScrapper struct {
	ScrapMock func(url string) (domain.Bet, error)
}

func (fs *fakeLastBetScrapper) Scrap(url string) (domain.Bet, error) {
	return fs.ScrapMock(url)
}

func newFakeScrapper() *fakeLastBetScrapper {
	return &fakeLastBetScrapper{
		ScrapMock: func(url string) (domain.Bet, error) {
			return domain.Bet{
				Numbers: []int{1, 2, 3, 4, 5, 6},
				Code:    17,
				Date:    time.Now(),
			}, nil
		},
	}
}
func TestScrap(t *testing.T) {
	s := newFakeScrapper()
	sut := NewLastBetScrapper(s)

	bet, _ := sut.Scrap("any-url")
	if bet.Code != 17 {
		t.Error("expected bet code is 17 but got", bet.Code)
	}
}
