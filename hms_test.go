package hms

import (
	"fmt"
	"testing"
)

var input = []float64{
	0, 0.32, 1, 123.43, 999.001,
}

var output = []string{
	"00:00:00.000", "00:00:00.320", "00:00:01.000", "00:02:03.430", "00:16:39.001",
}

func TestHMS(t *testing.T) {
	for i, d := range input {
		ts := Seconds(d)
		
		if ts != output[i] {
			t.Errorf("'%v' != '%v'", ts, output[i])
		}

		fmt.Println(ts)
	}
}

func BenchmarkHMS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Seconds(1234567890.0123456789)
	}
}
