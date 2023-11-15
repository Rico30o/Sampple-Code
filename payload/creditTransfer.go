package payload

type (
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
)
