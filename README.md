## Principles
1. KISS.
1. Avoid premature optimization.
1. Avoid third-party libs.

_I didn't follow SOLID because it's small test task._
_I did not write unit tests, so there may be some small bugs._ 


### What is "fair" price
It depends on business requirements. I found out one hint in task. 
> price should be the most relevant to the bar timestamp.

That's why SMA and EMA are represented here. My final choice it is EMA7, but it can be changed for more suitable variant in any moment due to the pattern "strategy".

### Producer
* The interface responsible for writing the values from our stream to the output channel. 
* All producers have one shared output channel. (Channel multiplexing)
##### Stream
Implements producer. 
1. Gets the channel from PriceStreamSubscriber and then reads the data from the channel.
1. If the channel is closed, gets a new channel from PriceStreamSubscriber.

### Consumer
* The interface responsible for reading the values from our output channel and process them. 
##### Consume
Implements Consumer. In a real task, you should put some logic in separate interfaces.
1. Starts minute ticker
1. Reads data from output channel, verifies them, and adds data to calculator.
1. Each beginning of a new minute gets the results from Calculator, resets calculator and timer.

### Calculator 
* The interface responsible for getting fair price for last minute.
##### SMA/SMAPeriod/EMA
Implements Calculator. 
1. SMA - simple moving average for all values.
1. SMAPeriod - moving average for any custom period(between 2 and 60). 
1. EMA - exponential moving average for any custom period(between 2 and 60)
First ema value = SMA for period(maybe it was not the best choice, and best will be SMA for 5-6 values)
1. Period unit - one second. 

##### Pool
1. Starts producers, consumers and controls them(graceful shutdown).


##### Additional comment. 
I forget to write that I could make several consumers + calculator with storage for async work + some kind of outer, but decided to simplify solution.