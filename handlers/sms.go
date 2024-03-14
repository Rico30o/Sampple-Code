package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sample/igateModel"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SendSMS(c *fiber.Ctx) error {
	// Parse JSON payload from request
	payload := &igateModel.MessagePayload{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Convert payload to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	Connection := "http://192.168.0.113:8000/api/public/v1/message/broadcast"

	// Create HTTP request
	req, err := http.NewRequest(http.MethodPost, Connection, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create the HTTP request",
		})
	}

	req.Header.Set("Content-Type", "application/json")

	// Create HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send SMS", "details": err.Error()})
	}
	defer resp.Body.Close()

	// Read response body
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": readErr.Error(),
		})
	}

	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {

		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"error":   "Request failed with status code " + strconv.Itoa(resp.StatusCode),
			"details": string(body),
		})
	}

	var Response igateModel.MessageResponse
	if unmarshalErr := json.Unmarshal(body, &Response); unmarshalErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": unmarshalErr.Error(),
		})
	}
	response := struct {
		Message string                     ` json:"message"`
		Header  http.Header                ` json:"header"`
		Data    igateModel.MessageResponse `json:"data"`
	}{
		Message: "success",
		Header:  resp.Header,
		Data:    Response,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
