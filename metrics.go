package gomeasure

import (
	"math"
	"sync"
	"time"
)

type metrics struct {
	min     time.Duration
	max     time.Duration
	mean    time.Duration
	total   time.Duration
	m2      time.Duration
	sigma   time.Duration
	samples int

	lock sync.RWMutex
}


// newMetrics returns a set of metrics based off the given timer. The sample
// size for the returned metrics will be 1.
func newMetrics(timer Timer) *metrics {
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

func (m *metrics) record(duration time.Duration) {
	m.lock.Lock()
	m.total += duration
	m.samples++

	d1 := duration - m.mean
	m.mean += m.average(d1)
	d2 := duration - m.mean
	m.m2 += d1 * d2
	m.sigma = m.sqrt(m.average(m.m2))

	if m.min == 0 || duration < m.min {
		m.min = duration
	}

	if m.max < duration {
		m.max = duration
	}

	m.lock.Unlock()
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

func (m *metrics) stats(action string) Stats {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return Stats{
		Action: action,
		Created: time.Now(),
		Min: m.min,
		Max: m.max,
		Mean: m.mean,
		Total: m.total,
		Sigma: m.sigma,
		Samples: m.samples,
	}
}
