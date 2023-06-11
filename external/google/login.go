package google

import (
	"context"
	"log"
	"net/http"

	"github.com/MuhAndriJP/personal-practice.git/helper"
	"golang.org/x/oauth2"
	userInfo "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"

	"github.com/labstack/echo/v4"
)

type Google struct {
}

func (g *Google) HandleGoogleLogin(c echo.Context) (err error) {
	url := GoogleOauthConfig().AuthCodeURL("state-token",
		oauth2.AccessTypeOffline,
		oauth2.ApprovalForce,
		oauth2.SetAuthURLParam("prompt", "select_account"))

	log.Println("Redirect URL Google Auth", url)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (g *Google) HandleGoogleCallback(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	code := c.QueryParam("code")
	token, err := GoogleOauthConfig().Exchange(ctx, code)
	if err != nil {
		return
	}

	client := GoogleOauthConfig().Client(ctx, token)
	userInfo, err := getUserInfo(ctx, client)
	if err != nil {
		return err
	}

	// CHECK DB IF USER STORED TO DB
	// userDB, err := g.user.GetUserByEmail(ctx, &pb.GetUserByEmailRequest{Email: userInfo.Email})
	// if err != nil {
	// 	return
	// }

	// if userDB.Email == "" {
	// 	req := pb.RegisterUserRequest{
	// 		Name:  userInfo.Name,
	// 		Email: userInfo.Email,
	// 		Token: token.AccessToken,
	// 	}
	// 	bytes, _ := json.Marshal(userInfo)
	// 	_ = json.Unmarshal(bytes, &req)

	// 	_, err = g.user.RegisterUser(ctx, &req)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	log.Println("Google Auth Register", &req)
	// 	resp := &helper.Response{
	// 		Code:    helper.SuccessCreated,
	// 		Message: helper.StatusMessage[helper.SuccessCreated],
	// 		Data: map[string]interface{}{
	// 			"token": token.AccessToken,
	// 		},
	// 	}

	// 	return c.JSON(helper.HTTPStatusFromCode(helper.SuccessCreated), resp)
	// }

	log.Println("Google Auth Login", userInfo)
	return c.JSON(helper.HTTPStatusFromCode(helper.Success), &helper.Response{
		Code:    helper.Success,
		Message: helper.StatusMessage[helper.Success],
		Data: map[string]interface{}{
			"token": token.AccessToken,
		},
	})
}

func getUserInfo(ctx context.Context, client *http.Client) (*userInfo.Userinfo, error) {
	userService, err := userInfo.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	userInfo, err := userService.Userinfo.Get().Do()
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

func NewGoogleAuth() *Google {
	return &Google{}
}
