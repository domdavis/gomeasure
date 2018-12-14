package gomeasure

import (
	"testing"
	"time"
)

func TestMetrics(t *testing.T) {
	m := &metrics{}

	// Check we don't get divide by zero error
	m.average(0)
	valid := m.mean == 0

	// record some stats
	m.record(9)
	m.record(2)
	m.record(5)
	m.record(4)
	m.record(12)

	// Check what we've recorded is what we expect
	valid = valid && m.min == 2
	valid = valid && m.max == 12
	valid = valid && m.mean == 7
	valid = valid && m.total == 32
	valid = valid && m.sigma == 3
	valid = valid && m.samples == 5

	if !valid {
		t.Fail()
	}
}

func BenchmarkMetrics_record(b *testing.B) {
	m := &metrics{}
	for n := 0; n < b.N; n++ {
		m.record(time.Duration(n))
	}
}
