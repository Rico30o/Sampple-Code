package payload

import "time"

// Trace- Netweork Alert ID//

type (
	NetworkAlertID struct {
		Network_alert_id string `json:"network_alert_id"`
		Format           string `json:"format"`
		Width            int64  `json:"wiidth"`
		Height           int64  `json:"height"`
		Legend           bool   `json:"legend"`
		Type             string `json:"type"`
		Colour_Mode      string `json:"Colour_Mode"`
		// SomeData         string `json:"somedata"`
	}
)

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
	Login struct {
		User     string `jason:"User"`
		Password string `jason:"password"`
	}
)

// ////FEEEDBACK/////
// type (
// 	RequestBody struct {
// 		AlertID      string `json:"alertID"`
// 		AlertType    string `json:"alertType"`
// 		EntityID     string `json:"entityID"`
// 		DecisionDate string `json:"decisionDate"`
// 		Feedback     string `json:"feedback"`
// 	}
// 	// ResponseBody represents the structure of the response JSON
// 	ResponseBody struct {
// 		FeedbackID string `json:"feedbackID"`
// 	}
// )

// Trace//
type (
	Transaction struct {
		TxnID string `json:"txn_id"`
		Type  string `json:"type"`
	}
	// TraceResponse struct to represent the response body when transaction is traced
	TraceResponse struct {
		ID                 string             `json:"id"`
		Time               time.Time          `json:"time"`
		NetworkID          string             `json:"networkID"`
		TransactionAlerts  []TransactionAlert `json:"transactionAlerts"`
		AccountAlerts      []AccountAlert     `json:"accountAlerts"`
		VizURL             string             `json:"vizURL"`
		SourceTxnID        string             `json:"sourceTxnID"`
		SourceTxnType      string             `json:"sourceTxnType"`
		Length             int                `json:"length"`
		Generations        int                `json:"generations"`
		TotalValue         int                `json:"totalValue"`
		SourceValue        int                `json:"sourceValue"`
		UniqueAccounts     int                `json:"uniqueAccounts"`
		MeanDwellTime      string             `json:"meanDwellTime"`
		MedianDwellTime    string             `json:"medianDwellTime"`
		MeanMuleScore      float64            `json:"meanMuleScore"`
		ElapsedTime        string             `json:"elapsedTime"`
		NumActionedMules   int                `json:"numActionedMules"`
		NumLegitimate      int                `json:"numLegitimate"`
		NumNotInvestigated int                `json:"numNotInvestigated"`
		ParentAlertID      string             `json:"parentAlertID"`
		DecisionDate       time.Time          `json:"decisionDate"`
		MostRecentFeedback string             `json:"mostRecentFeedback"`
	}
	// TransactionAlert struct to represent the transaction alerts
	TransactionAlert struct {
		ID             string    `json:"id"`
		TxnID          string    `json:"txnID"`
		NetworkAlertID string    `json:"networkAlertID"`
		NetworkID      string    `json:"networkID"`
		Time           time.Time `json:"time"`
		TxnTime        time.Time `json:"txnTime"`
		SourceID       string    `json:"sourceID"`
		DestID         string    `json:"destID"`
		SourceBankID   string    `json:"sourceBankID"`
		SourceBankName string    `json:"sourceBankName"`
		DestBankID     string    `json:"destBankID"`
		DestBankName   string    `json:"destBankName"`
		Value          int       `json:"value"`
	}
	// AccountAlert struct to represent the account alerts
	AccountAlert struct {
		ID             string    `json:"id"`
		NetworkAlertID string    `json:"networkAlertID"`
		AccountID      string    `json:"accountID"`
		NetworkID      string    `json:"networkID"`
		OwningBankID   string    `json:"owningBankID"`
		OwningBankName string    `json:"owningBankName"`
		Time           time.Time `json:"time"`
	}
)

type RequestBody struct {
	AlertID      string `json:"alertID"`
	AlertType    string `json:"alertType"`
	EntityID     string `json:"entityID"`
	DecisionDate string `json:"decisionDate"`
	Feedback     string `json:"feedback"`
	CustomerID   string `json:"customerID"`

	ResponseBody struct {
		FeedbackID string ` json:"feedbackID"`
	}
}
type ErrorDetail struct {
	Source      string      `json:"Source"`
	ReasonCode  string      `json:"ReasonCode"`
	Description string      `json:"Description"`
	Recoverable bool        ` json:"Recoverable"`
	Details     interface{} `json:"Details,omitempty"`
}

type ErrorResponses struct {
	Errors struct {
		Error []ErrorDetail `json:"Error"`
	} `json:"Errors"`
}
type RequestBodyArray struct {
	Items []RequestBody `json:"items"`
}
