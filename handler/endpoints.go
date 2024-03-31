package handler

import (
	"errors"
	"net/http"

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
	return ctx.String(http.StatusNotImplemented, helpers.DRNotImplemented)
}

// (GET /profile)
func (s *Server) GetProfile(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, helpers.DRNotImplemented)
}

// (PATCH /profile)
func (s *Server) PatchProfile(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, helpers.DRNotImplemented)
}
