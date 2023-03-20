package filter

import "time"

func SET_TIME_PROVIDER(provider func() time.Time) {
	timeProvider = provider
}
