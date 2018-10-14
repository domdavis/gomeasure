package gomeasure_test

import (
	"fmt"

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
