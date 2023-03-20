package ResponseFolder

type Cities struct {
	Code        int    `json:"code,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	City        string `json:"city,omitempty"`
	Address     string `json:"address"`
}

type FromCity struct {
	Cities
}

type ToCity struct {
	Cities
}

type Packages struct {
	Weight int `json:"weight"`
	Length int `json:"length,omitempty"`
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}
