package internal

// I've decided that for a simpler solution decorator subscriberWithRetry is not needed.

//
//type SubscriberWithRetry struct {
//	subscriber PriceStreamSubscriber
//
//}
//
//func (s *SubscriberWithRetry) SubscribePriceStream(t Ticker) (chan TickerPrice, chan error) {
//	resChan, resErr := s.subscriber.SubscribePriceStream(t)
//
//}
