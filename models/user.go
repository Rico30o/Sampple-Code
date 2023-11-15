package models

import (
	"encoding/xml"
	"net/http"
	"sample/bah"
	"time"

	"gorm.io/gorm"
)

type (
	User struct {
		ID       int
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json :"password"`
	}

	Login struct {
		Username string `xml:"username"`
		Email    string `xml:"email"`
		Password string `xml:"password"`
	}

	LoginResponse struct {
		XMlName xml.Name `xml:"response"`
		Message string   `xml:"message"`
	}

	// ////////////////////
	// /////////try///////
	Wew struct {
		XMLName  xml.Name `xml:"WewHead"`
		Name     string   `xml:"name"`
		Lastname string   `xml:"lastname"`
		Friend   Friend
	}

	Friend struct {
		XMLName  xml.Name `xml:"FriendHead"`
		Name     string   `xml:"name"`
		Lastname string   `xml:"lastname"`
	}

	// /////////////////////////
	// /Another try again /////
	Xample struct {
		XMLName  xml.Name `xml:"xamplehead"`
		Name     string   `xml:"name"`
		Address  string   `xml:"address"`
		Email    string   `xml:"email"`
		Employee Employee
	}

	Employee struct {
		XMLName      xml.Name `xml:"WewHead"`
		NameEmployee string   `xml:"nameemplyee"`
	}

	////////Boring roots muna kasi nakakalito////

	InstaPay struct {
		MLName          xml.Name `xml:"instapayHead"`
		Name3           string   `xml:"name3"`
		ISO_Description string   `xml:"iso_description"`
		Product_Usage   string   `xml:"product_usage"`
		Type            string   `xml:"type"`
		Occure          string   `xml:"occure"`
	}

	// /// HAHAHAHHAHA /////
	SignofInstap struct {
		MLName           xml.Name `xml:"signofinstapHead"`
		Name4            string   `xml:"Name4"`
		ISO_Description1 string   `xml:"iso_description1"`
		Product_Usage1   string   `xml:"product_usage1"`
		Type1            string   `xml:"type1"`
		Occurence1       string   `xml:"occurence1"`
		Group_Header     Group_Header
	}

	Group_Header struct {
		Name5            string `xml:"name5"`
		ISO_Description2 string `xml:"iso-description"`
		Type2            string `xml:"type2"`
		Occurence2       string `xml:"occurence2"`
	}

	Notif_status struct {
		gorm.Model
		Eventcode     string `json:"eventCode"`
		DateTime      string `json:"date_time"`
		Description   string `json:"description"`
		Parameters    string `json:"parameters"`
		Notifications string `json:"notificationsData"`
	}

	Notifications_Data struct {
		Is_signed_on bool      `json:"is_signed_on"`
		Remark       string    `json:"remark"`
		Authority    string    `json:"authority"`
		Signoff_date time.Time `json:"signoff_date"`
		Signoff_time time.Time `json:"signoff_time"`
		Signon_date  time.Time `json:"signon_date"`
		Signon_time  time.Time `json:"signon_time"`
		Create_at    time.Time `json:"create_at"`
	}
	Logs_Notification struct {
		Signed_on    bool      `json:"signed_on"`
		Remarks      string    `json:"remarks"`
		Signed_by    string    `json:"signed_by"`
		Signoff_date time.Time `json:"signoff_date"`
		Signoff_time time.Time `json:"signoff_time"`
		Signon_date  time.Time `json:"signon_date"`
		Signon_time  time.Time `json:"signon_time"`
		Create_at    time.Time `json:"create_at"`
	}

	AnotherTry struct {
		Signed_on      bool      `json:"signed_on"`
		Remarks        string    `json:"remarks"`
		Signed_by      string    `json:"signed_by"`
		Signoff_date   time.Time `json:"signoff_date"`
		Signoff_time   time.Time `json:"signoff_time"`
		Signon_date    time.Time `json:"signon_date"`
		Signon_time    time.Time `json:"signon_time"`
		Create_at      time.Time `json:"create_at"`
		CustomSignedBy string    `json:"custom_signed_by"`
	}
	// paysRequest represents the request body for the Pays endpoint.
	paysRequest struct {
		CustomSignedBy string `json:"custom_signed_by" example:"John Doe"`
		ExactDate      bool   `json:"Exactdate" example:"true"`
		SignonDate     string `json:"signon_date" example:"2023-10-11"`
	}

	// paysResponse represents the response body for the Pays endpoint.
	paysResponse struct {
		Message string `json:"message" example:"Data inserted successfully"`
	}

	// ErrorResponse represents an error response for Swagger documentation.
	ErrorResponse struct {
		Error string `json:"error" example:"Invalid request body"`
	}
	InsertedRecord struct {
		Signed_on bool      `json:"signed_on"`
		Signed_by string    `json:"signed_by"`
		Create_at time.Time `json:"create_at"`
	}
	TransferRequest struct {
		ReferenceNumber string  `json:"referenceNumber"`
		CreditAccount   string  `json:"creditAccount"`
		DebitAccount    string  `json:"debitAccount"`
		Description     string  `json:"description"`
		Amount          float64 `json:"amount"`
		AdminFee        float64 `json:"adminFee"`
	}
	TransferData struct {
		ReferenceNumber string  `json:"referenceNumber"`
		CreditAccount   string  `json:"creditAccount"`
		DebitAccount    string  `json:"debitAccount"`
		Description     string  `json:"description"`
		Amount          float64 `json:"amount"`
		AdminFee        float64 `json:"adminFee"`
	}
	InquiryTransferCredit struct {
		ResponseCode    string  `json:"responsecode"`
		Description     string  `json:"description"`
		CreditAccoun    string  `json:"creditAccount"`
		DebitAccount    string  `json:"debitAccount"`
		ReferenceNumber string  `json:"referenceNumber"`
		CreditName      string  `json:"creditname"`
		Amount          float64 `json:"amount"`
	}
)

type (
	TransferCredit struct {
		ReferenceNumber  string  `json:"referenceNumber"`
		CreditAccount    string  `json:"creditAccount"`
		DebitAccount     string  `json:"debitAccount"`
		TransactionFee   float32 `json:"transactionFee"`
		SourceBranchCode string  `json:"sourceBranchCode"`
		Amount           float32 `json:"amount"`
		AdminFee         float32 `json:"adminFee"`
		Description      string  `json:"description"`
	}

	TransferCreditResponse struct {
		ResponseCode    string  `json:"responseCode"`
		Description     string  `json:"description"`
		CreditAccount   string  `json:"creditAccount"`
		DebitAccount    string  `json:"debitAccount"`
		CustomerName    string  `json:"customerName"`
		AccountName     string  `json:"accountName"`
		ReferenceNumber string  `json:"referenceNumber"`
		Amount          float32 `json:"amount"`
		AdminFee        string  `json:"adminFee"`
		Reff            string  `json:"reff"`

		// "reff": "0718101137191816OPENAPITFCR-544303",
		// "coreReference": "FT20308RX9FV",
		// "sourceBranchCode": "PH1030032",
		// "destinationBranchCode": "PH1030001",
		// "sourceProductCode": "6003",
		// "destinationProductCode": "6007",
		// "debitCurrency": "PHP",
		// "creditCurrency": "PHP",
		// "availableBalance": "7141211",
		// "arNumber": "AR-AAA-000000000234346"
	}
)

type (
	Transfer_Request struct {
		ReferenceNumber    string  `json:"referenceNumber"`
		CreditAccount      string  `json:"creditAccount"`
		DebitAccount       string  `json:"debitAccount"`
		TransactionFee     float64 `json:"transactionFee"`
		SourceBranchCode   string  `json:"sourceBranchCode"`
		Amount             float64 `json:"amount"`
		Description        string  `json:"description"`
		AdminFee           float64 `json:"adminFee"`
		Transfer_Responces Transfer_Responces
	}

	Trans_Response struct {
		Message string      `json:"message"`
		Header  http.Header `json:"header"`
		Data    string      `json:"data"`
	}

	Transfer_Responces struct {
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
		AvailableBalance       string  `json:"availableBalance"`
		ARNumber               string  `json:"arNumber"`
	}
	// models/another_try.go

	// AnotherTry represents the structure of the request body
	AnotherTrys struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
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
)

// SYSTEM STATUS BROADCAST MESSAGE
type (
	SystemNotificationISO20022 struct {
		XMLName   xml.Name                       `xml:"Message"`
		XMLNS     string                         `xml:"xmlns,attr"`
		XMLNsHead string                         `xml:"head,attr"`
		XMLNsNe   string                         `xml:"ne,attr"`
		Header    bah.HCRequestApplicationHeader `xml:"AppHdr"`
		Body      SystemNotificationWrapper
	}

	SystemNotificationWrapper struct {
		XMLName xml.Name               `xml:"SystemNotificationEvent"`
		Body    SystemNotificationBody `xml:"SysEvtNtfctn"`
	}

	SystemNotificationBody struct {
		XMLName            xml.Name           `xml:"SysEvtNtfctn"`
		SystemNotification SystemNotification `xml:"EvtInf"`
	}
	SystemNotification struct {
		EventCode        string   `xml:"EvtCd"`
		EventParams      []string `xml:"EvtParam"`
		EventDescription []string `xml:"EvtDesc"`
		EvenTime         string   `xml:"EvtTm"`
	}
)

type (
	TransactCredit struct {
		ReceivingBIC     string `json:"receivingBIC"`
		ReceivingAccount string `json:"receivingAccountNumber"`
		ReceivingName    string `json:"receivingName"`
		SenderName       string `json:"senderName"`
		SenderBIC        string `json:"senderBIC"`
		SenderAccount    string `json:"senderAccountNumber"`
		Amount           string `json:"amount"`
		Currency         string `json:"currency"`
		ReferenceId      string `json:"referenceId"`
		InstructionId    string `json:"instapayReference"`
	}
	//used for getting the request body
	RequestMessageStatusReportISO20022 struct {
		XMLName   xml.Name                       `xml:"Message"`
		XMLNS     string                         `xml:"xmlns,attr"`
		XMLNSHead string                         `xml:"head,attr"`
		XMLNSPS   string                         `xml:"ps,attr"`
		Header    bah.HCRequestApplicationHeader `xml:"AppHdr"`
	}

	RejectMessageStatusReportISO20022 struct {
		XMLName   xml.Name              `xml:"Message"`
		XMLNS     string                `xml:"xmlns,attr"`
		XMLNSHead string                `xml:"xmlns:head,attr"`
		XMLNSPs   string                `xml:"xmlns:ps,attr"`
		Header    bah.ApplicationHeader `xml:"AppHdr"`
	}
	// used to get the request body
	RequestCreditTransferISO20022 struct {
		XMLName   xml.Name                           `xml:"Message"`
		XMLNS     string                             `xml:"xmlns,attr"`
		XMLNSct   string                             `xml:"ct,attr"`
		XMLNShead string                             `xml:"head,attr"`
		Header    bah.RequestBranchApplicationHeader `xml:"AppHdr"`
		// Body      RequestCreditTransferWrapper       `xml:"CreditTransfer"`
	}
	CreditTransferSending struct {
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
