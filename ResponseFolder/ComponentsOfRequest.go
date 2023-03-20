package ResponseFolder

type Cities struct {
	Code        int    `json:"code,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	City        string `json:"city,omitempty"`
	Address     string `json:"address"`
}

type SenderCity struct {
	Cities
}

type ReceiverCity struct {
	Cities
}

type Size struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
	Length int `json:"length,omitempty"`
	Weight int `json:"weight"`
}
