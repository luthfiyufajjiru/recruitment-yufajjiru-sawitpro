package handler

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/helpers"
	"github.com/labstack/echo/v4"
)

// (POST /registration)
func (s *Server) Register(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, helpers.DRNotImplemented)
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
