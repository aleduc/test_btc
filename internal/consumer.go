package internal

import (
	"context"
	"fmt"
	"strconv"
	"test_btc/config"
	"test_btc/internal/calculator"
	"time"
)

type Time interface {
	Now() time.Time
}

type Consume struct {
	calc           calculator.Calculator
	tim            Time
	previousMinute time.Time
	nextMinute     time.Time
}

func NewConsumer(calc calculator.Calculator) *Consume {
	return &Consume{
		calc: calc,
	}
}

// ConsumeResults reads tickerPrices, outputs results.
func (c *Consume) ConsumeResults(results <-chan TickerPrice, ctx context.Context) {
	now := time.Now()
	tick := time.NewTicker(calculateMinuteDuration(now))
	c.fillMinuteRanges(now)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case res := <-results:
			if inTimePeriod(c.previousMinute, c.nextMinute, res.Time) {
				if res.Price == "" {
					break
				}
				val, err := strconv.ParseFloat(res.Price, 64)
				if err != nil {
					break
				}
				c.calc.AddValue(res.Time.Second(), val)
			}
		case <-tick.C:
			now := time.Now()
			tick.Reset(calculateMinuteDuration(now))
			// This for situation, when we start exactly in tick and not get data.
			if val := c.calc.GetValue(); val != 0 {
				fmt.Println(fmt.Sprintf("%d %f", now.Unix(), val))
			}
			c.fillMinuteRanges(now)
			c.calc.ResetCalculator()
		}
	}
}

func (c *Consume) fillMinuteRanges(t time.Time) {
	c.previousMinute = t.Truncate(time.Minute)
	c.nextMinute = c.previousMinute.Add(time.Minute)
}

func calculateMinuteDuration(t time.Time) time.Duration {
	return time.Second * time.Duration(config.SecondsInMinute-t.Second())
}

func inTimePeriod(start, end, val time.Time) bool {
	return val.After(start) && val.Before(end)
}
