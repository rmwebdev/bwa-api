//go:build !integration
// +build !integration

package middleware

import (
	"testing"
)

func TestOauth2Authentication(t *testing.T) {
	// db := services.DBConnectTest()
	// app := fiber.New()
	// app.Use(Oauth2Authentication)
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	c.Set("X-User-ID", string(c.Request().Header.Peek("x-user-id")))
	// 	c.Set("X-Agent-ID", string(c.Request().Header.Peek("x-agent-id")))
	// 	c.Set("X-Corporate-ID", string(c.Request().Header.Peek("x-corporate-id")))
	// 	return c.Status(200).SendString("success")
	// })

	// language := model.Language{
	// 	LanguageAPI: model.LanguageAPI{
	// 		LanguageCode:       lib.Strptr("id"),
	// 		LanguageName:       lib.Strptr("indonesia"),
	// 		LanguageNativeName: lib.Strptr("indonesia"),
	// 		LanguageAlpha3Code: lib.Strptr("idn"),
	// 	},
	// }
	// db.Create(&language)

	// var agent model.Agent
	// agent = (*agent.Seed())[0]
	// db.Create(&agent)

	// user := &model.UserAccount{}
	// user.Seed()
	// db.Create(&user)

	// var agentUser model.AgentUser
	// agentUser.AgentID = agent.ID
	// agentUser.UserAccountID = user.ID
	// db.Create(&agentUser)

	// var corporateUser model.CorporateUser
	// corporateUser.UserAccountID = user.ID
	// corporateUser.CorporateID = lib.UUIDPtr(uuid.New())
	// db.Create(&corporateUser)

	// agentLanguage := model.AgentLanguage{
	// 	AgentID: agent.ID,
	// 	AgentLanguageAPI: model.AgentLanguageAPI{
	// 		LanguageID: language.ID,
	// 		IsPrimary:  lib.Boolptr(true),
	// 	},
	// }

	// db.Create(&agentLanguage)
	// viper.Set("AGENT_ID", agent.ID.String())
	// viper.Set("USER_ID", user.ID.String())

	// headers := map[string]string{
	// 	"authorization": ".eyJhdWQiOltdLCJjbGllbnRfaWQiOiJteS1jbGllbnQtaWQiLCJleHAiOjE2NjAzNzgyNDksImV4dCI6e30sImlhdCI6MTY1Nzc4NjI0OSwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo0NDQ0LyIsImp0aSI6IjU1YWMyODI2LWM5NGYtNGE5Yy1iZTQ5LWUwN2U3NmMwZDI4NSIsIm5iZiI6MTY1Nzc4NjI0OSwic2NwIjpbIm9wZW5pZCIsIm9mZmxpbmUiXSwic3ViIjoiYWRtaW5AbWFpbC5jb20ifQ.",
	// }

	// request := httptest.NewRequest("GET", "/", nil)
	// for h := range headers {
	// 	request.Header.Set(h, headers[h])
	// }

	// response, err := app.Test(request)
	// utils.AssertEqual(t, nil, err, "sending request")
	// utils.AssertEqual(t, 200, response.StatusCode, "oauth2 authenticated")
	// utils.AssertEqual(t, user.ID.String(), response.Header.Get("x-user-id"), "x-user-id header")
	// utils.AssertEqual(t, agent.ID.String(), response.Header.Get("x-agent-id"), "x-agent-id header")
	// utils.AssertEqual(t, corporateUser.CorporateID.String(), response.Header.Get("x-corporate-id"), "x-corporate-id header")

	// request2 := httptest.NewRequest("GET", "/", nil)
	// response, err = app.Test(request2)
	// utils.AssertEqual(t, nil, err, "sending request")
	// utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
}
