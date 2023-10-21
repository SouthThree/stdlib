package wait

import (
	"time"

	stdruntime "github.com/SouthThree/stdlib/pkg/utils/runtime"
)

func JitterUntil(f func(), period time.Duration, jitterFactor float64, sliding bool, stopCh <-chan struct{}) {
	var ticker *time.Ticker
	for {
		select {
		case <-stopCh:
			return
		default:
		}

		jitterPeriod := period
		if jitterFactor > 0 {
			jitterPeriod = Jitter(period, jitterFactor)
		}

		if !sliding {
			ticker = time.NewTicker(jitterPeriod)
		}

		func() {
			defer stdruntime.HandlePanic()
			f()
		}()

		if sliding {
			ticker = time.NewTicker(jitterPeriod)
		}

		select {
		case <-stopCh:
			return
		case <-ticker.C:

		}
	}
}
