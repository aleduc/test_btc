package internal

// I've decided that for a simpler solution I shouldn't use any custom/parametrised struct for ticker.
//import "time"
//
//
//type MinuteTick struct {
//	tick *time.Ticker
//}
//
//func StartNewMinuteTick(quit <-chan bool) *MinuteTick{
//	return &MinuteTick{tick: time.NewTicker(calculateMinuteDuration()) }
//}
//
//
