package gomeasure

type listener struct {
	in      chan Timer
	out     chan Metrics
	lookup  chan string
	metrics map[string]Metrics
}

func (l *listener) read(name string) Metrics {
	l.lookup <- name
	return <-l.out
}

func (l *listener) write(timer Timer) {
	l.in <- timer
}

func (l *listener) listen() {
	l.in = make(chan Timer)
	l.out = make(chan Metrics)
	l.lookup = make(chan string)
	l.metrics = map[string]Metrics{}

	go func() {
		for {
			select {
			case action := <-l.in:
				l.record(action)
			case name := <-l.lookup:
				l.respond(name)
			}
		}
	}()
}

func (l *listener) record(timer Timer) {
	if m := l.metrics[timer.Action()]; m == nil {
		l.metrics[timer.Action()] = NewMetrics(timer)
	} else {
		m.Record(timer.Duration())
	}
}

func (l *listener) respond(name string) {
	if m := l.metrics[name]; m != nil {
		l.out <- m.Copy()
	} else {
		l.out <- EmptyMetrics()
	}
}
