package gomeasure

import (
	"fmt"
	"testing"
	"time"
)

func ExampleMetrics_record() {
	m := &metrics{}
	m.record(9)
	m.record(2)
	m.record(5)
	m.record(4)
	m.record(12)
	fmt.Printf("Min: %s\n", m.min)
	fmt.Printf("Max: %s\n", m.max)
	fmt.Printf("Mean: %s\n", m.mean)
	fmt.Printf("Total: %s\n", m.total)
	fmt.Printf("Sigma: %s\n", m.sigma)
	fmt.Printf("Samples: %d\n", m.samples)

	// Output:
	// Min: 2ns
	// Max: 12ns
	// Mean: 7ns
	// Total: 32ns
	// Sigma: 3ns
	// Samples: 5
}

func ExampleMetrics_average() {
	m := &metrics{}

	// Zero samples, should not cause a divide by zero error
	m.average(0)
	fmt.Println(m.mean)

	// Output:
	// 0s
}

func BenchmarkMetrics_record(b *testing.B) {
	m := &metrics{}
	for n := 0; n < b.N; n++ {
		m.record(time.Duration(n))
	}
}
