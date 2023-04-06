package div1

import (
	"time"
)

// EventQuery is the users' command to search our auth logs
type EventQuery struct {
	// optional query end time, UTC in RFC3339 format
	EndTime time.Time `json:"end_time,omitempty"`
	// service name to query.
	Service []string `json:"service"`
	// query start time, UTC in RFC3339 format
	StartTime time.Time `json:"start_time"`
	// optional list of fields to return from query
	Fields []string `json:"fields,omitempty"`
	// Max number of rows to return
	Limit int64 `json:"limit,omitempty"`
	// optional string for specifying a full text query
	Q string `json:"q,omitempty"`
	// Specific query to search after, see x-* response headers for next values
	SearchAfter []string    `json:"search_after,omitempty"`
	SearchTerm  *SearchTerm `json:"search_term,omitempty"`
	// ASC or DESC order for timestamp
	Sort string `json:"sort,omitempty"`
}
