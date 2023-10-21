package wait

import (
	"context"
	"time"
)

func PollUntilContextCancel(ctx context.Context, interval time.Duration, immediate bool, condition ConditionWithContextFunc) error {
	return loopConditionUntilContext(ctx, interval, immediate, condition)
}

func PollUntilContextTimeout(ctx context.Context, interval, timeout time.Duration, immediate bool, condition ConditionWithContextFunc) error {
	timeoutContext, timeoutCancel := context.WithTimeout(ctx, timeout)
	defer timeoutCancel()
	return loopConditionUntilContext(timeoutContext, interval, immediate, condition)
}
