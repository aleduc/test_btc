package calculator

import (
	"test_btc/config"
)

type SMA struct {
	cnt int
	val float64
}

func NewSMACalculator() *SMA {
	return &SMA{}
}

func (s *SMA) AddValue(secondNum int, value float64) {
	s.cnt++
	s.val = s.val + (value-s.val)/float64(s.cnt)
}

func (s *SMA) ResetCalculator() {
	s.val = 0
	s.cnt = 0
}

func (s *SMA) GetValue() float64 {
	return s.val
}

// SMAPeriod implement SMA on last n seconds.
// I could implement some kind of cyclic buffer here, if it were not for the principles in the readme.
type SMAPeriod struct {
	period int
	val    []*price
}

func NewSMAPeriod(period int) *SMAPeriod {
	if period > maxPeriod {
		period = maxPeriod
	}
	if period <= 1 {
		period = minPeriod
	}
	return &SMAPeriod{
		period: period,
		val:    make([]*price, config.SecondsInMinute),
	}
}

func (s *SMAPeriod) AddValue(secondNum int, value float64) {
	if s.val[secondNum] == nil {
		s.val[secondNum] = &price{
			val: value,
			cnt: 1,
		}
		return
	}
	s.val[secondNum].val += value
	s.val[secondNum].cnt += 1
}

func (s *SMAPeriod) ResetCalculator() {
	// For garbage collector this reinit is not a problem, but for online solution is better.
	s.val = make([]*price, config.SecondsInMinute)
}

func (s *SMAPeriod) GetValue() (res float64) {
	realPeriod := s.period
	counter := realPeriod
	for i := len(s.val) - 1; i >= 0; i-- {
		if counter == 0 {
			break
		}
		v := s.val[i]
		if v != nil {
			res += v.val / float64(v.cnt)
			counter--
		}
	}
	if counter != 0 {
		realPeriod = realPeriod - counter
	}
	if realPeriod == 0 {
		return
	}
	return res / float64(realPeriod)
}
