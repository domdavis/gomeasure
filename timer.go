package gomeasure

import (
	"sync"
	"time"
)

// Timer is the interface implemented by types that can time the duration of an
// action. Timers are started on creation and stopped on a call to Stop.
type Timer interface {
	// Stop this Timer, recording the duration the timer was active for. Only
	// the first call to Stop will record the timers duration. Further calls to
	// Stop quietly fail.
	Stop()

	// Action returns the action this timer is timing.
	Action() string

	// Duration returns the recorded time this timer was active for. If the
	// timer is still running Duration will return the current duration the
	// timer has been running for.
	Duration() time.Duration
}

type timer struct {
	action string
	start  time.Time
	end    time.Time

	dead     bool
	listener *listener
}

var l = &listener{}
var once sync.Once

// Action returns a running Timer to measure an action. Multiple Timers can be
// can be created for any given action, representing multiple instances of that
// action. These timers can be concurrent.
func Action(action string) Timer {
	once.Do(func() {
		l.listen()
	})

	return &timer{action: action, start: time.Now(), listener: l}
}

// Report the Metrics for an action.
func Report(action string) Metrics {
	return l.read(action)
}

func (t *timer) Stop() {
	if t.dead {
		return
	}

	t.dead = true
	t.end = time.Now()
	t.listener.write(t)
}

func (t *timer) Action() string {
	return t.action
}

func (t *timer) Duration() time.Duration {
	if t.end.Before(t.start) {
		return time.Now().Sub(t.start)
	}

	return t.end.Sub(t.start)
}
