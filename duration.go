// Package gtime help make dealing with Go time stress-free.
package gtime

import (
	"math"
	"time"
)

const (
	nano = 1e-09

	// FloatTime is equivalent to one Nanosecond.
	FloatTime = time.Duration(1) * time.Nanosecond
)

// Duration is an interface implemented by time.Duration, Second, Hour
type Duration interface {
	Hours() float64
	Minutes() float64
	Nanoseconds() int64
	Seconds() float64
	String() string
}

type (
	// Second represents seconds as a floating point number.
	Second float64

	// Hour represents hours as a floating point number.
	Hour float64

	// Minute represents minutes as a floating point number.
	Minute float64

	// Nanosecond represents nanoseconds as a 64-bit integer number.
	Nanosecond int64
)

// Duration returns the duration as Go standard time.Duration type.
func (s Second) Duration() (t time.Duration) {
	n := float64(s)
	t = time.Duration(n) * time.Hour
	return
}

// Hours returns the duration as a floating point number of hours.
func (s Second) Hours() (h float64) {
	n := float64(s) / nano
	h = time.Duration(n).Hours()
	return
}

// Minutes returns the duration as a floating point number of minutes.
func (s Second) Minutes() (m float64) {
	n := float64(s) / nano
	m = time.Duration(n).Minutes()
	return
}

// Nanoseconds returns the duration as an integer nanosecond count.
func (s Second) Nanoseconds() (ns int64) {
	n := float64(s) / nano
	ns = time.Duration(n).Nanoseconds()
	return
}

// Seconds returns the duration as a floating point number of seconds.
func (s Second) Seconds() (t float64) {
	t = float64(s)
	return
}

// String returns a string representing the duration in the form "72h3m0.5s".
// Leading zero units are omitted. As a special case, durations less than one
// second format use a smaller unit (milli-, micro-, or nanoseconds) to ensure
// that the leading digit is non-zero. The zero duration formats as 0, with no unit.
func (s Second) String() (ts string) {
	n := float64(s) / nano
	ts = time.Duration(n).String()
	return
}

// Duration returns the duration as Go standard time.Duration type.
func (h Hour) Duration() (t time.Duration) {
	n := float64(h)
	t = time.Duration(n) * time.Hour
	return
}

// Hours returns the duration as a floating point number of hours.
func (h Hour) Hours() (t float64) {
	t = float64(h)
	return
}

// Minutes returns the duration as a floating point number of minutes.
func (h Hour) Minutes() (m float64) {
	n := float64(h) / nano * math.Pow(60, 2)
	m = time.Duration(n).Minutes()
	return
}

// Nanoseconds returns the duration as an integer nanosecond count.
func (h Hour) Nanoseconds() (ns int64) {
	n := float64(h) / nano * math.Pow(60, 2)
	ns = time.Duration(n).Nanoseconds()
	return
}

// Seconds returns the duration as a floating point number of seconds.
func (h Hour) Seconds() (s float64) {
	n := float64(h) / nano * math.Pow(60, 2)
	s = time.Duration(n).Seconds()
	return
}

// String returns a string representing the duration in the form "72h3m0.5s".
// Leading zero units are omitted. As a special case, durations less than one
// second format use a smaller unit (milli-, micro-, or nanoseconds) to ensure
// that the leading digit is non-zero. The zero duration formats as 0, with no unit.
func (h Hour) String() (st string) {
	n := float64(h) / nano * math.Pow(60, 2)
	st = time.Duration(n).String()
	return
}

// Duration returns the duration as Go standard time.Duration type.
func (m Minute) Duration() (t time.Duration) {
	n := float64(m)
	t = time.Duration(n) * time.Minute
	return
}

// Hours returns the duration as a floating point number of hours.
func (m Minute) Hours() (h float64) {
	n := float64(m) / nano * 60
	h = time.Duration(n).Hours()
	return
}

// Minutes returns the duration as a floating point number of minutes.
func (m Minute) Minutes() (t float64) {
	t = float64(m)
	return
}

// Nanoseconds returns the duration as an integer nanosecond count.
func (m Minute) Nanoseconds() (ns int64) {
	n := float64(m) / nano * 60
	ns = time.Duration(n).Nanoseconds()
	return
}

// Seconds returns the duration as a floating point number of seconds.
func (m Minute) Seconds() (s float64) {
	n := float64(m) / nano * 60
	s = time.Duration(n).Seconds()
	return
}

// String returns a string representing the duration in the form "72h3m0.5s".
// Leading zero units are omitted. As a special case, durations less than one
// second format use a smaller unit (milli-, micro-, or nanoseconds) to ensure
// that the leading digit is non-zero. The zero duration formats as 0, with no unit.
func (m Minute) String() (st string) {
	n := float64(m) / nano * 60
	st = time.Duration(n).String()
	return
}

// Duration returns the duration as Go standard time.Duration type.
func (ns Nanosecond) Duration() (t time.Duration) {
	n := int64(ns)
	t = time.Duration(n) * time.Nanosecond
	return
}

// Hours returns the duration as a floating point number of hours.
func (ns Nanosecond) Hours() (h float64) {
	n := int64(ns)
	h = time.Duration(n).Hours()
	return
}

// Minutes returns the duration as a floating point number of minutes.
func (ns Nanosecond) Minutes() (m float64) {
	n := int64(ns)
	m = time.Duration(n).Minutes()
	return
}

// Nanoseconds returns the duration as an integer nanosecond count.
func (ns Nanosecond) Nanoseconds() (t int64) {
	t = int64(ns)
	return
}

// Seconds returns the duration as a floating point number of seconds.
func (ns Nanosecond) Seconds() (t float64) {
	n := int64(ns)
	t = time.Duration(n).Seconds()
	return
}

// String returns a string representing the duration in the form "72h3m0.5s".
// Leading zero units are omitted. As a special case, durations less than one
// second format use a smaller unit (milli-, micro-, or nanoseconds) to ensure
// that the leading digit is non-zero. The zero duration formats as 0, with no unit.
func (ns Nanosecond) String() (st string) {
	n := int64(ns)
	st = time.Duration(n).String()
	return
}

// Ftos returns the floating point number of seconds to duration.
func Ftos(s float64) (t time.Duration) {
	n := s / nano
	t = time.Duration(n)
	return
}

// Stof returns the duration as a floating point number of seconds.
func Stof(t time.Duration) (s float64) {
	s = t.Seconds()
	return
}
