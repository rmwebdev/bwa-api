package user

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetUserID godoc
// @Summary Get a User by id
// @Description Get a User by id
// @Param Accept-Language header string false "2 character language code"
// @Param id path string true "User ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.User User data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /users/{id} [get]
// @Tags User
func GetUserID(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	data := getUserByID(&id, db)
	if nil == data {
		return lib.ErrorNotFound(c)
	}
	return lib.OK(c, data)
}

func getUserByID(id *uuid.UUID, db *gorm.DB) *model.User {
	var data model.User
	res := db.Model(&model.User{}).
		Select(`
		"user".id,
		"user".name,
		"user".occupation,
		"user".email,
		"user".avatar_file_name,
		"user".token`).
		Where(`"user".id = ?`, id).
		First(&data)
	if res.RowsAffected == 1 {
		return &data
	}

	return nil
}
