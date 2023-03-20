package ResponseFolder

type RequestData struct {
	Date     string     `json:"date,omitempty"`
	Type     int        `json:"type,omitempty"`
	Currency int        `json:"currency,omitempty"`
	Lang     string     `json:"lang,omitempty"`
	FromCity fromCity   `json:"from_location"`
	ToCity   toCity     `json:"to_location"`
	Packages []packages `json:"packages"`
}
