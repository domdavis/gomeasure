package gomeasure

import (
	"fmt"
	"time"
)

// Stats holds the cumulative statistics for an action. All recorded samples for
// that action contribute to the stats figures. Once captures a set of stats 
// will not be updated by the action is represents.
type Stats struct {
	// Action is the name of the action the stats relate to.
	Action string `json:"action"`

	// Created is the timestamp when this set of stats was taken.
	Created time.Time `json:"created"`

	// Min is the shortest duration taken to perform the action. If no
	// samples have been recorded then the duration will be 0.
	Min time.Duration `json:"min"`

	// Max is the longest duration taken to perform the action. If no samples
	// have been recorded then the duration will be 0.
	Max time.Duration `json:"max"`

	// Mean is the average duration taken to perform the action. If no samples
	// have been recorded then the duration will be 0.
	Mean time.Duration `json:"mean"`

	// Sigma is the standard deviation of durations for this action. If no
	// samples have been recorded the duration will be 0, however, it is also
	// possible to have a variance of 0 if actions have be recorded, but with no
	// variance.
	Sigma time.Duration `json:"s"`

	// Total is the cumulative duration spent performing this action. If
	// no samples have been recorded then the duration will be 0.
	Total time.Duration `json:"total"`

	// Samples is the total number of timings that have been recorded against
	// this action.
	Samples int `json:"samples"`
}

const format = "Action: %s [Min: %s, Max: %s, Mean: %s, Sigma: %s, " +
	"Total: %s, Samples: %d], recorded at %s"

// EmptyStats returns a new, empty set of stats for the given action with all
// values zeroed, and the current time as the created time.
func EmptyStats(action string) Stats {
	return Stats{Action: action, Created: time.Now()}
}

// String returns the stats as a human readable string.
func (s Stats) String() string {
	return fmt.Sprintf(format, s.Action, s.Min, s.Max, s.Mean, s.Sigma,
		s.Total, s.Samples, s.Created)
}

