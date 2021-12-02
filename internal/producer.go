package internal

import (
	"context"
	"fmt"
	"sync"
)

//go:generate go run github.com/golang/mock/mockgen --source=producer.go --destination=producer_mocks.go --package=internal

type PriceStreamSubscriber interface {
	SubscribePriceStream(Ticker) (chan TickerPrice, chan error)
}

type Stream struct {
	subscriber PriceStreamSubscriber
}

// Produce read from input channel and produce to output.
func (s *Stream) Produce(t Ticker, output chan<- TickerPrice, ctx context.Context, wg *sync.WaitGroup) {
	var (
		input   chan TickerPrice
		errChan chan error
	)
	defer wg.Done()
	input, errChan = s.subscriber.SubscribePriceStream(t)
	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-input:
			if !ok {
				input, errChan = s.subscriber.SubscribePriceStream(t)
			}
			output <- val
		case err := <-errChan:
			fmt.Println(err)

		}
	}
}

type Create struct{}

func (Create) CreateProducer(p PriceStreamSubscriber) *Stream {
	return &Stream{subscriber: p}
}
