package div1

type InlineResponse200 struct {
	Buckets                 []InlineResponse200Buckets `json:"buckets,omitempty"`
	DocCountErrorUpperBound int32                      `json:"doc_count_error_upper_bound,omitempty"`
	SumOtherDocCount        int32                      `json:"sum_other_doc_count,omitempty"`
}
