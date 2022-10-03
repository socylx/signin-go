package time

import (
	"time"
)

var cst *time.Location

// CSTLayout China Standard Time Layout
const CSTLayout = "2006-01-02 15:04:05"

type Duration = time.Duration
type Time = time.Time

var Now func() Time = time.Now
var Since func(t Time) time.Duration = time.Since

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)
