package response

import (
	"fmt"
	"github.com/ghiac/go-commons/log"
	"github.com/ghiac/social/internal/middlewares"
	"github.com/ghiac/social/pkg/access"
	"github.com/ghiac/social/pkg/text"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

func FromServerError(cc *middlewares.CustomContext, response access.ServerError) error {
	des := text.GetError(response.Response, text.Fa)
	if response.Description != nil {
		des = response.Description.English
	}
	fmt.Println(response.Error)
	return FromGeneralResponse(cc, GeneralResponse{
		Ok:          false,
		Description: des,
		Code:        response.Code,
		Error:       response.Response.String(),
	})
}

func ErrorNotFound(cc *middlewares.CustomContext) error {
	description := text.GetError(text.NotFound, text.Fa)
	return cc.JSON(http.StatusNotFound, GeneralResponse{
		Ok:          false,
		Description: description,
	})
}

func GetBadRequest(c echo.Context, description string) error {
	return c.JSON(http.StatusBadRequest, GeneralResponse{
		Ok:          false,
		Description: description,
	})
}

func FromWordLang(c *middlewares.CustomContext, wordLang text.WordLang, errorCode int) error {
	description := wordLang.Persian
	return FromGeneralResponse(c, GeneralResponse{
		Ok:          false,
		Description: description,
		Code:        errorCode,
	})
}

func FromGeneralResponse(c echo.Context, response GeneralResponse) error {
	return c.JSON(response.Code, response)
}

func GetErrorNotFound(c echo.Context, description string) error {
	return c.JSON(http.StatusNotFound, GeneralResponse{
		Ok:          false,
		Description: description,
	})
}

func GetUserNotFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, GeneralResponse{
		Ok:          false,
		Description: "کاربر پیدا نشد",
	})
}

func GetUnAuthorized(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, GeneralResponse{
		Ok:          false,
		Description: "UnAuthorized",
	})
}

func GetResult(c echo.Context, result interface{}) error {
	return c.JSON(http.StatusOK, GeneralResponse{
		Ok:     true,
		Result: result,
	})
}

func GetTrueSuccess(c echo.Context) error {
	return c.JSON(http.StatusOK, GeneralResponse{
		Ok:     true,
		Result: true,
	})
}

func GetTryAgain(c echo.Context, err error) error {
	log.Logger.WithFields(logrus.Fields{
		"method": "GetTryAgain", "message": "Internal Server Error", "path": c.Path(), "request": c.Request()}).
		Errorf("error: %v", err)
	return c.JSON(http.StatusInternalServerError, GeneralResponse{
		Ok:          false,
		Description: "دوباره تلاش کنید",
	})
}
