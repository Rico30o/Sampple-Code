package igate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sample/igateModel"
	"sample/middleware/encryptDecrypt"
	"sample/middleware/envRouting"
	"sample/middleware/loggers"
	"sample/models"
	"sample/util"
	webtool "sample/webTool"
	"strconv"
	"strings"

	"github.com/JohnRebellion/go-utils/database"
	"github.com/gofiber/fiber"
)

// Used for fund transfer successful transaction
func CreditTransfer(c *fiber.Ctx, instructionID, transactionType string) (bool, error) {
	fmt.Println("Start Transfer Credit")
	log.Println("Start Transfer Credit")
	transactCredit := &models.TransactCredit{}
	database.DBConn.Raw("SELECT * FROM rbi_instapay.ct_transaction WHERE instruction_id = ?", instructionID).Scan(transactCredit)

	// Fetch settlement account and decrypt the data
	settlementAccount := &webtool.SettlementAccount{}
	database.DBConn.Debug().Raw("SELECT account_number FROM rbi_instapay.settlement WHERE event = 'receiving'").Scan(settlementAccount)
	decryptedAccountNumber, _ := encryptDecrypt.Decrypt(settlementAccount.AccountNumber, envRouting.SecretKey)
	amount, _ := strconv.ParseFloat(transactCredit.Amount, 64)

	transferCredit := &igateModel.TransferCredit{}

	if transactionType == "receiving" {
		transferCredit = &igateModel.TransferCredit{
			ReferenceNumber: transactCredit.ReferenceId,
			CreditAccount:   transactCredit.ReceivingAccount,
			DebitAccount:    decryptedAccountNumber,
			Amount:          amount,
			Description:     fmt.Sprintf("%v %v", transactCredit.ReferenceId, "Instapay Receiving Fund Transfer"),
		}
	} else if transactionType == "sending" {
		transferCredit = &igateModel.TransferCredit{
			ReferenceNumber: transactCredit.ReferenceId,
			CreditAccount:   decryptedAccountNumber,
			DebitAccount:    transactCredit.SenderAccount,
			Amount:          amount,
			Description:     fmt.Sprintf("%v %v", transactCredit.ReferenceId, "Instapay Receiving Fund Transfer"),
		}
	}

	transferCreditRequirements, err := json.Marshal(transferCredit)
	if err != nil {
		fmt.Println("Error in JSON marshal:", err)
		return false, err
	}

	fmt.Println("Transfer Credit:", transferCredit)
	// This will get the endpoint from DB
	ServiceEP := util.GetServiceEP("CreditTransfer_igate", strings.ToLower(envRouting.Environment))

	client := &http.Client{}
	req, reqErr := http.NewRequest(http.MethodPost, ServiceEP, bytes.NewBuffer(transferCreditRequirements))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Merchant-ID", "QVBJMDAwMDU=")

	if reqErr != nil {
		fmt.Println("Error requesting:", err)
		return false, reqErr
	}

	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	resp, respErr := client.Do(req)
	if respErr != nil {
		fmt.Println("Getting error request:", err)
		return false, respErr
	}

	resultTransactCredit := &igateModel.TransferCreditResponse{}
	decodErr := json.NewDecoder(resp.Body).Decode(resultTransactCredit)
	if decodErr != nil {
		return false, decodErr
	}

	defer resp.Body.Close()

	fmt.Println("---------------------------------------")
	fmt.Println("SERVICE EP:", ServiceEP)
	fmt.Println("TRANSFER CREDIT:", transferCredit)
	fmt.Println("CORE REFERENCE ID:", resultTransactCredit.CoreReference)
	fmt.Println("RECEIVING BIC:", transactCredit.ReceivingBIC)
	fmt.Println("RECEIVING NAME:", transactCredit.ReceivingName)
	fmt.Println("RECEIVING ACCOUNT:", transactCredit.ReceivingAccount)
	fmt.Println("SENDER BIC:", transactCredit.SenderBIC)
	fmt.Println("SENDER NAME:", transactCredit.SenderName)
	fmt.Println("SENDER ACCOUNT:", transactCredit.SenderAccount)
	fmt.Println("AMOUNT:", transactCredit.Amount)
	fmt.Println("AVAILABLE BALANCE:", resultTransactCredit.AvalableBalance)
	fmt.Println("INSTRUCTION ID:", transactCredit.InstructionId)
	fmt.Println("REFERENCE ID:", resultTransactCredit.ReferenceNumber)
	fmt.Println("RESPONSE:", resp.Body)
	fmt.Println("---------------------------------------")

	loggers.TransactCredit(c.Path(), "igate", "Transfer_Credit_Receiving", ServiceEP, instructionID, transactCredit.ReferenceId, resultTransactCredit.CoreReference, transferCreditRequirements, resultTransactCredit)
	log.Println("End Transfer Credit")
	fmt.Println("End Transfer Credit")
	return true, nil
}

// POSTMAN
func TransferCreditProcess(c *fiber.Ctx) error {
	transferCreditFields := &igateModel.RequestTransferCredit{}
	if parsErr := c.BodyParser(transferCreditFields); parsErr != nil {
		return c.JSON(fiber.Map{
			"message": "error parsing",
			"data":    parsErr.Error(),
		})
	}

	transactCredit := &models.TransactCredit{}
	database.DBConn.Raw("SELECT * FROM rbi_instapay.ct_transaction WHERE instruction_id = ?", transferCreditFields.InstructionID).Scan(transactCredit)

	settlementAccount := &webtool.SettlementAccount{}
	database.DBConn.Debug().Raw("SELECT account_number FROM rbi_instapay.settlement WHERE event = 'receiving'").Scan(settlementAccount)
	decryptedAccountNumber, _ := encryptDecrypt.Decrypt(settlementAccount.AccountNumber, envRouting.SecretKey)

	transferCredit := &igateModel.TransferCredit{
		ReferenceNumber: transactCredit.ReferenceId,
		CreditAccount:   transferCreditFields.CreditAccount,
		DebitAccount:    decryptedAccountNumber,
		Amount:          transferCreditFields.Amount,
		Description:     transferCreditFields.Description,
	}

	transferCreditRequirements, marshalErr := json.Marshal(transferCredit)
	if marshalErr != nil {
		return c.JSON(fiber.Map{
			"message": "marshal error",
			"error":   marshalErr.Error(),
		})
	}

	// This will get the endpoint from DB
	ServiceEP := util.GetServiceEP("CreditTransfer_igate", strings.ToLower(envRouting.Environment))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, ServiceEP, bytes.NewBuffer(transferCreditRequirements))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Merchant-ID", "QVBJMDAwMDU=")

	fmt.Println("REQUEST:", req)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "http request error",
			"error":   err.Error(),
		})
	}

	res, err := client.Do(req)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "client request error",
			"error":   err.Error(),
		})
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "reading body error",
			"error":   err.Error(),
		})
	}

	// response := &igateModel.TransferCreditResponse{}
	response := &igateModel.TransferCreditResponse{}
	if unmarshalErr := json.Unmarshal(body, response); unmarshalErr != nil {
		return c.JSON(fiber.Map{
			"message": "unmarshal error",
			"error":   err.Error(),
		})
	}

	fmt.Println("RESPONSE:", response)
	return c.JSON(fiber.Map{
		"serviceEP": ServiceEP,
		"response":  response,
	})
}

//	func CreditTransferSending(c *fiber.Ctx, instructionID string) string {
//		return ""
//	}
//
// func CreditTransferSending(c *fiber.Ctx, transferCreditFields igateModel.RequestTransferCredit) error

// func CreditTransferSending(c *fiber.Ctx, igateModel.RequestTransferCredit) string {
// 	transactCredit := &models.TransactCredit{}
// 	database.DBConn.Raw("SELECT * FROM rbi_instapay.ct_transaction WHERE instruction_id = ?", transferCreditFields.InstructionID).Scan(transactCredit)

// 	settlementAccount := &webtool.SettlementAccount{}
// 	database.DBConn.Debug().Raw("SELECT account_number FROM rbi_instapay.settlement WHERE event = 'receiving'").Scan(settlementAccount)
// 	decryptedAccountNumber, _ := encryptDecrypt.Decrypt(settlementAccount.AccountNumber, envRouting.SecretKey)

// 	transferCredit := &igateModel.TransferCredit{
// 		ReferenceNumber: transactCredit.ReferenceId,
// 		CreditAccount:   transferCreditFields.CreditAccount,
// 		DebitAccount:    decryptedAccountNumber,
// 		Amount:          transferCreditFields.Amount,
// 		Description:     transferCreditFields.Description,
// 	}

// 	transferCreditRequirements, marshalErr := json.Marshal(transferCredit)
// 	if marshalErr != nil {
// 		return &fiber.Map{
// 			"message": "marshal error",
// 			"error":   marshalErr.Error(),
// 		}, marshalErr
// 	}

// 	// This will get the endpoint from DB
// 	ServiceEP := util.GetServiceEP("CreditTransfer_igate", strings.ToLower(envRouting.Environment))

// 	client := &http.Client{}
// 	req, err := http.NewRequest(http.MethodPost, ServiceEP, bytes.NewBuffer(transferCreditRequirements))
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Merchant-ID", "QVBJMDAwMDU=")

// 	fmt.Println("REQUEST:", req)
// 	if err != nil {
// 		return &fiber.Map{
// 			"message": "http request error",
// 			"error":   err.Error(),
// 		}, err
// 	}

// 	res, err := client.Do(req)
// 	if err != nil {
// 		return &fiber.Map{
// 			"message": "client request error",
// 			"error":   err.Error(),
// 		}, err
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return &fiber.Map{
// 			"message": "reading body error",
// 			"error":   err.Error(),
// 		}, err
// 	}

// 	response := &igateModel.TransferCreditResponse{}
// 	if unmarshalErr := json.Unmarshal(body, response); unmarshalErr != nil {
// 		return &fiber.Map{
// 			"message": "unmarshal error",
// 			"error":   err.Error(),
// 		}, unmarshalErr
// 	}

// 	fmt.Println("RESPONSE:", response)
// 	return &fiber.Map{
// 		"serviceEP": ServiceEP,
// 		"response":  response,
// 	}, nil
// }

func CreditTransferSending(c *fiber.Ctx, transferCreditFields igateModel.RequestTransferCredit) string {

	transactCredit := &models.TransactCredit{}
	// database.DBConn.Raw("SELECT * FROM rbi_instapay.ct_transaction WHERE instruction_id = ?", transferCreditFields.InstructionID).Scan(transactCredit)

	settlementAccount := &webtool.SettlementAccount{}
	// database.DBConn.Debug().Raw("SELECT account_number FROM rbi_instapay.settlement WHERE event = 'receiving'").Scan(settlementAccount)
	decryptedAccountNumber, _ := encryptDecrypt.Decrypt(settlementAccount.AccountNumber, envRouting.SecretKey)

	transferCredit := &igateModel.TransferCredit{
		ReferenceNumber: transactCredit.ReferenceId,
		CreditAccount:   transferCreditFields.CreditAccount,
		DebitAccount:    decryptedAccountNumber,
		Amount:          transferCreditFields.Amount,
		Description:     transferCreditFields.Description,
	}

	transferCreditRequirements, marshalErr := json.Marshal(transferCredit)
	if marshalErr != nil {
		return fmt.Sprintf(`{"message": "marshal error", "error": "%s"}`, marshalErr.Error())
	}

	// This will get the endpoint from DB
	ServiceEP := util.GetServiceEP("CreditTransfer_igate", strings.ToLower(envRouting.Environment))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, ServiceEP, bytes.NewBuffer(transferCreditRequirements))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Merchant-ID", "QVBJMDAwMDU=")

	fmt.Println("REQUEST:", req)
	if err != nil {
		return fmt.Sprintf(`{"message": "http request error", "error": "%s"}`, err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf(`{"message": "client request error", "error": "%s"}`, err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Sprintf(`{"message": "reading body error", "error": "%s"}`, err.Error())
	}

	response := &igateModel.TransferCreditResponse{}
	if unmarshalErr := json.Unmarshal(body, response); unmarshalErr != nil {
		return fmt.Sprintf(`{"message": "unmarshal error", "error": "%s"}`, unmarshalErr.Error())
	}

	fmt.Println("RESPONSE:", response)
	return fmt.Sprintf(`{"serviceEP": "%s", "response": %v}`, ServiceEP, response)
}
