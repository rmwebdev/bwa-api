package middleware

import (
	"api/app/services"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Oauth2Authentication authenticating oauth2 before calling next request
func Oauth2Authentication(c *fiber.Ctx) error {
	// do something here ...
	// dummy agent (development only)
	userID := c.Get("X-User-ID")
	db := services.DB

	if uid := getUserID(db, c); uid != "" {
		userID = uid
		c.Request().Header.Set("X-User-ID", userID)
	}

	if userID == "" {
		userID = viper.GetString("USER_ID")
		c.Request().Header.Set("X-User-ID", userID)
	}

	return c.Next()
}

func getUserID(db *gorm.DB, c *fiber.Ctx) (userID string) {
	authHeader := c.Get("authorization")
	if authHeader != "" {
		authHeaders := strings.Split(authHeader, ".")
		if len(authHeaders) == 3 && authHeaders[1] != "" {
			authData, err := base64.RawStdEncoding.DecodeString(authHeaders[1])
			if nil == err {
				mapData := map[string]interface{}{}
				err = json.Unmarshal(authData, &mapData)
				if nil == err {

					// if value, ok := mapData["sub"]; ok {
					// 	data := struct {
					// 		UserAccountID string
					// 		AgentID       string
					// 		CorporateID   string
					// 	}{}
					// 	db.Model(&model.UserAccount{}).
					// 		Select(`"user_account".id user_account_id, u.agent_id, c.corporate_id`).
					// 		Joins(`LEFT JOIN "agent_user" u ON u.user_account_id = "user_account".id`).
					// 		Joins(`LEFT JOIN "corporate_user" c ON c.user_account_id = "user_account".id`).
					// 		Where(`"user_account".email = ?`, value).
					// 		First(&data)
					// 	return data.UserAccountID, data.AgentID, data.CorporateID
					// }
				}
			}
		}
	}
	return ""
}
