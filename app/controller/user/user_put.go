package user

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PutCorporateEmployee Put CorporateEmployee
// @Summary Update User Profile for current session
// @Description Update User Profile for current session
// @Param data body model.UserData true "Employee data"
// @Param id path string true "User ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.UserData "User data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /users/{id} [put]
// @Tags User
func PutUser(c *fiber.Ctx) error {
	rid, _ := uuid.Parse(c.Params("id"))
	if rid == uuid.Nil {
		return lib.ErrorBadRequest(c, "invalid id format")
	}

	id := &rid

	db := services.DB

	api := new(model.UserAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	data := getUserByID(id, db)
	if nil == data {
		return lib.ErrorNotFound(c, "not found")
	}

	lib.Merge(api, data)

	values := fillValues(c, db, data)

	re := regexp.MustCompile(`[0-9]+$`)
	for k, v := range values {
		table := re.ReplaceAll([]byte(k), []byte(``))
		db.Table(string(table)).Where(`id = ?`, v["id"]).Updates(v["values"])
	}

	return lib.OK(c, data)
}

func fillValues(c *fiber.Ctx, db *gorm.DB, data *model.User) map[string]map[string]interface{} {

	createUser(db, data)

	values := map[string]map[string]interface{}{}
	if nil != data.ID {
		values["user"] = map[string]interface{}{
			"id": data.ID,
			"values": map[string]interface{}{
				"name":             data.Name,
				"email":            data.Email,
				"occupation":       data.Occupation,
				"avatar_file_name": data.AvatarFileName,
				"role":             data.Role,
			},
		}

	}

	return values
}

func createUser(db *gorm.DB, data *model.User) {
	if nil != data.ID {
		// keep existing, do not create new one
		return
	}

	user := model.User{}
	user.Name = data.Name
	user.Email = data.Email
	user.Occupation = data.Occupation
	user.AvatarFileName = data.AvatarFileName
	user.Role = data.Role
	db.Create(&user)

}
