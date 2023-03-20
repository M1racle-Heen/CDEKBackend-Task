package RequestFolder

type TariffRequest struct {
	tariffCodes  `json:"tariff_codes,omitempty"`
	tariffErrors `json:"errors,omitempty"`
}
