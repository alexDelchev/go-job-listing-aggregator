package listing

// Listing represents a job listing
type Listing struct {
	ID           uint64   `json:"id"`
	ExternalID   string   `json:"externalid"`
	Link         string   `json:"Link"`
	Name         string   `json:"name"`
	WorkSchedule string   `json:"workSchedule"`
	Company      string   `json:"company"`
	Location     string   `json:"location"`
	PostingDate  string   `json:"postingDate"`
	Description  string   `json:"description"`
	Keywords     []string `json:"keywords"`
	QueryID      uint64   `json:"queryId"`
	SourceName   string   `json:"sourceName"`
}
