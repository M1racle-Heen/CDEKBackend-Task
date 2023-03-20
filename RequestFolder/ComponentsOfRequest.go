package RequestFolder

type tariffErrors []struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type tariffCodes []struct {
	TariffCode        int     `json:"tariff_code"`
	TariffName        string  `json:"tariff_name"`
	TariffDescription string  `json:"tariff_description"`
	DeliveryMode      int     `json:"delivery_mode"`
	DeliverySum       float64 `json:"delivery_sum"`
	PeriodMin         int     `json:"period_min"`
	PeriodMax         int     `json:"period_max"`
	CalendarMin       int     `json:"calendar_min,omitempty"`
	CalendarMax       int     `json:"calendar_max,omitempty"`
}
