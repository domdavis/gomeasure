package gomeasure

import "time"

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

	dead    bool
	actions *actions
}

func (t *timer) Stop() {
	if t.dead {
		return
	}

	t.dead = true
	t.end = time.Now()
	t.actions.record(t)
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
