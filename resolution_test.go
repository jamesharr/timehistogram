package timehistogram_test

import "testing"
import (
	"github.com/bmizerany/assert"
	thg "github.com/jamesharr/timehistogram"
	"time"
)

func TestResolutionSelect(t *testing.T) {
	t_0 := time.Unix(0, 0)
	t_1d := time.Unix(24*3600, 0)

	d_5min := 5 * time.Minute
	d_10min := 10 * time.Minute

	res := thg.STANDARD.Select(t_0, t_1d, 100)
	assert.Equal(t, d_10min, res)

	res = thg.STANDARD.Select(t_0, t_1d, 180)
	assert.Equal(t, d_5min, res)
}
