package gomeasure_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/domdavis/gomeasure"
)

func ExampleEmptyMetrics() {
	m := gomeasure.EmptyMetrics()
	fmt.Println(m.Min(), m.Max(), m.Total(), m.Mean(), m.Sigma(), m.Samples())

	// Output: 0s 0s 0s 0s 0s 0
}

func ExampleMetrics_String() {
	m := gomeasure.EmptyMetrics()
	fmt.Println(m.String())

	// Output:
	// Min: 0s, Max: 0s, Mean: 0s, Total: 0s, Sigma: 0s, Samples: 0
}

func ExampleMetrics_Record() {
	m := gomeasure.EmptyMetrics()
	m.Record(9)
	m.Record(2)
	m.Record(5)
	m.Record(4)
	m.Record(12)
	fmt.Println(m.String())

	// Output:
	// Min: 2ns, Max: 12ns, Mean: 7ns, Total: 32ns, Sigma: 3ns, Samples: 5
}

func ExampleMetrics_Copy() {
	m := gomeasure.EmptyMetrics()
	m.Record(1)
	c := m.Copy()
	c.Record(2)

	fmt.Println(m.Samples(), c.Samples())

	// Output:
	// 1 2
}

func BenchmarkMetrics_Record(b *testing.B) {
	m := gomeasure.EmptyMetrics()
	for n := 0; n < b.N; n++ {
		m.Record(time.Duration(n))
	}
}
