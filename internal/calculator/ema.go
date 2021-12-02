package calculator

import (
	"test_btc/config"
)


type EMA struct {
	groupSize int
	period int
	val []*price
}

func NewEMACalculator(period int) *EMA{
	if period > maxPeriod  {
		period = maxPeriod
	}

	if period <= 1 {
		period = minPeriod
	}

	return &EMA{
		period: period,
		val: make([]*price, config.SecondsInMinute),
	}
}

func (e *EMA) AddValue(secondNum int, value float64) {
	if e.val[secondNum] == nil {
		e.val[secondNum] = &price{
			val: value,
			cnt: 1,
		}
		return
	}
	e.val[secondNum].val += value
	e.val[secondNum].cnt +=1
}

func (e *EMA) ResetCalculator() {
	e.val = make([]*price, config.SecondsInMinute)
}

func (e *EMA) GetValue() (res float64) {
	resData := make([]float64,0)

	for _, v := range e.val {
		if v != nil {
			resData= append(resData, v.val / float64(v.cnt))
		}
	}

	var (
		alpha = 2 / float64(e.period + 1)
		cnt int
	)


	for k, v := range resData {
		if k < e.period {
			cnt ++
			res = res + (v - res) / float64(cnt)
		} else {
			res = (v-res) * alpha + res
		}
	}

	return res
}
