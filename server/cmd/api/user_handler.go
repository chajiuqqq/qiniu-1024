package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"qiniu-1024-server/types"
	"qiniu-1024-server/utils/xecho"
	"qiniu-1024-server/utils/xerr"
	"strconv"
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
	token, u, err := h.srv.UserLogin(c.Request().Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{"token": token, "user": u})
}
func (h *Handler) GetUser(c echo.Context) error {
	id := c.Param("id")
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return xerr.New(400, "InvalidParam", "invalid id")
	}
	u, err := h.srv.UserDetail(c.Request().Context(), idNum)
	if err != nil {
		return err
	}
	return c.JSON(200, u)
}

func (h *Handler) PostUserAction(c echo.Context) error {
	from, err := xecho.CurUserID(c)
	if err != nil {
		return err
	}
	to := c.Param("id")
	toNum, err := strconv.ParseInt(to, 10, 64)
	if err != nil {
		return xerr.New(400, "InvalidParam", "invalid id")
	}

	action := c.QueryParam("action")
	if action == "" {
		return xerr.New(400, "InvalidParam", "invalid action")
	}
	u, err := h.srv.PostUserAction(c.Request().Context(), from, toNum, action)
	if err != nil {
		return err
	}
	return c.JSON(200, u)
}

func (h *Handler) GetCurUser(c echo.Context) error {
	uid, err := xecho.CurUserID(c)
	if err != nil {
		return err
	}
	u, err := h.srv.UserDetail(c.Request().Context(), uid)
	if err != nil {
		return err
	}
	return c.JSON(200, u)
}
