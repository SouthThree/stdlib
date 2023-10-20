package wait

import (
	"context"
	"log"
	"time"
)

// loopConditionUntilContext execute the condition in period until ctx is cancelled, condition returns true or condition
//
//	@Description: execute the condition in period until ctx is cancelled, condition returns true or condition return an
//	              error. If immediate is true, condition is executed immediately, otherwise it is executed after period time.
//	@param ctx
//	@param period
//	@param immediate
//	@param condition
//	@return error return by condition or ctx.Err()
func loopConditionUntilContext(ctx context.Context, period time.Duration, immediate bool, condition ConditionWithContextFunc) error {
	recoverCondition := func() (bool, error) {
		defer func() {
			if e := recover(); e != nil {
				log.Println("panic: ", e)
			}
		}()
		return condition(ctx)
	}

	ticker := time.NewTicker(period)
	defer ticker.Stop()

	if immediate {
		if ok, err := recoverCondition(); err != nil || ok {
			return err
		}
	} else {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			ticker.Reset(period)
		}
	}

	for {
		select {
		case <-ticker.C:
			ticker.Reset(period)
		case <-ctx.Done():
			return ctx.Err()
		}

		if ok, err := recoverCondition(); err != nil || ok {
			return err
		}
	}
}
