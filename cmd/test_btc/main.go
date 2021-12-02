package main

import (
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"test_btc/internal"
	"test_btc/internal/calculator"
	"time"
)

func main() {
	pool := internal.NewPool(internal.Create{}, internal.NewConsumer(calculator.NewEMACalculator(7)))

	l := logrus.New()
	ctrl := gomock.NewController(l)
	defer ctrl.Finish()

	c1 := make(chan internal.TickerPrice)
	c2 := make(chan internal.TickerPrice)
	c3 := make(chan internal.TickerPrice)
	e1 := make(chan error)
	e2 := make(chan error)
	e3 := make(chan error)
	mockStream :=  internal.NewMockPriceStreamSubscriber(ctrl)
	mockStream.EXPECT().SubscribePriceStream(internal.BTCUSDTicker).Return(c1, e1)
	mockStream.EXPECT().SubscribePriceStream(internal.BTCUSDTicker).Return(c2, e2)
	mockStream.EXPECT().SubscribePriceStream(internal.BTCUSDTicker).Return(c3, e3)
	go func() {
		for {

			c2 <- internal.TickerPrice{
				Ticker: "",
				Time:   time.Now().Add(5 * time.Second),
				Price:  "100",
			}
			c2 <- internal.TickerPrice{
				Ticker: "",
				Time:   time.Now().Add(6 * time.Second),
				Price:  "100",
			}
			c2 <- internal.TickerPrice{
				Ticker: "",
				Time:   time.Now().Add(7 * time.Second),
				Price:  "100",
			}
			c2 <- internal.TickerPrice{
				Ticker: "",
				Time:   time.Now().Add(8 * time.Second),
				Price:  "200",
			}
			c2 <- internal.TickerPrice{
				Ticker: "",
				Time:   time.Now().Add(9 * time.Second),
				Price:  "200",
			}
			c2 <- internal.TickerPrice{
				Ticker: "",
				Time:   time.Now().Add(10 * time.Second),
				Price:  "200",
			}
			time.Sleep(time.Second)
			c2 <- internal.TickerPrice{
				Ticker: "",
				Time:   time.Now().Add(-1 * time.Minute),
				Price:  "1000000",
			}

			time.Sleep(1 * time.Minute)
			close(c2)
			c3 <- internal.TickerPrice{
				Ticker: "",
				Time:   time.Now().Add(9 * time.Second),
				Price:  "200",
			}
			c3 <- internal.TickerPrice{
				Ticker: "",
				Time:   time.Now().Add(10 * time.Second),
				Price:  "200",
			}
			time.Sleep(1 * time.Minute)
		}
	}()

	pool.Start(2, mockStream)


}


