package div1

// SearchTerm is the filter portion of the query it contains only one of 'and', 'or', or 'not' conjunction maps
type SearchTerm struct {
	And *map[string]interface{} `json:"and,omitempty"`
	Not *map[string]interface{} `json:"not,omitempty"`
	Or  *map[string]interface{} `json:"or,omitempty"`
}
