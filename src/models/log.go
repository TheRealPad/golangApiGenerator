package models

import "time"

type Log struct {
	Method  string
	Url     string
	Address string
	Time    time.Time
}
