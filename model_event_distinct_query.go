package div1

import (
	"time"
)

// EventDistinctQuery is the users' command to search our auth logs for distinct values of the specified field
type EventDistinctQuery struct {
	// optional query end time, UTC in RFC3339 format
	EndTime time.Time `json:"end_time,omitempty"`
	// service name to query.
	Service []string `json:"service"`
	// query start time, UTC in RFC3339 format
	StartTime time.Time `json:"start_time"`
	// field is what they wish to query on
	Field      string      `json:"field"`
	SearchTerm *SearchTerm `json:"search_term,omitempty"`
}
