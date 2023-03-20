package RequestFolder

type TariffRequest struct {
	TariffCodes  []TariffCodes  `json:"tariff_codes,omitempty"`
	TariffErrors []TariffErrors `json:"errors,omitempty"`
}
