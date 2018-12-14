package gomeasure_test

import (
	"fmt"
	"testing"

	"github.com/domdavis/gomeasure"
)

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

func BenchmarkTime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		t := gomeasure.Action("action")
		t.Stop()
	}
}
