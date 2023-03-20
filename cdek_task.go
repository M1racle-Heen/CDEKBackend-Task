package cdekCostCalculator

type Client struct {
	Account string `json:"client_id"`
	Secret  string `json:"client_secret"`
	Token   string `json:"access_token"`
	URL     string `json:"api_url"`
}

func NewClient(account string, securePassword string, apiURL string) *Client {
	c := &Client{
		Account: account,
		Secret:  securePassword,
		URL:     apiURL,
	}
	return c
}
