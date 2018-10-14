package gomeasure_test

import (
	"fmt"

	"github.com/domdavis/gomeasure"
)

func ExampleAction() {
	// Action timers are arbitrary and can be started anywhere. The resultant
	// timer needs to be stopped once the action has completed.
	l := gomeasure.Action("example")
	defer l.Stop()

	for i := 0; i < 100; i++ {
		// Multiple timers of the same name are allowed, and their results are
		// added to the metrics set to give an overview of how the action is
		// performing over multiple runs.
		s := gomeasure.Action("sprintf")
		_ = fmt.Sprintf("String %d", i)
		s.Stop()
	}

	// A report on a named action can be pulled, and the metrics collected
	// analysed.
	r := gomeasure.Report("sprintf")
	fmt.Println(r.Samples())

	// An action timer that is still running will not have its metrics added to
	// the set.
	r = gomeasure.Report("example")
	fmt.Println(r.Total())

	// Output:
	// 100
	// 0s

}

func ExampleTimer_Stop() {
	// Calling Stop on an action more than once has no effect. The second call
	// is simply ignored
	t := gomeasure.Action("action")
	t.Stop()
	t.Stop()

	// Output:
}

func ExampleTimer_Duration() {
	t := gomeasure.Action("action")
	defer t.Stop()

	// A running time will return the current run duration
	mid := t.Duration()
	t.Stop()
	end := t.Duration()

	fmt.Println(mid != end)
	// Output: true
}
