package calculator


const (
	maxPeriod = 60
	minPeriod = 2
)

// Calculator can add value for calculation, reset values, and get fair price.
type Calculator interface {
	AddValue(secondNum int, value float64)
	ResetCalculator()
	// GetValue get fair price.
	GetValue() float64
}

type price struct {
	val float64
	cnt int
}

// Calc implement Calculator strategy.
type Calc struct {
	c Calculator
}

func NewCalc(c Calculator) *Calc {
	return &Calc{c: c}
}

func (c *Calc) SetCalculatorType(calc Calculator) {
	c.c = calc
}

func (c *Calc) AddValue(secondNum int, value float64) {
	c.c.AddValue(secondNum, value)
}

func (c *Calc) ResetCalculator() {
	c.c.ResetCalculator()
}

func (c *Calc) GetValue() float64 {
	return c.c.GetValue()
}


