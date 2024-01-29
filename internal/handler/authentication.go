package handler

import (
	"fmt"
	"github.com/ghiac/social/internal/mysql"
	"github.com/ghiac/social/internal/repositories"
	"github.com/ghiac/social/internal/requests"
	"github.com/ghiac/social/internal/response"
	"github.com/ghiac/social/pkg/access"
	"github.com/ghiac/social/pkg/text"
	"github.com/ghiac/social/pkg/utils"
	"net/http"
)

func (h *Handler) Login(u requests.LoginRequest) (*response.LoginResponse, *access.ServerError) {
	loggedInUser, found := h.UserRepo.CheckAuth(u.Email, mysql.GetMD5Hash(u.Password))

	if !found || loggedInUser.ID == 0 {
		return nil, &access.ServerError{
			Response: text.UserOrPassInvalid,
			Code:     http.StatusBadRequest,
			Where:    "login method",
		}
	}
	token := utils.ComplexString(100)
	if sErr := h.SessionRepo.Add(&repositories.Session{
		UserID: loggedInUser.ID,
		Token:  token,
	}); sErr != nil {
		return nil, &access.ServerError{
			Response: text.TryAgain,
			Code:     http.StatusInternalServerError,
			Where:    "login method",
			Error:    sErr,
		}
	}
	result := response.LoginResponse{
		Token: token,
		User:  loggedInUser,
	}
	return &result, nil
}

func (h *Handler) UserInfo(userID int) (*repositories.User, *access.ServerError) {
	loggedInUser, found := h.UserRepo.Get(userID)
	if !found || loggedInUser.ID == 0 {
		return nil, &access.ServerError{
			Response: text.UserOrPassInvalid,
			Code:     http.StatusBadRequest,
			Where:    "UserInfo method",
		}
	}
	return &loggedInUser, nil
}

func (h *Handler) Register(u requests.LoginRequest) (*response.LoginResponse, *access.ServerError) {
	if len(u.Email) < 6 {
		return nil, access.GetError(text.NotEnoughChar)
	}
	if len(u.Password) < 6 {
		return nil, access.GetError(text.NotEnoughChar)
	}

	var duplicatedEmail = h.UserRepo.IsExistByEmail(u.Email)
	if duplicatedEmail {
		return nil, &access.ServerError{
			Response: text.InputDuplicated,
			Code:     http.StatusBadRequest,
			Where:    "Register method",
		}
	}

	randomUserID := h.UserRepo.GetRandomID()
	fmt.Println(randomUserID)
	user := repositories.User{
		Email:    u.Email,
		Password: mysql.GetMD5Hash(u.Password),
	}
	user.ID = uint(randomUserID)
	fmt.Println(user)
	var err = h.UserRepo.Add(&user)
	if err != nil {
		return nil, &access.ServerError{
			Response: text.TryAgain,
			Code:     http.StatusInternalServerError,
			Where:    "Register method",
			Error:    err,
		}
	}
	loginResponse, sErr := h.Login(requests.LoginRequest{
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		return nil, sErr
	}
	return loginResponse, nil
}
