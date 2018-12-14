# Measurement and Instrumentation package for Go

[![Build Status](https://travis-ci.org/domdavis/gomeasure.svg?branch=master)](https://travis-ci.org/domdavis/gomeasure)
[![Coverage Status](https://coveralls.io/repos/github/domdavis/gomeasure/badge.svg?branch=master)](https://coveralls.io/github/domdavis/gomeasure?branch=master)
[![Maintainability](https://api.codeclimate.com/v1/badges/01918f422f98828fa262/maintainability)](https://codeclimate.com/github/domdavis/gomeasure/maintainability)
[![](https://godoc.org/github.com/domdavis/gomeasure?status.svg)](http://godoc.org/github.com/domdavis/gomeasure)

`gomeasure` is designed to provide fine grained (μs) timings for arbitrary
actions which can be reported on during runtime. Actions are named, and metrics
for all actions of the same name can be queried at any time.

## Installation

```
go get github.com/domdavis/gomeasure
```

## Usage

Recording an action is simply a case of:

```
t := gomeasure.Action("my action")
// Do something
t.Stop()
```

A set of Stats can then be retrieved using the same name:

```
s := gomeasure.Report("my action")
```

The returned Stats provide timings as `time.Duration`, which provides ns
results, however, the overhead for starting and stopping a timer is ~750ns
making the results accurate to μs granularity.

A full report on all actions being measured is provided using:

```
r := gomeasure.Snapshot()
```

## Actions

Segregated sets of actions can be created if required:

```
a := gomeasure.NewActions()
t := a.TimerFor("my action")
t.Stop()
fmt.Println(a.Snapshot())
```
