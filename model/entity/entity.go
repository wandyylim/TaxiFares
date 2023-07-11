package entity

import "time"

type Record struct {
	Time     time.Time
	Distance float64
	Diff     float64
}
