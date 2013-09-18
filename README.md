# Time Histogram generator

[![build status](https://secure.travis-ci.org/jamesharr/timehistogram.png)](http://travis-ci.org/jamesharr/timehistogram)

A silly little library for generating histograms off of time-ranged events that can overlap

```go
// Data types used (already defined)
type EventList []Event
type Event struct {
    Begin time.Time
    End   time.Time
    Data  int64
}

// Generate a histogram
hist := Histogram(
    myEventList,               // Some list of events
    time.Hour,                 // 1hr resolution on render
    time.Now(),                // Start time
    time.Now().Add(time.Hour), // End time
)

// Marshal to flot-friendly JSON
flotSeries := json.Marshal(hist)
```

For concrete, see the [tests](https://github.com/jamesharr/timehistogram/tree/master/histogram_test.go).

