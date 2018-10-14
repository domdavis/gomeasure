package gomeasure

import (
	"fmt"
	"math"
	"time"
)

// Metrics are used to record the cumulative statistics for a set of actions.
type Metrics interface {
	// Min returns the shortest duration required to perform an action in this
	// set. If no actions have been recorded then the duration will be 0.
	Min() time.Duration

	// Max returns the longest duration required to perform an action in this
	// set. If no actions have been recorded then the duration will be 0.
	Max() time.Duration

	// Mean is the average duration required to perform actions in this set. If
	// no actions have been recorded then the duration will be 0.
	Mean() time.Duration

	// Total is the cumulative duration spent performing actions in this set. If
	// no actions have been recorded then the duration will be 0.
	Total() time.Duration

	// Sigma is the standard deviation of durations for actions in this set. If
	// no actions have been recorded the duration will be 0, however, it is also
	// possible to have a variance of 0 if actions have be recorded, but with no
	Sigma() time.Duration

	// Samples is the total number of actions that have been recorded against
	// this set.
	Samples() int

	// Record a timed action against this set of statistics, adding its data to
	// the set of actions. Metrics are recorded using Welford's Algorithm as
	// described on the Wikipedia page for Algorithms for Calculating Variance
	// (https://en.wikipedia.org/wiki/Algorithms_for_calculating_variance).
	Record(duration time.Duration)

	// Copy this Metrics type into a new Metrics type occupying a different
	// memory address.
	Copy() Metrics

	// String returns a human readable version of a Metrics type.
	String() string
}

type metrics struct {
	min     time.Duration
	max     time.Duration
	mean    time.Duration
	total   time.Duration
	m2      time.Duration
	sigma   time.Duration
	samples int
}

const format = "Min: %s, Max: %s, Mean: %s, Total: %s, Sigma: %s, Samples: %d"

// NewMetrics returns a set of metrics based off the given timer. The sample
// size for the returned metrics will be 1.
func NewMetrics(timer Timer) Metrics {
	duration := timer.Duration()
	return &metrics{
		min:     duration,
		max:     duration,
		mean:    duration,
		total:   duration,
		m2:      0,
		sigma:   0,
		samples: 1,
	}
}

// EmptyMetrics returns a new, empty set of metrics with no samples recorded.
func EmptyMetrics() Metrics {
	return &metrics{}
}

func (m *metrics) Min() time.Duration {
	return m.min
}

func (m *metrics) Max() time.Duration {
	return m.max
}

func (m *metrics) Mean() time.Duration {
	return m.mean
}

func (m *metrics) Total() time.Duration {
	return m.total
}

func (m *metrics) Sigma() time.Duration {
	return m.sqrt(m.average(m.m2))
}

func (m *metrics) Samples() int {
	return m.samples
}

func (m *metrics) Record(duration time.Duration) {
	m.total += duration
	m.samples++

	d1 := duration - m.mean
	m.mean += m.average(d1)
	d2 := duration - m.mean
	m.m2 += d1 * d2

	if m.min == 0 || duration < m.min {
		m.min = duration
	}

	if m.max < duration {
		m.max = duration
	}
}

func (m *metrics) Copy() Metrics {
	c := *m
	return &c
}

func (m *metrics) String() string {
	return fmt.Sprintf(
		format, m.min, m.max, m.mean, m.total, m.Sigma(), m.samples)
}

func (m *metrics) average(duration time.Duration) time.Duration {
	if m.samples == 0 {
		return 0
	}

	return duration / time.Duration(m.samples)
}

func (m *metrics) sqrt(duration time.Duration) time.Duration {
	return time.Duration(math.Sqrt(float64(duration)))
}
