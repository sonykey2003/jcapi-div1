package div1

import (
	"time"
)

// EventServiceTimeQuery is the base struct for most event queries it includes the minimum-required fields of service and start/end times
type EventServiceTimeQuery struct {
	// optional query end time, UTC in RFC3339 format
	EndTime time.Time `json:"end_time,omitempty"`
	// service name to query.
	Service []string `json:"service"`
	// query start time, UTC in RFC3339 format
	StartTime time.Time `json:"start_time"`
}
