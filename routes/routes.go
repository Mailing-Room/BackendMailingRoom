package routes

import (
	departemen_controller "backendmailingroom/controller/departemen-controller"
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

	groupes := grp.Group("/admin")
	//Departemen Routes
	groupes.Post("/inputdepartemen", admin.InputDepartemen)

	return
}

func PublicRoutes(grp fiber.Router) (err error) {
	user := user_controller.NewUserController(UserRepository)

	groupes := grp.Group("/public")
	//User Routes
	groupes.Post("/login", user.Login)

	return
}
