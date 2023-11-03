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
	err = h.srv.ActionService.PlayVideo(c.Request().Context(), uid, vidNum)
	if err != nil {
		return err
	}
	return c.NoContent(200)
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
	err = h.srv.ActionService.LikeVideo(c.Request().Context(), uid, vidNum)
	if err != nil {
		return err
	}
	return c.NoContent(200)
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
	err = h.srv.ActionService.CollectVideo(c.Request().Context(), uid, vidNum)
	if err != nil {
		return err
	}
	return c.NoContent(200)
}

// PostCommentVideo
func (h *Handler) PostCommentVideo(c echo.Context) error {
	uid, err := xecho.CurUserID(c)
	if err != nil {
		return err
	}
	vid := c.Param("id")
	vidNum, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		return xerr.New(400, "InvalidParam", "invalid id")
	}
	var req struct {
		Content string `json:"content"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	err = h.srv.ActionService.CommentVideo(c.Request().Context(), uid, vidNum, req.Content)
	if err != nil {
		return err
	}
	return c.NoContent(200)
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
	err = h.srv.ActionService.UnLikeVideo(c.Request().Context(), uid, vidNum)
	if err != nil {
		return err
	}
	return c.NoContent(200)
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
	err = h.srv.ActionService.UnCollectVideo(c.Request().Context(), uid, vidNum)
	if err != nil {
		return err
	}
	return c.NoContent(200)
}
