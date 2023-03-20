package ResponseFolder

type cities struct {
	Code        int    `json:"code,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	City        string `json:"city,omitempty"`
	Address     string `json:"address"`
}

type fromCity struct {
	cities
}

type toCity struct {
	cities
}

type packages struct {
	Weight int `json:"weight"`
	Length int `json:"length,omitempty"`
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}
