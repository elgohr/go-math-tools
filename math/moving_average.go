package math

import "context"

func MovingAverage[T Number](ctx context.Context, input <-chan T, windowSize int) <-chan T {
	movingAverage := make(chan T)
	go func() {
		window := make([]T, windowSize)
		var counter int
		var initialized bool
		for {
			select {
			case <-ctx.Done():
				close(movingAverage)
				return
			case i := <-input:
				window[counter] = i
				counter++
				if counter == windowSize {
					counter = 0
					initialized = true
				}
				if !initialized {
					continue
				}
				var sum T
				for _, slot := range window {
					sum += slot
				}
				movingAverage <- sum / T(windowSize)
			}
		}
	}()
	return movingAverage
}

type Number interface {
	~float32 | ~float64 | ~int8 | ~int16 | ~int32 | ~int64 | ~int | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
