package routes

import (
	departemen_controller "backendmailingroom/controller/departemen-controller"
	office_controller "backendmailingroom/controller/office-contoller"
	user_controller "backendmailingroom/controller/user-controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(grp fiber.Router) (err error) {
	user := user_controller.NewUserController(UserRepository)

	groupes := grp.Group("/user")
	//User Routes
	groupes.Post("/inputuser", user.InputUser)
	groupes.Get("/getallusers", user.GetAllUsers)
	groupes.Get("/getuserbyid/:id", user.GetUserByID)
	groupes.Get("/getuserbyemail/:email", user.GetUserByEmail)
	groupes.Delete("/deleteuserbyid/:id", user.DeleteUserByID)
	groupes.Put("/updateuser/:id", user.UpdateUser)

	return
}

func AdminRoutes(grp fiber.Router) (err error) {
	admin := departemen_controller.NewDepartemenController(DepartemenRepository)
	office := office_controller.NewOfficeController(OfficeRepository)

	groupes := grp.Group("/admin")
	//Departemen Routes
	groupes.Post("/inputdepartemen", admin.InputDepartemen)
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
