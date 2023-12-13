package igateModel

import "net/http"

type (
	RequestAccountNumber struct {
		AccountNumber string `json:"accountNumber"`
		InstructionID string `json:"instructionID"`
	}

	RequestDeduct struct {
		AccountNumber string  `json:"accountNumber"`
		Amount        float64 `json:"amount"`
	}

	AccountValidationResponse struct {
		RetCode          int     `json:"retCode"`
		ResponseCode     string  `json:"responseCode"`
		Description      string  `json:"description"`
		ReferenceID      string  `json:"referenceID"`
		AccountNumber    string  `json:"accountNumber"`
		AccountName      string  `json:"accountName"`
		ProductCode      string  `json:"productCode"`
		ProductName      string  `json:"productName"`
		Amount           float64 `json:"amount"`
		CurrencyCode     string  `json:"currencyCode"`
		ReferenceName    string  `json:"referenceName"`
		AvailableBalance float64 `json:"availableBalance"`
		CurrentBalance   float64 `json:"currentBalance"`
	}
	ResponseFeeDetails struct {
		FeeID             int     `json:"fee_id"`
		StartRange        int     `json:"start_range"`
		EndRange          int     `json:"end_range"`
		TotalCharge       float64 `json:"total_charge"`
		AgentIncome       float64 `json:"agent_income"`
		FdsFee            float64 `json:"fds_fee"`
		CmitFee           float64 `json:"cmit_fee"`
		BankIncome        float64 `json:"bank_income"`
		CreatedBy         int     `json:"created_by"`
		CreatedDate       string  `json:"created_date"`
		LastUpdatedBy     int     `json:"last_updated_by"`
		LastUpdatedDate   string  `json:"last_updated_date"`
		TransType         string  `json:"trans_type"`
		TelcoFee          float64 `json:"telco_fee"`
		BankIncomeFlag    bool    `json:"bank_income_flag"`
		ClientType        string  `json:"client_type"`
		AgentTargetIncome float64 `json:"agent_target_income"`
		BancnetIncome     float64 `json:"bancnet_income"`
	}

	// inqbalance

	InqBalanceRequest struct {
		ReferenceNumber string `json:"referenceNumber"`
		AccountNumber   string `json:"accountNumber"`
		Response_Data   Response_Data
	}

	InqBalanceResponse struct {
		Message string      `json:"message"`
		Header  http.Header `json:"header"`
		Data    string      `json:"data"`
	}

	Response_Data struct {
		ResponseCode     string  `json:"responseCode"`
		Description      string  `json:"description"`
		AccountNumber    string  `json:"accountNumber"`
		AccountName      string  `json:"accountName"`
		ProductCode      string  `json:"productCode"`
		ProductName      string  `json:"productName"`
		CurrencyCode     string  `json:"currencyCode"`
		ReferenceName    string  `json:"referenceName"`
		AvailableBalance float64 `json:"availableBalance"`
		CurrentBalance   float64 `json:"currentBalance"`
		JointHolder      string  `json:"jointHolder"`
	}

	// JSON

	JSONRequestCreditTransfer struct {
		ReceivingBIC           string  `json:"receivingBIC"`
		ReceivingAccountNumber string  `json:"receivingAccountNumber"`
		ReceivingName          string  `json:"receivingName"`
		SenderBIC              string  `json:"senderBIC"`
		SenderName             string  `json:"senderName"`
		SenderAccountNumber    string  `json:"senderAccountNumber"`
		Amount                 float64 `json:"amount"`
		Currency               string  `json:"currency"`
		LocalInstrument        string  `json:"localInstrument"`
		ReferenceID            string  `json:"referenceId"`
		AppId                  string  `json:"appId"`
	}

	RequestCreditTransfer struct {
		SourceAccountDetails SourceAccountDetails
		TargetAccountDetails TargetAccountDetails
	}

	SourceAccountDetails struct {
		AccountType   string `json:"accountType"`
		AccountNumber string `json:"accountNumber"`
	}

	TargetAccountDetails struct {
		BankName          string  `json:"bankName"`
		AccountNumber     string  `json:"accountNumber"`
		AccountHolderName string  `json:"accountHolderName"`
		TransactionAmount float64 `json:"transactionAmount"`
		TransactionCharge float64 `json:"transactionCharge"`
	}

	CreditTransferJSON struct {
		TransactionType     string  `json:"transactionType"`
		Status              string  `json:"status"`
		ReasonCode          string  `json:"reasonCode"`
		LocalInstrument     string  `json:"localInstrument"`
		InstructionID       string  `json:"instructionId"`
		TransactionID       string  `json:"transactionId"`
		ReferenceID         string  `json:"referenceId"`
		SenderBIC           string  `json:"senderBIC"`
		SenderName          string  `json:"senderName"`
		SenderAccount       string  `json:"senderAccount"`
		AmountCurrency      string  `json:"amountCurrency"`
		SenderAmount        float64 `json:"senderAmount"`
		ReceivingBIC        string  `json:"receivingBIC"`
		ReceivingName       string  `json:"receivingName"`
		ReceivingAccount    string  `json:"receivingAccount"`
		TransactionDateTime string  `json:"transactionDateTime"`
		BNResponse          string  `json:"bnResponse"`
	}

	SystemParameter struct {
		Parameter string `json:"parameter"`
		Value     string `json:"value"`
	}
)

type (
	RequestTransferCredits struct {
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

	TransferCreditss struct {
		ReferenceNumber  string  `json:"referenceNumber"`
		CreditAccount    string  `json:"creditAccount"`
		DebitAccount     string  `json:"debitAccount"`
		TransactionFee   float64 `json:"transactionFee"`
		SourceBranchCode string  `json:"sourceBranchCode"`
		Amount           float64 `json:"amount"`
		AdminFee         float64 `json:"adminFee"`
		Description      string  `json:"description"`
	}

	TransferCreditResponses struct {
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
)
