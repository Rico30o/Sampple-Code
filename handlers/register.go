package handlers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sample/db"
	"sample/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Add(c *fiber.Ctx) error {

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	db.DB.Create(&user)
	return c.XML(user)
}

// @Summary Delete a user by ID
// @Description Delete a specific user in the database by their ID.
// @Accept json
// @Produce json
// @Param id path int true "User ID to delete"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Router /delete/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var Reg models.User

	result := db.DB.Debug().Raw("DELETE FROM users WHERE id = ?", id).Scan(&Reg)

	if result != nil {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{

			"messege": "User Deleted"})

	}
	db.DB.Delete(&Reg)
	return c.JSON(fiber.Map{"message": "User Not Deleted"})

}

// UpdateUserByID godoc
// @Summary Update a user by ID
// @Description Update a specific user in the database by their ID.
// @Accept json
// @Produce json
// @Param id path int true "User ID to update"
// @Param user body models.User true "User object to update"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Router /update/{id} [put]
// Update Users
func UpdateHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var Reg models.User
	if err := db.DB.First(&Reg, id).Error; err != nil {

		return err
	}

	if err := c.BodyParser(&Reg); err != nil {

		return err
	}

	db.DB.Save(&Reg)
	return c.JSON(Reg)

}

func ShowUserIdHandler(c *fiber.Ctx) error {

	id := c.Params("id")
	var user models.User
	db.DB.Find(&user, id)
	return c.XML(user)

}

// UpdateUserByID godoc
// @Summary Update a user by ID
// @Description Update a specific user in the database by their ID.
// @Accept json
// @Produce json
// @Param id path int true "User ID to update"
// @Param user body models.User true "User object to update"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Router /update/{id} [put]

func ShowAllsHandler(c *fiber.Ctx) error {

	var user []models.User
	db.DB.Find(&user)
	return c.JSON(user)
}

// func LoginHandler(c *fiber.Ctx) error {

// 	var user, password models.User
// 	// var username, password string

// 	// if err := c.BodyParser(&user); err != nil {
// 	// 	return err

// 	// }

// 	fmt.Print("Enter username: ")
// 	fmt.Scan(&user)

// 	fmt.Print("Enter password: ")
// 	fmt.Scan(&password)

//		if login(user, password) {
//			fmt.Println("Login successful")
//		} else {
//			fmt.Println("Invalid username or password")
//			os.Exit(1)
//		}
//	}
func LoginHandler(c *fiber.Ctx) error {

	fmt.Println("username:")
	var username string
	fmt.Scanln(&username)

	fmt.Println("Password:")
	var password string
	fmt.Scanln(&password)

	var user models.User

	result := db.DB.First(&user, "username = ?", username)
	if result.Error != nil {
		fmt.Println("Invalid")
		os.Exit(1)
	}

	if password != user.Password {
		fmt.Println("Invalid")
		os.Exit(1)
	}

	response := models.LoginResponse{

		Message: fmt.Sprintf("Hello dear, %s!", user.Email),
	}

	xmlResponse, err := xml.MarshalIndent(response, "", " ")
	if err != nil {
		log.Fatal("Failed to generat XML response")
	}

	fmt.Println("Login Successful")
	fmt.Println(string(xmlResponse))

	db.DB.First(&xmlResponse)
	return c.XML(xmlResponse)
}

///////try//////

// // root ///
func GG(c *fiber.Ctx) error {
	Wtf := models.Wew{

		Name:     "gg",
		Lastname: "gg",

		Friend: models.Friend{

			Name:     "wew",
			Lastname: "wew",
		},
	}

	xmlInfo, err := xml.MarshalIndent(Wtf, "", " ")
	if err != nil {

		return c.Status(http.StatusInternalServerError).SendString("error generating XMl response")
	}
	c.Response().Header.Set("Context-Type", "application/xml")
	c.Set("Content-Type", "application/xml")
	return c.Send(xmlInfo)
}

func Https(c *fiber.Ctx) error {
	modes, err := xml.MarshalIndent(models.Wew{}, "", " ")
	if err != nil {
		return c.SendString(err.Error())
	}
	response, responseErr := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/gg", bytes.NewBuffer(modes))
	if responseErr != nil {
		return c.SendString(responseErr.Error())
	}
	response.Header.Set("Content-Type", "application/xml")
	client := &http.Client{}

	//send response
	resp, clientErr := client.Do(response)
	if clientErr != nil {
		return c.SendString(clientErr.Error())
	}
	defer resp.Body.Close()

	//response
	respBody, respErr := io.ReadAll(resp.Body)
	if respErr != nil {
		return c.SendString(respErr.Error())
	}
	c.Set("Content-Type", "application/xml")
	return c.Send(respBody)
}

/////root again///

func Sample(c *fiber.Ctx) error {
	Apply := models.Xample{

		Name:    "Rico Vergara",
		Address: "Sto.tomas Calauan Laguna",
		Email:   "@gmail.com",
	}
	xmlInfo, err := xml.MarshalIndent(Apply, "", " ")
	if err != nil {

		return c.Status(http.StatusInternalServerError).SendString("error generating XMl response")
	}
	c.Response().Header.Set("Context-Type", "application/xml")
	c.Set("Content-Type", "application/xml")
	return c.Send(xmlInfo)
}

/////boring po kasi ako!!

func AdmnSignOnReq(c *fiber.Ctx) error {
	Insta := models.InstaPay{

		Name3:           "Sign On Request",
		ISO_Description: "he Sign On Request message is used to initiate a sign on.",
		Product_Usage:   "The Instructing Agent uses the Sign-On Request Message to request sign on to InstaPay IPS.The initial status of the Instructing Agent is 'Signed-Off'. Sign-On is required when the previous status of Participant is 'Signed-Off'.The Instructing Agent sends the Sign-On message to InstaPay IPS when they first connect to InstaPay IPS, when re-connecting to InstaPay IPS after a planned or unplanned log-off. Signing-On indicates to the InstaPay IPS community that the bank is now available for receiving and sending payment requests.",
		Type:            "ignOnRequest",
		Occure:          "[1..1]",
	}
	c.Response().Header.Set("Context-Type", "application/xml")
	xmlInfo, err := xml.MarshalIndent(Insta, "", " ")
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("error generating XMl responcse")
	}
	c.Set("Context-Type", "application/xml")
	return c.Send(xmlInfo)

	// 	xmlInfo, err := xml.MarshalIndent(Wtf, "", " ")
	// 	if err != nil {

	// 		return c.Status(http.StatusInternalServerError).SendString("error generating XMl response")
	// 	}
	// 	c.Response().Header.Set("Context-Type", "application/xml")
	// 	c.Set("Content-Type", "application/xml")
	// 	return c.Send(xmlInfo)
	// }

}

func AdmnSignOffReqHttps(c *fiber.Ctx) error {

	modes, err := xml.MarshalIndent(models.InstaPay{}, "", " ")
	if err != nil {
		return c.SendString(err.Error())
	}
	response, responseErr := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/insta", bytes.NewBuffer(modes))
	if responseErr != nil {
		return c.SendString(responseErr.Error())
	}
	response.Header.Set("Content-Type", "application/xml")
	client := &http.Client{}

	//send response
	resp, clientErr := client.Do(response)
	if clientErr != nil {
		return c.SendString(clientErr.Error())
	}
	defer resp.Body.Close()

	//response
	respBody, respErr := io.ReadAll(resp.Body)
	if respErr != nil {
		return c.SendString(respErr.Error())
	}
	c.Set("Content-Type", "application/xml")
	return c.Send(respBody)
}

//// HAHAHHAHAHA//////

func SignOffReq(c *fiber.Ctx) error {

	instap := models.SignofInstap{

		Name4:            "Sign Off Request",
		ISO_Description1: "The Sign-Off Request message is used to initiate a Sign-Off.",
		Product_Usage1:   "An Instructing Agent uses the Sign-Off Request Message to perform a Sign-Off request to InstaPay IPS.An Instructing Agent sends the Sign-Off message to InstaPay IPS before bringing down their RealTime services for scheduled maintenance.",
		Type1:            "SignOffRequest",
		Occurence1:       "[1..1]",
	}
	c.Response().Header.Set("Context-Type", "application/xmzl")
	xmlInfo, err := xml.MarshalIndent(instap, "", " ")
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("error generating XMl responcse")
	}
	c.Response().Header.Set("Context-Type", "application/xml")
	c.Set("Content-Type", "application/xml")
	return c.Send(xmlInfo)
}
func GrpHd(c *fiber.Ctx) error {

	instap := models.Group_Header{

		Name5:            "Group Header",
		ISO_Description2: "Set of characteristics shared by all individual transactions included in the message",
		Type2:            "GrpHdr",
		Occurence2:       "[1..1]",
	}
	c.Response().Header.Set("Context-Type", "application/xml")
	xmlInfo, err := xml.MarshalIndent(instap, "", " ")
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("error generating XMl responcse")
	}
	c.Response().Header.Set("Context-Type", "application/xml")
	c.Set("Content-Type", "application/xml")
	return c.Send(xmlInfo)

}

// ///// task //////

func CheckInstaPaySign(c *fiber.Ctx) error {
	logFiles := []string{
		"logs/rick.log",
		// Add more file paths if needed
	}

	// Read contents of all log files
	allLogContents := make(map[string]string)

	for _, filePath := range logFiles {
		contents, err := os.ReadFile(filePath)
		if err != nil {
			// If there's an error reading a file, you can handle it or ignore it
			allLogContents[filePath] = fmt.Sprintf("Error reading file %s: %v\n", filePath, err)
		} else {
			allLogContents[filePath] = string(contents)
		}
	}

	// Extract and compare dates
	dates := make(map[string]string)

	for filePath, content := range allLogContents {
		date := extractDate(content)
		if date != "" {
			dates[filePath] = date
		}
	}

	// Check if all log files have the same date
	var result bool = true
	firstDate := "JJ"
	for _, date := range dates {
		if firstDate == "JJ" {
			firstDate = date
		} else if firstDate != date {
			result = false
			break
		}
	}

	// Return the result as a response
	return c.JSON(map[string]interface{}{
		"result":      result,
		"logContents": allLogContents,
		"dates":       dates,
	})
}

func extractDate(logContent string) string {
	// Implement your logic to extract the date from log content
	// This is just a placeholder, you may need to adjust it based on your log format
	if strings.Contains(logContent, "Date:") {
		return "2023-10-05" // Replace with the actual date extraction logic
	}
	return ""
}

// result := db.DB.First(&user, "username = ?", username)
// 	if result.Error != nil {
// 		fmt.Println("Invalid")
// 		os.Exit(1)
// 	}

func SignOn(c *fiber.Ctx) error {
	notificationsData := models.Notifications_Data{
		Is_signed_on: false,
		Remark:       "This account is Good",
		Authority:    "Rico",
		Create_at:    time.Now(), // This line is fine
	}

	// Insert data into the database
	result := db.DB.Debug().Exec("INSERT INTO signon_notification (is_signed_on, remark, authority, create_at) VALUES (?, ?, ?, ?)",
		notificationsData.Is_signed_on,
		notificationsData.Remark,
		notificationsData.Authority,
		notificationsData.Create_at,
	)
	if result.Error != nil {
		// Handle the error if the database query fails
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data inserted successfully",
	})
}

func Notif_status(c *fiber.Ctx) error {
	// Define the server status
	serverStatus := models.Logs_Notification{

		// Signed_on:    "",
		// Remarks:      "",
		// Signed_by:    "",
		// Signoff_date: "",
		// Signon_date:  "",
		// Signon_time:  "",
		// Create_at:    "",
	}

	// Set SignedOn and Remark based on time, considering lunch break
	currentTime := time.Now()
	if currentTime.Hour() == 12 {
		// Lunch break from 12:00 PM to 1:00 PM
		serverStatus.Signed_on = false
		serverStatus.Remarks = "Server is down (Lunch break)"
	} else if currentTime.Hour() < 12 || currentTime.Hour() >= 13 {
		// Server is up before 12:00 PM and after 1:00 PM
		serverStatus.Signed_on = true
		serverStatus.Remarks = "Server is up"
		serverStatus.Signon_time = currentTime
	} else {
		// Lunch break ongoing
		serverStatus.Signed_on = false
		serverStatus.Remarks = "Server is down (Lunch break)"
	}

	// Insert data into the database
	result := db.DB.Debug().Exec(`
		INSERT INTO ricologs (
			signed_on, remarks, signed_by, signoff_date, signoff_time, signon_date, signon_time, created_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		serverStatus.Signed_on,
		serverStatus.Remarks,
		serverStatus.Signed_by,
		serverStatus.Signoff_date,
		serverStatus.Signoff_time,
		serverStatus.Signon_date,
		serverStatus.Signon_time,
		serverStatus.Create_at,
	)

	if result.Error != nil {
		// Handle the error if the database query fails
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data inserted successfully",
	})
}

// Pays handles the POST request to insert data into the trytable.
//
// @Summary Check if the user is online or offline based on input parameters.
// @Description Inserts user data into the trytable and provides a response message.
// @ID Post-Pays
// @Accept json
// @Produce json
// @Param request body models.AnotherTry true "JSON request body"
// @Success 200 {object} models.AnotherTry
// @Failure 400 {object} models.ErrorResponse
// @Router /13 [post]
func Pays(c *fiber.Ctx) error {
	// Parse the JSON body from the request
	var requestBody struct {
		CustomSignedBy string `json:"custom_signed_by"`
		ExactDate      bool   `json:"Exactdate"`
		SignonDate     string `json:"signon_date"` // New field for custom signon_date
		// Add other fields as needed

	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Create the AnotherTry struct
	notif := models.AnotherTry{
		Signed_on:    requestBody.ExactDate, // Set Signed_on based on ExactDate
		Remarks:      "",                    // Initialize Remarks with an empty string
		Signed_by:    requestBody.CustomSignedBy,
		Signoff_date: time.Now(),
		Signoff_time: time.Now(),
		Create_at:    time.Now(),
	}

	// Check if Signon_date is provided in the request body
	if requestBody.SignonDate != "" {
		// Parse the custom Signon_date from the request body
		signonDate, err := time.Parse("2006-01-02", requestBody.SignonDate)
		if err == nil {
			// Update Signon_date if parsing is successful
			notif.Signon_date = signonDate
		} else {
			// Handle error if parsing fails
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid date format for signon_date",
			})
		}
	} else {
		// Use the current time if Signon_date is not provided
		notif.Signon_date = time.Now()
	}

	// Determine the Signed_on value based on ExactDate
	if !requestBody.ExactDate && notif.Signon_date.Before(time.Now()) {
		notif.Signed_on = true
	} else {
		notif.Signed_on = false
	}

	// Determine the Remarks value based on whether Signon_date is outdated
	if notif.Signed_on {
		notif.Remarks = "Your Online"
	} else {
		notif.Remarks = "Your Offline"
	}

	// Execute the SQL INSERT statement
	const timeFormat = "2006-01-02 15:04:05"

	result := db.DB.Debug().Exec(`
        INSERT INTO trytable (
            signed_on, remarks, signed_by, signoff_date, signoff_time, signon_date, signon_time, create_at
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		notif.Signed_on,
		notif.Remarks,
		notif.Signed_by,
		notif.Signoff_date.Format(timeFormat),
		notif.Signoff_time.Format(timeFormat),
		notif.Signon_date.Format(timeFormat),
		notif.Signon_time.Format(timeFormat),
		notif.Create_at.Format(timeFormat),
	)

	if result.Error != nil {
		// Handle the error if the database query fails
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data inserted successfully",
	})
}
func GetOnlineRecords(c *fiber.Ctx) error {
	// Channel to receive results from goroutines
	resultCh := make(chan models.AnotherTry, 10) // Adjust the buffer size as needed

	// Query the database concurrently
	go func() {
		// Query the database to fetch records where Signed_on is true
		rows, err := db.DB.Debug().Table("routines").Where("signed_on = ?", true).Rows()
		if err != nil {
			// Handle the error if the database query fails
			fmt.Println("Error fetching data:", err)
			close(resultCh)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var record models.AnotherTry
			if err := db.DB.ScanRows(rows, &record); err != nil {
				// Handle the error if scanning fails
				fmt.Println("Error scanning row:", err)
				continue
			}
			resultCh <- record
		}

		close(resultCh)
	}()

	// Create a slice to collect the fetched records
	var fetchedRecords []models.AnotherTry

	// Process the results
	for record := range resultCh {
		// Optionally process the fetched records
		fmt.Printf("Fetched record: %+v\n", record)
		// You can do something with the fetched records, for example, add them to a slice
		fetchedRecords = append(fetchedRecords, record)

		// Insert the fetched record into a different table, assuming you have a different model for insertion
		insertRecord := models.InsertedRecord{
			Signed_on: record.Signed_on,
			Signed_by: record.Signed_by,
			Create_at: record.Create_at,
			// Add other fields as needed
		}

		// Assuming you have a function to insert records
		if err := db.DB.Model(&insertRecord).Create(&insertRecord).Error; err != nil {
			fmt.Println("Error inserting record:", err)
			// Handle the error as needed
		}
	}

	// Return the fetched records as JSON in the response
	return c.JSON(fiber.Map{
		"message": "Data fetched and inserted successfully",
		"records": fetchedRecords,
	})
}
