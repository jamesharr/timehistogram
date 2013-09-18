package timehistogram

import (
	"encoding/json"
	"fmt"
	"github.com/bmizerany/assert"
	"testing"
	"time"
)

func at(t string) time.Time {
	rv, err := time.Parse(time.Kitchen, t)
	if err != nil {
		panic(err)
	}
	return rv
}

func TestHistogram(t *testing.T) {

	events := EventList{
		{at("2:10PM"), at("2:15PM"), 5},
		{at("2:11PM"), at("2:16PM"), 5},
	}

	// Full histogram
	hist := Histogram(events, time.Minute, at("2:10PM"), at("2:16PM"))
	exp := EventList{
		{at("2:10PM"), at("2:11PM"), 5},
		{at("2:11PM"), at("2:12PM"), 10},
		{at("2:12PM"), at("2:13PM"), 10},
		{at("2:13PM"), at("2:14PM"), 10},
		{at("2:14PM"), at("2:15PM"), 10},
		{at("2:15PM"), at("2:16PM"), 5},
	}
	assert.Equal(t, exp, hist)

	// Small slice
	hist = Histogram(events, time.Minute, at("2:13PM"), at("2:15PM"))
	exp = EventList{
		{at("2:13PM"), at("2:14PM"), 10},
		{at("2:14PM"), at("2:15PM"), 10},
	}
	assert.Equal(t, exp, hist)

	// Wide
	hist = Histogram(events, time.Hour, at("2:00PM"), at("3:00PM"))
	exp = EventList{
		{at("2:00PM"), at("3:00PM"), 10},
	}
	assert.Equal(t, exp, hist)

}

func TestMarshal(t *testing.T) {
	t1 := at("2:00PM")
	t2 := time.Now().Truncate(time.Second)
	data := TimeSeries{
		{t1, 11},
		{t2, 12},
	}
	exp := fmt.Sprintf("[[%d,11],[%d,12]]", t1.Unix()*1000, t2.Unix()*1000)
	m, err := json.Marshal(data)
	assert.Equal(t, nil, err)
	assert.Equal(t, exp, string(m))
}
