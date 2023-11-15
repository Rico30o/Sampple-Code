package webtool

type (
	SettlementFields struct {
		AccountNumber string `json:"accountNumber"`
		Event         string `json:"event"`
		IsEnabled     bool   `json:"isEnabled"`
		Description   string `json:"description"`
	}

	SettlementAccount struct {
		AccountNumber string `json:"accountNumber"`
	}
)
