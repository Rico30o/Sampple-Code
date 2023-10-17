package models

import (
	"encoding/xml"
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
)
