package model

// Post is the model for the Post resource in the RestApiExample provider.
type Post struct {
	ID     int64  `json:"id,omitempty"`
	Title  string `json:"tirle"`
	Body   string `json:"body"`
	UserId int64  `json:"userId"`
}
