package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"qiniu-1024-server/model"
	"qiniu-1024-server/service"
	"qiniu-1024-server/types"
	"qiniu-1024-server/utils/xecho"
	"qiniu-1024-server/utils/xerr"
	"strconv"
	"strings"
)

func (h *Handler) UploadFile(c echo.Context) error {
	uid, err := xecho.CurUserID(c)
	if err != nil {
		return err
	}
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// 检查文件扩展名是否为 .mp4
	if !strings.HasSuffix(strings.ToLower(file.Filename), ".mp4") {
		return xerr.New(400, "InvalidExtension", "invalid file extension")
	}

	// read file
	src, err := file.Open()
	if err != nil {
		h.sugar.Warnln("Open file err", "filename", file.Filename, "error", err)
	}
	defer src.Close()

	videoNum, err := h.srv.GetVideoSeq(c.Request().Context())
	if err != nil {
		return err
	}

	vid := service.GenVideoID(videoNum)
	var key = fmt.Sprintf("%d.mp4", vid)

	// get bytes from file
	fileBytes, err := io.ReadAll(src)
	if err != nil {
		h.sugar.Warnln("ReadAll file err", "filename", file.Filename, "error", err)
	}
	go func() {
		// upload
		_, err = h.srv.Oss.ByteUpload(fileBytes, key)
		if err != nil {
			h.sugar.Warnln("Oss ByteUpload err", "filename", file.Filename, "error", err)
		}
		// change status
		err = h.srv.VideoStatusUpdate(context.Background(), vid, model.VideoStatusNew)
		if err != nil {
			h.sugar.Warnln("VideoStatusUpdate err", "vid", vid, "error", err)
		}
	}()

	_, err = h.srv.PreSaveVideo(c.Request().Context(), uid, vid)
	if err != nil {
		return err
	}
	return c.JSON(200, echo.Map{
		"vid": vid,
		"key": key,
		"url": h.srv.Oss.ResourceUrl(key),
	})
}
func (h *Handler) GetMainCategories(c echo.Context) error {
	data, err := h.srv.MainCategories(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(200, data)
}
func (h *Handler) GetMainVideos(c echo.Context) error {
	param := c.QueryParam("category_id")
	cid, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return xerr.New(400, "InvalidParam", "invalid category_id")
	}
	data, err := h.srv.MainVideos(c.Request().Context(), cid)
	if err != nil {
		return err
	}
	return c.JSON(200, data)
}

func (h *Handler) PostMainVideo(c echo.Context) error {
	uid, err := xecho.CurUserID(c)
	if err != nil {
		return err
	}
	var req types.MainVideoSubmit
	if err := c.Bind(&req); err != nil {
		return err
	}
	data, err := h.srv.SaveVideo(c.Request().Context(), uid, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}
