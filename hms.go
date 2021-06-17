// hms, https://tpkn.me
package hms

import (
	"fmt"
	"math"
)

type timestamp struct {
	hours   float64
	minutes float64
	seconds float64
	msec    float64
}

func (t timestamp) stringify() string {
	return fmt.Sprintf("%02v:%02v:%02v.%03v", t.hours, t.minutes, t.seconds, t.msec)
}

// Seconds converts seconds to pretty 'HH:MM:SS.MSEC' style
func Seconds(duration float64) string {
	t := timestamp{
		math.Floor(duration / 3600),
		math.Floor(math.Mod(duration / 60, float64(60))),
		math.Floor(math.Mod(duration, float64(60))),
		math.Floor(math.Mod(duration * 1000, 1000)),
	}
	
	return t.stringify()
}
