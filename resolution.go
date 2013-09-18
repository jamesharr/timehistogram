package timehistogram

import (
	"sort"
	"time"
)

// ResolutionList - A list of Histogram resolutions
type ResolutionList []time.Duration

// Select a sane resolution that provides a minimum amount of data points between start & end.
//
// Note: ResolutionList must be sorted in DESCENDING order.
func (resList ResolutionList) Select(start, end time.Time, points int64) time.Duration {
	roughResolution := end.Sub(start) / time.Duration(points)
	n := sort.Search(len(resList), func(i int) bool {
		return resList[i] < roughResolution
	})

	if n >= len(resList) {
		n = len(resList) - 1
	}
	return resList[n]
}

// STANDARD - a sane list of time resolutions
var STANDARD = ResolutionList{
	24 * 365 * time.Hour,
	24 * 28 * time.Hour,
	24 * 7 * time.Hour,

	24 * time.Hour,
	12 * time.Hour,
	6 * time.Hour,
	4 * time.Hour,
	1 * time.Hour,

	30 * time.Minute,
	20 * time.Minute,
	15 * time.Minute,
	10 * time.Minute,
	5 * time.Minute,
	1 * time.Minute,

	30 * time.Second,
	20 * time.Second,
	15 * time.Second,
	10 * time.Second,
	5 * time.Second,
	1 * time.Second,
}
