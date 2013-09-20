package timehistogram

import (
	"encoding/json"
	"math"
	"sort"
	"time"
)

// WHY DOES THIS NOT EXIST?!?!
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// Event object
type Event struct {
	Begin time.Time
	End   time.Time
	Data  int64
}

func (a Event) Overlaps(b Event) bool {
	return a.Begin.Before(b.End) && b.Begin.Before(a.End)
}

// EventList - Sort()able, Range()able list of Events
type EventList []Event

func (events EventList) Sort() {
	sort.Sort(events)
}

// MarshalJSON spits out data that's useful to JS apps. Meant to be fed to flot
//
// JSON Structure: [[e1.Begin,e1.Data], [e2.Begin,e2.Data], ...]
// Time Format: time.Unix()*1000 [milliseconds since epoch]
func (events EventList) MarshalJSON() ([]byte, error) {
	m := make([][2]int64, len(events))
	for i, pt := range events {
		m[i][0] = pt.Begin.Unix() * 1000
		m[i][1] = pt.Data
	}
	return json.Marshal(m)
}

func (events EventList) Len() int {
	return len(events)
}

func (events EventList) Less(i, j int) bool {
	return events[i].Begin.Before(events[j].Begin)
}

func (events EventList) Swap(i, j int) {
	events[i], events[j] = events[j], events[i]
}

func (events EventList) Range(begin, end time.Time) EventList {
	start_idx := sort.Search(len(events), func(i int) bool {
		return begin.Before(events[i].End)
	})
	end_idx := sort.Search(len(events), func(i int) bool {
		return end.Before(events[i].End)
	})

	return events[start_idx:end_idx]
}

// Histogram renders a fixed interval histogram out of a list of events
func Histogram(events EventList, resolution time.Duration, begin, end time.Time) EventList {
	// Establish a list of all times mentioned
	all_times := make(map[time.Time]bool)
	for _, evt := range events {
		all_times[evt.Begin] = true
		all_times[evt.End] = true
	}

	// Make a unique list of all times mentioned
	reduced := make(EventList, 0, len(all_times))

	// Put in list, sort
	for k := range all_times {
		// Set begin time
		reduced = append(reduced, Event{Begin: k})
	}
	reduced.Sort()

	// Set an end time, chop off last one
	for i := 0; i < len(reduced)-1; i++ {
		reduced[i].End = reduced[i+1].Begin
	}
	reduced = reduced[0 : len(reduced)-1]

	// Run through times again, add into reduced/flat set
	for _, evt := range events {
		relevant := reduced.Range(evt.Begin, evt.End)
		for i := range relevant {
			relevant[i].Data += evt.Data
		}
	}

	// Setup render plot
	bucketCount := int(math.Ceil(end.Sub(begin).Seconds() / resolution.Seconds()))
	buckets := make(EventList, bucketCount)
	for i := range buckets {
		buckets[i].Begin = begin.Add(resolution * time.Duration(i))
		buckets[i].End = begin.Add(resolution * time.Duration(i+1))
	}

	// Render
	i := 0 // For reduced
	j := 0 // For buckets
	for i < len(reduced) && j < len(buckets) {
		// Update if they overlap
		if reduced[i].Overlaps(buckets[j]) {
			buckets[j].Data = max(buckets[j].Data, reduced[i].Data)
		}

		// Advance the least
		if buckets[j].End.Before(reduced[i].End) {
			j++
		} else {
			i++
		}
	}

	return buckets
}
