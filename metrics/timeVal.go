package main

import "time"

var upTimeMetrics *upTimeVar

type upTimeVar struct {
	value time.Time
}

func (u *upTimeVar) Set(date time.Time) {
	u.value = date
}

func (u *upTimeVar) Add(duration time.Duration) {
	u.value = u.value.Add(duration)
}

func (u *upTimeVar) String() string {
	return u.value.Format(time.UnixDate)
}
