package timehistogram

import "time"

func TimeMin(a, b time.Time) time.Time {
	if a.Before(b) {
		return a
	}
	return b
}

func TimeMax(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}
