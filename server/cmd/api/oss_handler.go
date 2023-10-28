package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"qiniu-1024-server/model"
	"qiniu-1024-server/service"
	"qiniu-1024-server/utils/oss"
)

func (h *Handler) OssVideoCallback(c echo.Context) error {
	pass, err := h.srv.Oss.VerifyCallback(c.Request())
	if err != nil {
		return err
	}
	if !pass {
		return c.NoContent(403)
	}

	key := c.Param("key")
	vid, err := service.GetVideoIDFromKey(key)
	if err != nil {
		return fmt.Errorf("get video id from key failed: %w", err)
	}
	err = h.srv.VideoStatusUpdate(c.Request().Context(), vid)
	if err != nil {
		h.sugar.Warnln("VideoStatusUpdate err", "vid", vid, "error", err)
		return err
	}
	return c.NoContent(200)
}

func (h *Handler) OssTaskCallback(c echo.Context) error {
	var ossTask oss.TaskCallbackBody
	if err := c.Bind(&ossTask); err != nil {
		return err
	}
	vid, err := service.GetVideoIDFromKey(ossTask.Input.KodoFile.Key)
	if err != nil {
		return fmt.Errorf("get video id from key failed: %w", err)
	}
	var coverStatus = model.CoverStatusUploading
	if ossTask.Code == 0 {
		coverStatus = model.CoverStatusSuccess
	}
	if ossTask.Code == 3 {
		coverStatus = model.CoverStatusFailed
	}
	err = h.srv.VideoCoverStatusUpdate(c.Request().Context(), vid, coverStatus)
	if err != nil {
		h.sugar.Warnln("VideoCoverStatusUpdate err", "vid", vid, "error", err)
		return err
	}
	return c.NoContent(200)
}
