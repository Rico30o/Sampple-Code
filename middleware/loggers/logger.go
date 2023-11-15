package loggers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sample/middleware/envRouting"
	"strconv"
	"strings"
	"time"

	"github.com/JohnRebellion/go-utils/database"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	Separator     *log.Logger
)

func OpenLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, err
}

func SetLogs() {
	logFileName := time.Now().Format("2006-01-02") + ".log"
	// create a new log with file name xxx or more the existing one
	f, err := os.OpenFile("./logs/insta-"+logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// set default log output to the 'new' file
	log.SetOutput(f)
	log.Println("This is a test log entry")
	defer f.Close()
}

func CreateDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

// ------------------------------------------------------------
func HealthCheckLog(path, class, senderRequest, receivingBankIdentifierCode, sendingBankIdentifierCode, messageDefinationID, senderDigestValue, senderSignatureValue, responseDigestValue, responseSignatureValue string, senderHeader, senderBody, responseHeader, responseBody, response interface{}, signedInfo []byte, instructionID string) {
	currentTime := time.Now()
	folderName := "./logs/healthCheck/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/ReceivingHealthCheck-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	strReceivingBankCode := fmt.Sprintf("%v", receivingBankIdentifierCode)
	strSendingBankCode := fmt.Sprintf("%v", sendingBankIdentifierCode)
	strMessageDefinition := fmt.Sprintf("%v", messageDefinationID)
	strSenderHeader := fmt.Sprintf("%v", senderHeader)
	strSenderBody := fmt.Sprintf("%v", senderBody)
	strSenderRequest := fmt.Sprintf("%v", senderRequest)
	strResponseHeader := fmt.Sprintf("%v", responseHeader)
	strResponseBody := fmt.Sprintf("%v", responseBody)
	strResponse := fmt.Sprintf("%v", response)
	strSDigestValue := fmt.Sprintf("%v", senderDigestValue)
	strSSignatureValue := fmt.Sprintf("%v", senderSignatureValue)
	strResponseDigestValue := fmt.Sprintf("%v", responseDigestValue)
	strResponseSignatureValue := fmt.Sprintf("%v", responseSignatureValue)
	strSignedInfo := fmt.Sprintf("%v", string(signedInfo))

	Separator.Println("")
	InfoLogger.Println(class + ": - - - HEALTH CHECKS - - -")
	InfoLogger.Println(class + ": INSTRUCTION ID: " + instructionID)
	InfoLogger.Println(class + ": REQUEST HEADER: " + strSenderHeader)
	InfoLogger.Println(class + ": REQUEST BODY: " + strSenderBody)
	InfoLogger.Println(class + ": REQUEST: " + strSenderRequest)
	InfoLogger.Println(class + ": REQUEST DIGEST VALUE: " + strSDigestValue)
	InfoLogger.Println(class + ": REQUEST SIGNATURE VALUE: " + strSSignatureValue)
	InfoLogger.Println(class + ": RECEIVING BANK CODE: " + strReceivingBankCode)
	InfoLogger.Println(class + ": SENDING BANK CODE: " + strSendingBankCode)
	InfoLogger.Println(class + ": MESSAGE DEFINITION: " + strMessageDefinition)
	InfoLogger.Println(class + ": RESPONSE HEADER: " + strResponseHeader)
	InfoLogger.Println(class + ": REPONSE BODY: " + strResponseBody)
	InfoLogger.Println(class + ": REPONSE DIGEST VALUE: " + strResponseDigestValue)
	InfoLogger.Println(class + ": RESPONSE SIGNATURE VALUE: " + strResponseSignatureValue)
	InfoLogger.Println(class + ": RESPONSE SIGNED INFO: " + strSignedInfo)
	InfoLogger.Println(class + ": RESPONSE: " + strResponse)
	defer file.Close()
}

func IPSSignLog(path, class string, requestAppHeader, requestPayload, requestBody interface{}, foodXml, signedInfo, disgestValue, signatureValue, endpoint, responseResult string) {
	currentTime := time.Now()
	folderName := "./logs/" + path + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/"+path+"-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	strRequestAppHeader := fmt.Sprintf("%v", requestAppHeader)
	strRequestPayload := fmt.Sprintf("%v", requestPayload)
	strRequestBody := fmt.Sprintf("%v", requestBody)
	strRequestFoodXML := fmt.Sprintf("%v", foodXml)
	strSignedInfo := fmt.Sprintf("%v", signedInfo)
	strDigestValue := fmt.Sprintf("%v", disgestValue)
	strSignatureValue := fmt.Sprintf("%v", signatureValue)
	strEndpoint := fmt.Sprintf("%v", endpoint)
	strResponse := fmt.Sprintf("%v", responseResult)

	Separator.Println("")
	InfoLogger.Printf(class+": - - - %s - - -\n", path)
	InfoLogger.Println(class + ": SEND TO: " + strEndpoint)
	InfoLogger.Println(class + ": REQUEST HEADER: " + strRequestAppHeader)
	InfoLogger.Println(class + ": REQUEST PAYLOAD: " + strRequestPayload)
	InfoLogger.Println(class + ": REQUEST: " + string(strRequestBody))
	InfoLogger.Println(class + ": FOOD XML: " + strRequestFoodXML)
	InfoLogger.Println(class + ": SIGNED INFO: " + strSignedInfo)
	InfoLogger.Println(class + ": DIGEST VALUE: " + strDigestValue)
	InfoLogger.Println(class + ": SIGNATURE VALUE: " + strSignatureValue)
	InfoLogger.Println(class + ": RESPONSE: " + strResponse)
	defer file.Close()
}

func ServiceHealthCheckLog(class, requestBody, endpoint, signInfo, foodXML, digestValue, signatureValue string, echoRequestHeader, echoRequestBody interface{}, response *http.Response, responseBody, instructionID string) {
	currentTime := time.Now()
	folderName := "./logs/serviceHealthCheck/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/SendingHealthCheck-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	strEndpoint := fmt.Sprintf("%v", endpoint)
	strRequestHeader := fmt.Sprintf("%v", echoRequestHeader)
	strRequestBody := fmt.Sprintf("%v", echoRequestBody)
	strRequest := fmt.Sprintf("%v", requestBody)
	strFoodXML := fmt.Sprintf("%v", foodXML)
	strSignedInfo := fmt.Sprintf("%v", signInfo)
	strDigestValue := fmt.Sprintf("%v", digestValue)
	strSignatureValue := fmt.Sprintf("%v", signatureValue)
	strResponseStatus := fmt.Sprintf("%v", response.Status)
	strResponse := fmt.Sprintf("%v", responseBody)

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - HEALTH CHECK REQUEST - - - -")
	InfoLogger.Println(class + ": SEND TO: " + strEndpoint)
	InfoLogger.Println(class + ": INSTRUCTION ID: " + instructionID)
	InfoLogger.Println(class + ": REQUEST HEADER: " + strRequestHeader)
	InfoLogger.Println(class + ": REQUEST BODY: " + strRequestBody)
	InfoLogger.Println(class + ": REQUEST: " + strRequest)
	InfoLogger.Println(class + ": FOOD XML: " + strFoodXML)
	InfoLogger.Println(class + ": SIGNED INFO: " + strSignedInfo)
	InfoLogger.Println(class + ": DIGEST VALUE: " + strDigestValue)
	InfoLogger.Println(class + ": SIGNATURE VALUE: " + strSignatureValue)
	InfoLogger.Println(class + ": RESPONSE STATUS: " + strResponseStatus)
	InfoLogger.Println(class + ": RESPONSE: " + strResponse)
	defer file.Close()
}

func ServiceHealthCheckErrorLog(class, step string, errors error, echoRequestHeader, echoRequestBody interface{}, resp *http.Response, requestBody, endpoint, signInfo, foodXML, digestValue, signatureValue, instructionID string) {
	currentTime := time.Now()
	folderName := "./logs/serviceHealthCheck/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/SendingHealthCheck-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	strError := fmt.Sprintf("%v", errors)
	strRequestHeader := fmt.Sprintf("%v", echoRequestHeader)
	strRequestBody := fmt.Sprintf("%v", echoRequestBody)
	strResponse := fmt.Sprintf("%v", resp)
	strRequest := fmt.Sprintf("%v", requestBody)
	strEndpoint := fmt.Sprintf("%v", endpoint)
	strSignedInfo := fmt.Sprintf("%v", signInfo)
	strFoodXml := fmt.Sprintf("%v", foodXML)
	strDigestValue := fmt.Sprintf("%v", digestValue)
	strSignaturetValue := fmt.Sprintf("%v", signatureValue)

	Separator.Println("")
	ErrorLogger.Println(class + "-" + step + ": - - - SERVICE HEALTH CHECK ERROR - - -")
	ErrorLogger.Println(class + "-" + step + ": SEND TO: " + strEndpoint)
	ErrorLogger.Println(class + "-" + step + ": INSTRUCTION ID: " + instructionID)
	ErrorLogger.Println(class + "-" + step + ": REQUEST HEADER: " + strRequestHeader)
	ErrorLogger.Println(class + "-" + step + ": REQUEST BODY: " + strRequestBody)
	ErrorLogger.Println(class + "-" + step + ": REQUEST: " + strRequest)
	ErrorLogger.Println(class + "-" + step + ": FOOD XML: " + strFoodXml)
	ErrorLogger.Println(class + "-" + step + ": SIGNED INFO: " + strSignedInfo)
	ErrorLogger.Println(class + "-" + step + ": DIGEST VALUE: " + strDigestValue)
	ErrorLogger.Println(class + "-" + step + ": SIGNATURE VALUE: " + strSignaturetValue)
	ErrorLogger.Println(class + "-" + step + ": CLIENT RESPONSE: " + strResponse)
	ErrorLogger.Println(class + "-" + step + ": ERROR " + strError)
	defer file.Close()
}

func ServiceHealthCheckErrorMarshalLog(class, step string, errors error, echoRequestHeader, echoRequestBody interface{}, echoRequest, endpoint string) {
	currentTime := time.Now()
	folderName := "./logs/serviceHealthCheck/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/SendingHealthCheck-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	strError := fmt.Sprintf("%v", errors)
	strRequestHeader := fmt.Sprintf("%v", echoRequestHeader)
	strRequestBody := fmt.Sprintf("%v", echoRequestBody)
	strRequest := fmt.Sprintf("%v", echoRequest)
	strEndpoint := fmt.Sprintf("%v", endpoint)

	Separator.Println("")
	ErrorLogger.Println(class + "-" + step + ": - - - SERVICE HEALTH CHECK ERROR - - -")
	ErrorLogger.Println(class + "-" + step + ": SEND TO: " + strEndpoint)
	ErrorLogger.Println(class + "-" + step + ": REQUEST HEADER: " + strRequestHeader)
	ErrorLogger.Println(class + "-" + step + ": REQUEST BODY: " + strRequestBody)
	ErrorLogger.Println(class + "-" + step + ": REQUEST : " + strRequest)
	ErrorLogger.Println(class + "-" + step + ": ERROR: " + strError)
	defer file.Close()
}

func CTLogger(class, step, receivingBIC, receivingAccountNumber, receivingName, currency, request, endpoint string, amount float64, header, payload interface{}, response, instructionID, transactionID, senderName, senderAccountNumber, referenceID, senderBIC, localInstrument, requestTriggerTime string) {
	currentTime := time.Now()
	folderName := "./logs/creditTransfer/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/CreditTransfer-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	currentTimeFormat := currentTime.Format("03:04:05")

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	strAmount := fmt.Sprintf("%.2f", amount)
	strHeader := fmt.Sprintf("%v", header)
	strPayload := fmt.Sprintf("%v", payload)

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - CREDIT TRANSFER - - - -")
	InfoLogger.Println(class + ": PROCESS START: " + requestTriggerTime)
	InfoLogger.Println(class + ": PROCESS END: " + currentTimeFormat)
	InfoLogger.Println(class + ": SEND TO: " + endpoint)
	InfoLogger.Println(class + ": SENDER BIC: " + senderBIC)
	InfoLogger.Println(class + ": SENDER AMOUNT: " + currency + " " + strAmount)
	InfoLogger.Println(class + ": SENDER NAME: " + senderName)
	InfoLogger.Println(class + ": SENDER ACCOUNT: " + senderAccountNumber)
	InfoLogger.Println(class + ": INSTRUCTION ID: " + instructionID)
	InfoLogger.Println(class + ": TRANSACTION ID: " + transactionID)
	InfoLogger.Println(class + ": LOCAL INSTRUMENT: " + localInstrument)
	InfoLogger.Println(class + ": REFERENCE ID: " + referenceID)
	InfoLogger.Println(class + ": RECEIVING BIC: " + receivingBIC)
	InfoLogger.Println(class + ": RECEIVING NAME: " + receivingName)
	InfoLogger.Println(class + ": RECEIVING ACCOUNT NUMBER: " + receivingAccountNumber)
	InfoLogger.Println(class + ": REQUEST HEADER: " + strHeader)
	InfoLogger.Println(class + ": REQUEST BODY: " + strPayload)
	InfoLogger.Println(class + ": REQUEST: " + request)
	InfoLogger.Println(class + ": RESPONSE: " + response)
	defer file.Close()

	database.DBConn.Exec("SELECT rbi_instapay.logger(?,?,?,?,?,?,?)", requestTriggerTime, currentTimeFormat, step, instructionID, localInstrument, request, response)
}

func ServiceResponseLogger(class, step, response, instructionID, transactionID, transactionDateTime, transactionStatus, senderBIC, senderCurrency, senderAmount, receivingBIC, referenceID, reasonCode, requestTriggerTime string, status int) {
	currentTime := time.Now()
	folderName := "./logs/serviceResponse/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/ServiceResponse-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	currentTimeFormat := currentTime.Format("03:04:05")
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - SERVICE RESPONSE - - - -")
	InfoLogger.Println(class + ": PROCESS START: " + requestTriggerTime)
	InfoLogger.Println(class + ": PROCESS END: " + currentTimeFormat)
	InfoLogger.Println(class + ": SENDER BIC: " + senderBIC)
	InfoLogger.Println(class + ": SENDER CURRENCY: " + senderCurrency)
	InfoLogger.Println(class + ": SENDER AMOUNT: " + senderAmount)
	InfoLogger.Println(class + ": RECEIVING BIC: " + receivingBIC)
	InfoLogger.Println(class + ": INSTRUCTION ID: " + instructionID)
	InfoLogger.Println(class + ": TRANSACTION ID: " + transactionID)
	InfoLogger.Println(class + ": TRANSACTION DATE/TIME: " + transactionDateTime)
	InfoLogger.Println(class + ": TRANSACTION STATUS: " + transactionStatus)
	InfoLogger.Println(class + ": STATUS: " + strconv.Itoa(status))
	InfoLogger.Println(class + ": REASON CODE: " + reasonCode)
	InfoLogger.Println(class + ": REFERENCE ID: " + referenceID)
	InfoLogger.Println(class + ": RESPONSE: " + response)
	defer file.Close()
	database.DBConn.Exec("SELECT rbi_instapay.logger(?,?,?,?,?,?,?)", requestTriggerTime, currentTimeFormat, step, instructionID, "CHECK LEG-1", "CHECK LEG-1", response)
}

func ServiceRequestLogger(class, step, request, response, instructionID, transactionID, debtorName, debtorAccount, creditorName, creditorAccount, messageDefinitionNameID, instructingAgent, instructedAgent, currency, amount, referenceID, localInstrument, acctValidation, requestTriggerTime, creationDateTime string) {
	currentTime := time.Now()
	folderName := "./logs/serviceRequest/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/ServiceRequest-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	currentTimeFormat := currentTime.Format("03:04:05")
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - SERVICE REQUEST - - - -")
	InfoLogger.Println(class + ": PROCESS START: " + requestTriggerTime)
	InfoLogger.Println(class + ": PROCESS END: " + currentTimeFormat)
	InfoLogger.Println(class + ": HEADER CREATION DATE TIME: " + creationDateTime)
	InfoLogger.Println(class + ": MESSAGE DEFINITION: " + messageDefinitionNameID)
	InfoLogger.Println(class + ": INSTRUCTION ID: " + instructionID)
	InfoLogger.Println(class + ": TRANSACTION ID: " + transactionID)
	InfoLogger.Println(class + ": DEBTOR BIC: " + instructingAgent)
	InfoLogger.Println(class + ": DEBTOR NAME: " + debtorName)
	InfoLogger.Println(class + ": DEBTOR ACCOUNT: " + debtorAccount)
	InfoLogger.Println(class + ": AMOUNT: " + currency + " " + amount)
	InfoLogger.Println(class + ": CREDITOR BIC: " + instructedAgent)
	InfoLogger.Println(class + ": CREDITOR NAME: " + creditorName)
	InfoLogger.Println(class + ": CREDITOR ACCOUNT: " + creditorAccount)
	InfoLogger.Println(class + ": LOCAL INSTRUMENT: " + localInstrument)
	InfoLogger.Println(class + ": ACCOUNT STATUS: " + acctValidation)
	InfoLogger.Println(class + ": REFERENCE ID: " + referenceID)
	InfoLogger.Println(class + ": REQUEST FROM BN: " + request)
	InfoLogger.Println(class + ": RESPONSE TO BN: " + response)

	defer file.Close()

	database.DBConn.Exec("SELECT rbi_instapay.logger(?,?,?,?,?,?,?)", requestTriggerTime, currentTimeFormat, step, instructionID, localInstrument, request, response)
}

func PaymentInstructionsLogger(class, step, loggerType, request, instructionStatus, response, instructionID, currency, amount, rsnCode, requestTriggerTime string) {
	currentTime := time.Now()
	folderName := "./logs/paymentInstructions/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/PaymentInstructions-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	currentTimeFormat := currentTime.Format("03:04:05")

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	if rsnCode == "" {
		rsnCode = "None"
	}

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - PAYMENT INSTRUCTIONS - " + loggerType + " - - - -")
	InfoLogger.Println(class + ": PROCESS START: " + requestTriggerTime)
	InfoLogger.Println(class + ": PROCESS END: " + currentTimeFormat)
	InfoLogger.Println(class + ": INSTRUCTION ID: " + instructionID)
	InfoLogger.Println(class + ": INSTRUCTION STATUS: " + instructionStatus)
	InfoLogger.Println(class + ": AMOUNT: " + currency + " " + amount)
	InfoLogger.Println(class + ": REASON CODE: " + rsnCode)
	InfoLogger.Println(class + ": REQUEST FROM BN: " + request)
	InfoLogger.Println(class + ": RESPONSE TO BN: " + response)
	defer file.Close()
	database.DBConn.Exec("SELECT rbi_instapay.logger(?,?,?,?,?,?,?)", requestTriggerTime, currentTimeFormat, step, instructionID, "CHECK LEG-3", request, response)
}

func SystemNotificationLogger(class, eventCode, eventDateTime, description, parameter, notification string) {
	currentTime := time.Now()
	folderName := "./logs/notification/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/SystemNotification-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	removeString := strings.Replace(parameter, "<string>", " ", -1)
	parameter = strings.Replace(removeString, "</string>", ",", -1)

	if description == "" {
		description = "NO DESCRIPTION"
	} else {
		removeDescriptionString := strings.Replace(description, "<string>", "", -1)
		description = strings.Replace(removeDescriptionString, "</string>", "", -1)
	}

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - NOTIFICATION - - - -")
	InfoLogger.Println(class + ": EVENT CODE: " + eventCode)
	InfoLogger.Println(class + ": DATE/TIME: " + eventDateTime)
	InfoLogger.Println(class + ": DESCRIPTION: " + description)
	InfoLogger.Println(class + ": PARAMETERS: " + parameter)
	InfoLogger.Println(class + ": NOTIFICATION: " + notification)
	defer file.Close()
}

func MessageSignedLogger(class, signedXML, businessMessageID, instructionID, step, result, logName string) {
	currentTime := time.Now()
	folderName := "./logs/signed/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/"+logName+"-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - GENERATING " + step + " - - - -")
	InfoLogger.Println(class + ": PROCESS: " + logName)
	InfoLogger.Println(class + ": INSTRUCTION ID: " + instructionID)
	InfoLogger.Println(class + ": BUSINESS MESSAGE ID: " + businessMessageID)
	InfoLogger.Println(class + ": XML: " + signedXML)
	InfoLogger.Println(class + ": GENERATED " + step + " VALUE: " + result)

	defer file.Close()
}

// Geromme
func CallbackLogs(class, referenceId, instructionId string, request, reponse interface{}) {
	currentTime := time.Now()
	folderName := "./logs/callback/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/CallBack-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	strRequest := fmt.Sprintf("%v", request)
	strResponse := fmt.Sprintf("%v", reponse)

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - CALLBACK - - - -")
	InfoLogger.Println(class + ": INSTRUCTION ID: " + instructionId)
	InfoLogger.Println(class + ": REFERENCE ID: " + referenceId)
	InfoLogger.Println(class + ": REQUEST: " + strRequest)
	InfoLogger.Println(class + ": RESPONSE: " + strResponse)

	defer file.Close()
}

func BalanceInquiry(class, folder, logName, instructionID, responseCode, description, accountNumber, accountName, productCode, productName, referenceID string, availableBalance, currentBalance float64) {
	loc, err := time.LoadLocation(envRouting.PostgresTimeZone)
	if err != nil {
		panic(err)
	}

	currentTime := time.Now().In(loc)
	folderName := "./logs/" + folder + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/"+logName+"-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - " + strings.ToUpper(folder) + " - BALANCE INQUIRY - - - -")
	InfoLogger.Println(class + ": REFERENCE ID: " + referenceID)
	InfoLogger.Println(class + ": INSTRUCTION ID: " + instructionID)
	InfoLogger.Println(class + ": RESPONSE CODE: " + responseCode)
	InfoLogger.Println(class + ": DESCRIPTION: " + description)
	InfoLogger.Println(class + ": ACCOUNT NUMBER: " + accountNumber)
	InfoLogger.Println(class + ": ACCOUNT NAME: " + accountName)
	InfoLogger.Println(class + ": PRODUCT CODE: " + productCode)
	InfoLogger.Println(class + ": PRODUCT NAME: " + productName)
	InfoLogger.Println(class + ": AVAILABLE BALANCE: " + fmt.Sprintf("%.2f", availableBalance))
	InfoLogger.Println(class + ": CURRENT BALANCE: " + fmt.Sprintf("%.2f", currentBalance))

	defer file.Close()
}

// Transaction Credit Logs
func TransactCredit(class, folder, logName, serviceEP, instructionID, referenceId, coreReferenceId string, request, response interface{}) {
	loc, err := time.LoadLocation(envRouting.PostgresTimeZone)
	if err != nil {
		panic(err)
	}

	currentTime := time.Now().In(loc)
	folderName := "./logs/" + folder + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/"+logName+"-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	strRequest := fmt.Sprintf("%s", request)
	strResponse := fmt.Sprintf("%s", response)

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - " + strings.ToUpper(folder) + " - TRANSFER CREDIT RECEIVER- - - -")
	InfoLogger.Println(class + ": SERVICE EP: " + serviceEP)
	InfoLogger.Println(class + ": CORE REFERENCE ID: " + coreReferenceId)
	InfoLogger.Println(class + ": IPS REFERENCE ID: " + referenceId)
	InfoLogger.Println(class + ": INSTRUCTION ID: " + instructionID)
	InfoLogger.Println(class + ": REQUEST: " + strRequest)
	InfoLogger.Println(class + ": RESPONSE: " + strResponse)

	defer file.Close()
}

// TRANSACTION CT RECEIVER LOGS
func TransactCreditReceiver(class, folder, logName, instructionID string, request, response interface{}) {
	currentTime := time.Now()
	folderName := "./logs/" + folder + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/"+logName+"-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	strResponse := fmt.Sprintf("%s", response)
	strRequest := fmt.Sprintf("%s", request)

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - " + strings.ToUpper(folder) + " - TRANSFER CREDIT RECEIVER- - - -")
	InfoLogger.Println(class + ": INSTRUCTION ID: " + instructionID)
	InfoLogger.Println(class + ": REQUEST: " + strRequest)
	InfoLogger.Println(class + ": RESPONSE: " + strResponse)

	defer file.Close()
}

func MessageReject(class, process, instructionID, reason, additionalData string) {
	currentTime := time.Now()
	folderName := "./logs/messageReject/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/MessageReject-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	Separator.Println("")
	InfoLogger.Println(class + ": - - - - MESSAGE REJECT - - - -")
	InfoLogger.Println(class + ": PROCESS: " + process)
	InfoLogger.Println(class + ": BUSINESS MESSAGE ID: " + instructionID)
	InfoLogger.Println(class + ": REJECTING REASON: " + reason)
	InfoLogger.Println(class + ": ADDITIONAL DATA: " + additionalData)

	defer file.Close()
}

func SystemLogger(class, folder, filename, process string, logContent map[string]interface{}) {
	currentTime := time.Now()
	folderName := "./logs/" + folder + "/" + currentTime.Format("01-January")
	CreateDirectory(folderName)
	file, err := os.OpenFile(folderName+"/"+filename+"-"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	Separator = log.New(file, "", log.Ldate|log.Ltime)

	Separator.Println("")

	InfoLogger.Println(class + ": - - - -  " + process + " - - - -")
	for k, v := range logContent {
		InfoLogger.Printf("%v: %v: %v \n", class, k, v)
	}

	defer file.Close()
}
