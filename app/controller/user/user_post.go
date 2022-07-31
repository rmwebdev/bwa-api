package user

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostUser godoc
// @Summary Create new user
// @Description Create new user
// @Param data body model.UserAPI true "User data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.UserData "User data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /users [post]
// @Tags Page
func PostUser(c *fiber.Ctx) error {
	api := new(model.UserAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB

	var data model.User

	lib.Merge(api, &data)

	fillValues(c, db, &data)
	// if err := db.Create(&data).Error; err != nil {
	// 	return lib.ErrorConflict(c, err)
	// }
	res := getUserByID(data.ID, db)

	return lib.OK(c, res)
}
