package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"qiniu-1024-server/types"
)

func (h *Handler) PostRegister(c echo.Context) error {
	var req types.UserRegisterPayload
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}
	u, err := h.srv.UserRegister(c.Request().Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func (h *Handler) PostLogin(c echo.Context) error {
	var req types.UserLoginPayload
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}
	token, err := h.srv.UserLogin(c.Request().Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
