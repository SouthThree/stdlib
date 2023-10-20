package wait

import "context"

type ConditionFunc func() (done bool, err error)

type ConditionWithContextFunc func(ctx context.Context) (done bool, err error)

func (cf ConditionFunc) WithContext() ConditionWithContextFunc {
	return func(ctx context.Context) (done bool, err error) {
		return cf()
	}
}
