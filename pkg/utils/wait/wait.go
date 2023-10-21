package wait

import (
	"context"
	"math/rand"
	"time"
)

var NeverStop <-chan struct{} = make(<-chan struct{})

func Forever(f func(), period time.Duration) {
	Until(f, period, NeverStop)
}

func Until(f func(), period time.Duration, stopCh <-chan struct{}) {
	JitterUntil(f, period, 0.0, true, stopCh)
}

func Jitter(duration time.Duration, maxFactor float64) time.Duration {
	if maxFactor <= 0.0 {
		maxFactor = 1.0
	}
	duration = duration + time.Duration(float64(duration)*rand.Float64()*maxFactor)
	return duration
}

type ConditionFunc func() (done bool, err error)

type ConditionWithContextFunc func(ctx context.Context) (done bool, err error)

func (cf ConditionFunc) WithContext() ConditionWithContextFunc {
	return func(ctx context.Context) (done bool, err error) {
		return cf()
	}
}
