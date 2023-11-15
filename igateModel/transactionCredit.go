package igateModel

type (
	RequestTransferCredit struct {
		InstructionID    string  `json:"instructionID"`
		ReferenceNumber  string  `json:"referenceNumber"`
		CreditAccount    string  `json:"creditAccount"`
		DebitAccount     string  `json:"debitAccount"`
		TransactionFee   float64 `json:"transactionFee"`
		SourceBranchCode string  `json:"sourceBranchCode"`
		Amount           float64 `json:"amount"`
		AdminFee         float64 `json:"adminFee"`
		Description      string  `json:"description"`
	}

	TransferCredit struct {
		ReferenceNumber  string  `json:"referenceNumber"`
		CreditAccount    string  `json:"creditAccount"`
		DebitAccount     string  `json:"debitAccount"`
		TransactionFee   float64 `json:"transactionFee"`
		SourceBranchCode string  `json:"sourceBranchCode"`
		Amount           float64 `json:"amount"`
		AdminFee         float64 `json:"adminFee"`
		Description      string  `json:"description"`
	}

	TransferCreditResponse struct {
		ResponseCode           string  `json:"responseCode"`
		Description            string  `json:"description"`
		CreditAccount          string  `json:"creditAccount"`
		DebitAccount           string  `json:"debitAccount"`
		CustomerName           string  `json:"customerName"`
		AccountName            string  `json:"accountName"`
		ReferenceNumber        string  `json:"referenceNumber"`
		Amount                 float64 `json:"amount"`
		AdminFee               string  `json:"adminFee"`
		Reff                   string  `json:"reff"`
		CoreReference          string  `json:"coreReference"`
		SourceBranchCode       string  `json:"sourceBranchCode"`
		DestinationBranchCode  string  `json:"destinationBranchCode"`
		SourceProductCode      string  `json:"sourceProductCode"`
		DestinationProductCode string  `json:"destinationProductCode"`
		DebitCurrency          string  `json:"debitCurrency"`
		CreditCurrency         string  `json:"creditCurrency"`
		AvalableBalance        string  `json:"availableBalance"`
		ArNumber               string  `json:"arNumber"`
	}
	TransferCredits struct {
		ResponseCode           string  `json:"responseCode"`
		Description            string  `json:"description"`
		DebitAccount           string  `json:"debitAccount"`
		CreditAccount          string  `json:"creditAccount"`
		CustomerName           string  `json:"customerName"`
		AccountName            string  `json:"accountName"`
		ReferenceNumber        string  `json:"referenceNumber"`
		Amount                 float64 `json:"amount"`
		AdminFee               string  `json:"adminFee"`
		Reff                   string  `json:"reff"`
		CoreReference          string  `json:"coreReference"`
		SourceBranchCode       string  `json:"sourceBranchCode"`
		DestinationBranchCode  string  `json:"destinationBranchCode"`
		SourceProductCode      string  `json:"sourceProductCode"`
		DestinationProductCode string  `json:"destinationProductCode"`
		DebitCurrency          string  `json:"debitCurrency"`
		CreditCurrency         string  `json:"creditCurrency"`
		AvalableBalance        string  `json:"availableBalance"`
		ArNumber               string  `json:"arNumber"`
	}
)
