package gomeasure

import (
	"sync"
	"time"
)

// Actions holds metrics on a set of actions.
type Actions interface{

	// TimerFor returns a running Timer to measure the named action. Multiple
	// Timers can be can be created for any given action, representing multiple
	// instances of that action. These timers can be concurrent. When the timer
	// is stopped the metrics for that timer will be added to this set of
	// actions.
	TimerFor(action string) Timer

	// Report the Stats for the named action. If the action has not been
	// recorded in this set then a set of empty Stats will be returned.
	Report(action string) Stats

	// Snapshot records a snapshot of Stats for all actions in this set.
	Snapshot() []Stats
}

type actions struct {
	lock sync.RWMutex
	metrics map[string]*metrics
}

var defaultActions = NewActions()

// Action returns a running Timer to measure an action. Multiple Timers can be
// can be created for any given action, representing multiple instances of that
// action. These timers can be concurrent.
func Action(action string) Timer {
	return defaultActions.TimerFor(action)
}

// Report the Stats for the named action. If the action has not been recorded
// then a set of empty Stats will be returned.
func Report(action string) Stats {
	return defaultActions.Report(action)
}

// Snapshot records a snapshot of Stats for all recorded actions.
func Snapshot() []Stats {
	return defaultActions.Snapshot()
}

// NewActions returns a new, empty set of actions.
func NewActions() Actions {
	return &actions{metrics: map[string]*metrics{}}
}

// TimerFor returns a running Timer to measure the named action. Multiple Timers
// can be can be created for any given action, representing multiple instances
// of that action. These timers can be concurrent.
func (s *actions) TimerFor(action string) Timer {
	return &timer{action: action, start: time.Now(), actions: s}
}

// Report the Stats for the named action. If the action has not been recorded
// in this set then a set of empty Stats will be returned.
func (s *actions) Report(action string) Stats {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if m, ok := s.metrics[action]; ok {
		return m.stats(action)
	}

	return EmptyStats(action)
}

// Snapshot records a snapshot of Stats for all recorded actions is this set.
func (s *actions) Snapshot() []Stats {
	var i int
	s.lock.RLock()
	defer s.lock.RUnlock()


	m := make([]Stats, len(s.metrics))

	for k, v := range s.metrics {
		m[i] = v.stats(k)
		i++
	}

	return m
}

func (s *actions) record(timer Timer) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if m := s.metrics[timer.Action()]; m == nil {
		s.metrics[timer.Action()] = newMetrics(timer)
	} else {
		m.record(timer.Duration())
	}
}
