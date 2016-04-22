package gtime

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func logError(result, expected interface{}) {
	fmt.Errorf("Mismatched values: Expect %v. Got %v", expected, result)
}

var ftosTable = []struct {
	n        float64
	expected time.Duration
}{
	{0.1, time.Duration(100) * time.Millisecond},
	{1.0, time.Duration(1) * time.Second},
	{10.0, time.Duration(10) * time.Second},
	{60.0, time.Duration(60) * time.Second},
}

func TestFtos(t *testing.T) {
	for _, tt := range ftosTable {
		result := Ftos(tt.n)
		if !(result >= tt.expected-FloatTime || result <= tt.expected-FloatTime) {
			logError(result, tt.expected)
			t.Error("Result isn't accurate.")
		}
	}
}

func TestStof(t *testing.T) {
	for _, tt := range ftosTable {
		result := Stof(tt.expected)
		if !(result >= tt.n-nano || result <= tt.n+nano) {
			logError(result, tt.expected)
			t.Error("Result isn't accurate.")
		}
	}
}

var testTable = []struct {
	dr  time.Duration
	hr  float64
	min float64
	ns  int64
	s   float64
	st  string
}{
	{
		time.Duration(24) * time.Hour,
		24.0,
		1440.0,
		8.64e+13,
		86400.0,
		"24h0m0s",
	},
}

func TestHour(t *testing.T) {
	for _, tt := range testTable {
		h := Hour(tt.hr)
		if h.Duration() != tt.dr {
			logError(h.Duration(), tt.dr)
		}
		if h.Hours() != tt.hr {
			logError(h.Hours(), tt.hr)
		}
		if h.Minutes() != tt.min {
			logError(h.Minutes(), tt.min)
		}
		if h.Nanoseconds() != tt.ns {
			logError(h.Nanoseconds(), tt.ns)
		}
		if h.Seconds() != tt.s {
			logError(h.Seconds(), tt.s)
		}
		if h.String() != tt.st {
			logError(h.String(), tt.st)
		}
	}
}

func TestMinute(t *testing.T) {
	for _, tt := range testTable {
		m := Minute(tt.min)
		if m.Duration() != tt.dr {
			logError(m.Duration(), tt.dr)
		}
		if m.Hours() != tt.hr {
			logError(m.Hours(), tt.hr)
		}
		if m.Minutes() != tt.min {
			logError(m.Minutes(), tt.min)
		}
		if m.Nanoseconds() != tt.ns {
			logError(m.Nanoseconds(), tt.ns)
		}
		if m.Seconds() != tt.s {
			logError(m.Seconds(), tt.s)
		}
		if m.String() != tt.st {
			logError(m.String(), tt.st)
		}
	}
}

func TestNanosecond(t *testing.T) {
	for _, tt := range testTable {
		ns := Nanosecond(tt.ns)
		if ns.Duration() != tt.dr {
			logError(ns.Duration(), tt.dr)
		}
		if ns.Hours() != tt.hr {
			logError(ns.Hours(), tt.hr)
		}
		if ns.Minutes() != tt.min {
			logError(ns.Minutes(), tt.min)
		}
		if ns.Nanoseconds() != tt.ns {
			logError(ns.Nanoseconds(), tt.ns)
		}
		if ns.Seconds() != tt.s {
			logError(ns.Seconds(), tt.s)
		}
		if ns.String() != tt.st {
			logError(ns.String(), tt.st)
		}
	}
}

func TestSecond(t *testing.T) {
	for _, tt := range testTable {
		s := Second(tt.s)
		if s.Duration() != tt.dr {
			logError(s.Duration(), tt.dr)
		}
		if s.Hours() != tt.hr {
			logError(s.Hours(), tt.hr)
		}
		if s.Minutes() != tt.min {
			logError(s.Minutes(), tt.min)
		}
		if s.Nanoseconds() != tt.ns {
			logError(s.Nanoseconds(), tt.ns)
		}
		if s.Seconds() != tt.s {
			logError(s.Seconds(), tt.s)
		}
		if s.String() != tt.st {
			logError(s.String(), tt.st)
		}
	}
}

var durationTable = []struct {
	n        Duration
	expected float64
}{
	{Second(1.0), 1.0},
	{Second(0.2), 0.2},
	{Minute(60.0), 3600.0},
	{time.Duration(1) * time.Second, 1.0},
	{time.Duration(60) * time.Minute, 3600.0},
	{time.Duration(200) * time.Millisecond, 0.2},
}

func TestDuration(t *testing.T) {
	handler := func(d Duration) (s float64) {
		s = d.Seconds()
		return
	}
	for _, tt := range durationTable {
		result := handler(tt.n)
		if result != tt.expected {
			logError(result, tt.expected)
			t.Error(errors.New("Mismatched!"))
		}
	}
}
