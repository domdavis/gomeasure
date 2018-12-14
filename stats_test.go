package gomeasure_test

import (
	"fmt"
	"strings"

	"github.com/domdavis/gomeasure"
)

func ExampleEmptyStats() {
	s := gomeasure.EmptyStats("my action")
	fmt.Println(s.Action, s.Min, s.Max, s.Total, s.Mean, s.Sigma, s.Samples)

	// Output:
	// my action 0s 0s 0s 0s 0s 0
}

func ExampleStats_String() {
	s := gomeasure.EmptyStats("example")
	split := strings.SplitAfter(s.String(), "]")
	fmt.Println(split[0])

	// Output:
	// Action: example [Min: 0s, Max: 0s, Mean: 0s, Sigma: 0s, Total: 0s, Samples: 0]
}
