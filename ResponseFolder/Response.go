package ResponseFolder

type RequestData struct {
	Date         string       `json:"date,omitempty"`
	Version      int          `json:"type,omitempty"`
	Currency     int          `json:"currency,omitempty"`
	Lang         string       `json:"lang,omitempty"`
	SenderCity   SenderCity   `json:"from_location"`
	ReceiverCity ReceiverCity `json:"to_location"`
	Packages     []Size       `json:"packages"`
}
