package RequestFolder

type TariffRequest struct {
	TariffCodes  `json:"tariff_codes,omitempty"`
	TariffErrors `json:"errors,omitempty"`
}
