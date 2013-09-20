package timehistogram

import (
	"sort"
	"time"
)

// Item for generating a multi-series flot chart
type FlotBuilder struct {
	groups map[string]EventList
}

func (fb *FlotBuilder) AddPoints(group string, events ...Event) {
	if len(events) == 0 {
		return
	}

	// Initialize structure
	if fb.groups == nil {
		fb.groups = make(map[string]EventList)
	}

	// Add to Group
	fb.groups[group] = append(fb.groups[group], events...)
}

func (fb *FlotBuilder) Generate(resolution time.Duration, min, max time.Time) []FlotSeries {
	// Sort labels
	labels := make([]string, 0, len(fb.groups))
	for k := range fb.groups {
		labels = append(labels, k)
	}
	sort.Strings(labels)

	// Generate flot series
	data := []FlotSeries{}
	for _, l := range labels {
		data = append(data, FlotSeries{
			Label: l,
			Data:  Histogram(fb.groups[l], resolution, min, max),
		})
	}

	// Done
	return data
}

type FlotSeries struct {
	Label string    `json:"label"`
	Data  EventList `json:"data"`
}
