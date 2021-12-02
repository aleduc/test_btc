package internal

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	defaultExchanges = 2
	BTCUSDTicker Ticker = "BTC_USD"
)

// Pool struct for control our producers/consumers.
type Pool struct {
	producerCreator ProducerCreator
	consumer  Consumer
}


// Producer write values from PriceStreamSubscriber to output channel.
type Producer interface {
	Produce(t Ticker, output chan<- TickerPrice, ctx context.Context, wg *sync.WaitGroup)
}

type ProducerCreator interface {
	CreateProducer(p PriceStreamSubscriber) *Stream
}

type Consumer interface {
	ConsumeResults(results <-chan TickerPrice, ctx context.Context)
}

// NewPool init new Pool.
func NewPool(p ProducerCreator, consumer Consumer) *Pool {
	return &Pool{producerCreator: p, consumer: consumer}
}

// Start setup all producers/consumer, wait results/exit signals.
func (p *Pool) Start(maxExchanges int, streamSubscriber PriceStreamSubscriber) {
	if maxExchanges == 0 {
		maxExchanges = defaultExchanges
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	resultChan := make(chan TickerPrice, maxExchanges)

	go p.consumer.ConsumeResults(resultChan, ctx)

	for i := 1; i <= maxExchanges; i++ {
		producer := p.producerCreator.CreateProducer(streamSubscriber)
		wg.Add(1)
		go producer.Produce(BTCUSDTicker, resultChan, ctx,  wg)
	}


	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan
	cancelFunc()

	wg.Wait()
	close(resultChan)
	fmt.Println("Graceful shutdown")
}


