// Package gomeasure allows for arbitrary actions in code to be timed. Actions
// are grouped into sets by name, and metrics are provided for each named set.
//
// Timing actions is simply a case of starting and stopping a time:
//
//    t := gomeasure.Action("my action")
//    // do something
//    t.Stop()
//
// Metrics for a set of actions are retrieved using the same name passed to the
// Time function:
//
//    m := gomeasure.Report("my action")
//
// The gomeasure package is thread safe, and multiple actions of the same name
// can be timed concurrently. While the metrics report to nanosecond precision,
// this package isn't meant to replace tools such as pprof as the timers
// themselves introduce an overhead (~750ns) on each call. Instead it's meant
// for longer running actions, like servicing http requests to give an overview
// of how aspects of an application are performing.
package gomeasure
