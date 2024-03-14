package routes

import (
	_ "sample/docs"
	"sample/handlers"
	"sample/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	// Import the http-swagger package
)

func SetupRoutes(app *fiber.App) {
	//swagger routes
	// app.Get("/InstaPay-docs*", swagger.HandlerDefault)
	app.Get("sample-code_docs/*", swagger.HandlerDefault)

	apiEndpoint := app.Group("/api")
	v1Endpoint := apiEndpoint.Group("/v1")

	fdsapEndpoint := v1Endpoint.Group("/ips/fdsap")
	//Login
	app.Post("/login", handlers.LoginHandler)
	// Add other routes as needed
	app.Post("/add", handlers.Add)
	// delete the rables//
	app.Delete("/delete/:id", handlers.DeleteUser)
	// Update yung tables
	app.Put("/update/:id", handlers.UpdateHandler)
	//Show spicific
	app.Get("/ShowUserId/:id", handlers.ShowUserIdHandler)
	//Show All
	app.Get("/showalls", handlers.ShowAllsHandler)
	///root
	app.Post("/gg", handlers.GG)
	// http
	app.Post("gg-copy", handlers.Https)
	/////tryinggg routess//
	app.Post("try", handlers.Sample)
	////instaPAy//
	app.Post("/insta", handlers.AdmnSignOnReq)

	app.Post("/insta-pay", handlers.AdmnSignOnReq)
	//////HAHAHHAHAH/////////
	app.Post("/signof", handlers.SignOffReq)

	fdsapEndpoint.Get("/read", handlers.CheckInstaPaySign)

	app.Post("/try12", handlers.SignOn)

	app.Post("/11", handlers.Notif_status)

	fdsapEndpoint.Post("/SignedOn", handlers.Pays)

	fdsapEndpoint.Get("/routines", handlers.GetOnlineRecords)

	fdsapEndpoint.Post("/credits", handlers.CreditsTransfer)

	app.Post("/transfer", handlers.TransCredit)
	fdsapEndpoint.Put("/Inquiry", handlers.InquiryTransferCredit)
	app.Post("/generate-token", handlers.Token)
	app.Post("/23", handlers.CreditTransferSending)
	app.Post("/24", handlers.TransferCreditProcess)
	app.Post("/balance", handlers.DeductBalance)

	app.Post("/sms", handlers.SendSMS)
	// app.Post("/Login", handlers.Manage)

	app.Post("/send-email", handlers.Email)      //Email
	app.Post("/feedbackId", handlers.FeedbackId) //feedbackID
	// app.Post("/feedback-Acc", handlers.Feedback) //feedbackACc
	app.Post("/Trace", handlers.Trace) //Trace
}
func AuthenticatedRoutes(app *fiber.App) {
	// Apply JWTMiddleware to these routes
	authenticated := app.Group("/authenticated", middleware.JWTMiddleware())
	authenticated.Get("/secure-endpoint", secureEndpoint)
}

func secureEndpoint(c *fiber.Ctx) error {
	// Handle the secure endpoint here
	return c.JSON(fiber.Map{
		"message": "You are authorized to access this endpoint!",
	})
}
