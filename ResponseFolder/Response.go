package ResponseFolder

type RequestData struct {
	Date     string     `json:"date,omitempty"`
	Type     int        `json:"type,omitempty"`
	Currency int        `json:"currency,omitempty"`
	Lang     string     `json:"lang,omitempty"`
	FromCity FromCity   `json:"from_location"`
	ToCity   ToCity     `json:"to_location"`
	Packages []Packages `json:"packages"`
}
