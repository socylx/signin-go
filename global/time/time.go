package time

import (
	"time"
)

type Format = string
type Duration = time.Duration
type Time = time.Time

var (
	cst              *time.Location
	DateTimeZeroTime Time                       = time.Time{}
	Now              func() Time                = time.Now
	Since            func(t Time) time.Duration = time.Since
)

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute

	CSTLayout Format = "2006-01-02 15:04:05" // CSTLayout China Standard Time Layout
	YYYYMMDD  Format = "2006-01-02"

	DateTimeZeroString string = "0000-00-00 00:00:00"
)
