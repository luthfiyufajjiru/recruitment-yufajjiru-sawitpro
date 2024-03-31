package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/helpers"
	"github.com/SawitProRecruitment/UserService/helpers/errorIndex"
	"github.com/labstack/echo/v4"
)

// (POST /registration)
func (s *Server) Register(ctx echo.Context) error {
	user := generated.UserRegistrationRequest{
		Name:        ctx.FormValue("name"),
		Password:    ctx.FormValue("password"),
		PhoneNumber: ctx.FormValue("phone_number"),
	}

	err := helpers.RegistrationValidator(user)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.MessageResponse{Message: err.Error()})
	}

	resp, err := s.Repository.SetProfile(ctx.Request().Context(), user)
	if err != nil && errors.Is(err, errorIndex.ErrPhoneNumberExist) {
		return ctx.JSON(http.StatusConflict, generated.MessageResponse{
			Message: helpers.DRPhoneNumberExist,
		})
	} else if err != nil {
		return ctx.JSON(http.StatusInternalServerError, generated.MessageResponse{
			Message: helpers.DRInternalServerError,
		})
	}

	return ctx.JSON(http.StatusOK, resp)
}

// (POST /login)
func (s *Server) Login(ctx echo.Context) error {
	usr := generated.UserLoginRequest{
		Password:    ctx.FormValue("password"),
		PhoneNumber: ctx.FormValue("phone_number"),
	}
	if usr.Password == "" || usr.PhoneNumber == "" {
		return ctx.JSON(http.StatusForbidden, generated.MessageResponse{
			Message: helpers.DRForbidden,
		})
	}

	name, id, err := s.Repository.ComparePassword(ctx.Request().Context(), usr.PhoneNumber, usr.Password)
	if err != nil {
		return ctx.JSON(http.StatusForbidden, generated.MessageResponse{
			Message: helpers.DRForbidden,
		})
	}

	accessToken, refreshToken, err := helpers.GenJWTTokens(id, name)
	if err != nil {
		return ctx.JSON(http.StatusForbidden, generated.MessageResponse{
			Message: helpers.DRForbidden,
		})
	}

	return ctx.JSON(http.StatusOK, generated.JWTTokens{
		AccessToken:  &accessToken,
		RefreshToken: &refreshToken,
	})
}

// (GET /profile)
func (s *Server) GetProfile(ctx echo.Context) error {
	tokenStr := ctx.Request().Header.Get("authorization")
	idx := strings.Index(tokenStr, " ")
	if tokenStr == "" || idx < 0 {
		return ctx.JSON(http.StatusForbidden, generated.MessageResponse{
			Message: helpers.DRForbidden,
		})
	}
	tokenStr = tokenStr[idx+1:]

	claims, err := helpers.TokenCheck(ctx, tokenStr)
	if err != nil {
		return ctx.JSON(http.StatusForbidden, generated.MessageResponse{
			Message: helpers.DRForbidden,
		})
	}

	userId := int(claims[helpers.ClaimUserId].(float64))

	profile, err := s.Repository.GetProfile(ctx.Request().Context(), userId)
	if err != nil {
		ctx.JSON(http.StatusForbidden, generated.MessageResponse{
			Message: helpers.DRForbidden,
		})
	}

	return ctx.JSON(http.StatusOK, profile)
}

// (PATCH /profile)
func (s *Server) PatchProfile(ctx echo.Context) error {
	tokenStr := ctx.Request().Header.Get("authorization")
	idx := strings.Index(tokenStr, " ")
	if tokenStr == "" || idx < 0 {
		return ctx.JSON(http.StatusForbidden, generated.MessageResponse{
			Message: helpers.DRForbidden,
		})
	}
	tokenStr = tokenStr[idx+1:]

	claims, err := helpers.TokenCheck(ctx, tokenStr)
	if err != nil {
		return ctx.JSON(http.StatusForbidden, generated.MessageResponse{
			Message: helpers.DRForbidden,
		})
	}

	userId := int(claims[helpers.ClaimUserId].(float64))

	name := ctx.FormValue("name")
	pn := ctx.FormValue("phone_number")
	reqData := generated.UserProfilePresenter{
		Name:        &name,
		PhoneNumber: &pn,
	}

	err = s.Repository.UpdateProfile(ctx.Request().Context(), userId, reqData)
	if errors.Is(err, errorIndex.ErrPhoneNumberExist) {
		return ctx.JSON(http.StatusConflict, generated.MessageResponse{
			Message: helpers.DRPhoneNumberExist,
		})
	} else if err != nil {
		return ctx.JSON(http.StatusUnauthorized, generated.MessageResponse{
			Message: helpers.DRUnauthorized,
		})
	}

	return ctx.JSON(http.StatusOK, generated.MessageResponse{
		Message: helpers.DROk,
	})
}
