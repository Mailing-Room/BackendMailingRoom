package routes

import (
	departemen "backendmailingroom/repository/departemen"
	office "backendmailingroom/repository/office"
	user "backendmailingroom/repository/users"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(db *mongo.Client) {
	UserRepository = user.NewUser(db)
	DepartemenRepository = departemen.NewDepartemen(db)
	OfficeRepository = office.NewOffice(db)
}

func Router(app *fiber.App) (err error) {
	api := app.Group("/api")

	err = UserRoutes(api)
	if err != nil {
		return err
	}

	err = AdminRoutes(api)
	if err != nil {
		return err
	}

	err = PublicRoutes(api)
	if err != nil {
		return err
	}

	return
}
