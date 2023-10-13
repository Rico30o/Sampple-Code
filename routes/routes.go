package routes

import (
	_ "sample/docs"
	"sample/handlers"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	//swagger routes
	app.Get("/InstaPay_docs/*", swagger.HandlerDefault)

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

	app.Get("/read", handlers.CheckInstaPaySign)
	app.Post("/try12", handlers.SignOn)
	app.Post("/11", handlers.Notif_status)
	app.Post("/13", handlers.Pays)

}
