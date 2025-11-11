package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/toffettl/explicAI/internal/application"
	"github.com/toffettl/explicAI/internal/infrastructure/errors"
	"github.com/toffettl/explicAI/internal/infrastructure/log"
)

type ExplicaServer struct {
}

func NewExplicaServer() *ExplicaServer {
	return &ExplicaServer{}
}

func (api *ExplicaServer) Register(server *echo.Echo) {
	server.POST("/upload", api.Upload)
	server.GET("/summaries", api.ListSummaries)
	server.GET("/summaries/:externalId", api.GetSummaryByExternalId)
	server.GET("/summaries/:externalId", api.DeleteSummaryByExternalId)
}

func (api *ExplicaServer) Upload(c echo.Context) error {
	ctx := c.Request().Context()
	_, err := api.getFileFromRequest(ctx, c)
	if err != nil {
		return errors.Handle(c, err)
	}

	//TODO: init flow

	return c.JSON(http.StatusCreated, nil)
}

func (api *ExplicaServer) ListSummaries(c echo.Context) error {
	//ctx := c.Request().Context()
	// init get flow
	return c.JSON(http.StatusOK, nil)
}

func (api *ExplicaServer) GetSummaryByExternalId(c echo.Context) error {
	//ctx := c.Request().Context()
	//externalId := c.Param("externalId")

	//parsedExternalId, err := uuid.Parse(externalId)
	//if err != nil {
	//	return echo.ErrBadRequest
	//}

	// TODO get flow

	return c.JSON(http.StatusOK, nil)
}

func (api *ExplicaServer) DeleteSummaryByExternalId(c echo.Context) error {
	//ctx := c.Request().Context()
	//externalId := c.Param("externalId")

	//parsedExternalId, err := uuid.Parse(externalId)
	//if err != nil {
	//	return echo.ErrBadRequest
	//}

	// TODO delete flow
	return c.JSON(http.StatusOK, map[string]string{
		"message": "sumary has been removed",
	})
}

func (api *ExplicaServer) getFileFromRequest(ctx context.Context, c echo.Context) ([]byte, error) {
	file, err := c.FormFile("file")
	if err != nil {
		log.LogError(ctx, "missing file", err)
		return nil, application.MissingFile
	}

	allowedExtensions := map[string]bool{
		".mp3":  true,
		".mp4":  true,
		".mpeg": true,
		".mpga": true,
		".m4a":  true,
		".wav":  true,
		".webm": true,
	}

	fileExtension := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[fileExtension] {
		return nil, application.InvalidFile
	}

	src, err := file.Open()
	if err != nil {
		log.LogError(ctx, "fail to open file", err)
		return nil, application.FailedReadFile
	}
	defer src.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, src); err != nil {
		log.LogError(ctx, "fail to read file", err)
		return nil, application.FailedReadFile
	}

	return buf.Bytes(), nil
}
