package routes

import (
	"backendmailingroom/config/middleware"
	office_controller "backendmailingroom/controller/office-contoller"
	subdirektorat_controller "backendmailingroom/controller/subdirektorat-controller"
	user_controller "backendmailingroom/controller/user-controller"

	"github.com/gofiber/fiber/v2"
)

func GetHome(ctx *fiber.Ctx) error {
	ipAddress := ctx.IP()
	if ipAddress == "" {
		ipAddress = "Unknown"
	}

	return ctx.JSON(fiber.Map{
		"ip_address": ipAddress,
	})
}

func UserRoutes(grp fiber.Router) (err error) {
	user := user_controller.NewUserController(UserRepository)

	grp.Get("/", GetHome)
	groupes := grp.Group("/user")
	groupes.Use(middleware.AuthMiddleware("kurir"))
	//User Routes

	groupes.Get("/getallusers", user.GetAllUsers)
	groupes.Get("/getuserbyid/:id", user.GetUserByID)
	groupes.Get("/getuserbyemail/:email", user.GetUserByEmail)
	groupes.Delete("/deleteuserbyid/:id", user.DeleteUserByID)
	groupes.Put("/updateuser/:id", user.UpdateUser)

	return
}

func AdminRoutes(grp fiber.Router) (err error) {
	admin := subdirektorat_controller.NewSubdirektoratController(SubdirektoratRepository)
	office := office_controller.NewOfficeController(OfficeRepository)
	user := user_controller.NewUserController(UserRepository)

	groupes := grp.Group("/admin")
	groupes.Use(middleware.AuthMiddleware("admin"))
	//Departemen Routes
	groupes.Post("/inputuser", user.InputUser)
	groupes.Post("/inputdepartemen", admin.InputSubdirektorat)
	groupes.Post("/inputoffice", office.InputOffice)
	groupes.Get("/getofficebyid/:id", office.GetOfficeByID)
	groupes.Get("/getalloffice", office.GetAllOffice)
	groupes.Get("/getofficebykota/:kota", office.GetOfficeByKota)
	groupes.Delete("/deleteofficebyid/:id", office.DeleteOfficeByID)
	groupes.Put("/updateoffice/:id", office.UpdateOffice)

	return
}

func PublicRoutes(grp fiber.Router) (err error) {
	user := user_controller.NewUserController(UserRepository)

	groupes := grp.Group("/public")
	//User Routes
	groupes.Post("/login", user.Login)

	return
}
