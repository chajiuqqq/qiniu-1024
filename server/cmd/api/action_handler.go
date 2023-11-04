package main

import (
	"github.com/labstack/echo/v4"
	"qiniu-1024-server/utils/xecho"
	"qiniu-1024-server/utils/xerr"
	"strconv"
)

func (h *Handler) PostPlayVideo(c echo.Context) error {
	uid, err := xecho.CurUserID(c)
	if err != nil {
		uid = -1
	}
	vid := c.Param("id")
	vidNum, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		return xerr.New(400, "InvalidParam", "invalid id")
	}
	v, err := h.srv.ActionService.PlayVideo(c.Request().Context(), uid, vidNum)
	if err != nil {
		return err
	}
	return c.JSON(200, v)
}

// PostLikeVideo
func (h *Handler) PostLikeVideo(c echo.Context) error {
	uid, err := xecho.CurUserID(c)
	if err != nil {
		return err
	}
	vid := c.Param("id")
	vidNum, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		return xerr.New(400, "InvalidParam", "invalid id")
	}
	v, err := h.srv.ActionService.LikeVideo(c.Request().Context(), uid, vidNum)
	if err != nil {
		return err
	}
	return c.JSON(200, v)
}

// PostCollectVideo
func (h *Handler) PostCollectVideo(c echo.Context) error {
	uid, err := xecho.CurUserID(c)
	if err != nil {
		return err
	}
	vid := c.Param("id")
	vidNum, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		return xerr.New(400, "InvalidParam", "invalid id")
	}
	v, err := h.srv.ActionService.CollectVideo(c.Request().Context(), uid, vidNum)
	if err != nil {
		return err
	}
	return c.JSON(200, v)
}

// PostCommentVideo
func (h *Handler) PostCommentVideo(c echo.Context) error {
	uid, err := xecho.CurUserID(c)
	if err != nil {
		return err
	}
	var req struct {
		Content string `json:"content" validate:"required"`
		ID      int64  `param:"id" validate:"required"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}
	v, err := h.srv.ActionService.CommentVideo(c.Request().Context(), uid, req.ID, req.Content)
	if err != nil {
		return err
	}
	return c.JSON(200, v)
}

// DeleteLikeVideo
func (h *Handler) DeleteLikeVideo(c echo.Context) error {
	uid, err := xecho.CurUserID(c)
	if err != nil {
		return err
	}
	vid := c.Param("id")
	vidNum, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		return xerr.New(400, "InvalidParam", "invalid id")
	}
	v, err := h.srv.ActionService.UnLikeVideo(c.Request().Context(), uid, vidNum)
	if err != nil {
		return err
	}
	return c.JSON(200, v)
}

// DeleteCollectVideo
func (h *Handler) DeleteCollectVideo(c echo.Context) error {
	uid, err := xecho.CurUserID(c)
	if err != nil {
		return err
	}
	vid := c.Param("id")
	vidNum, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		return xerr.New(400, "InvalidParam", "invalid id")
	}
	v, err := h.srv.ActionService.UnCollectVideo(c.Request().Context(), uid, vidNum)
	if err != nil {
		return err
	}
	return c.JSON(200, v)
}
