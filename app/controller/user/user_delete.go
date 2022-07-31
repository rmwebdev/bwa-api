package user

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// DeleteUser godoc
// @Summary Delete User by id
// @Description Delete User by id
// @Param id path string true "User ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /users/{id} [delete]
// @Tags User
func DeleteUser(c *fiber.Ctx) error {
	db := services.DB

	var user model.User
	result2 := db.Model(&user).Where(`id = ?`, c.Params("id")).First(&user)
	if result2.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}
	db.Delete(&user)

	return lib.OK(c)
}
