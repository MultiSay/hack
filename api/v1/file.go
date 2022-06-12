package v1

import (
	"database/sql"
	"hack/internal/app/model"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

// AddFile Сохранить файл с датасетом
// AddFile godoc
// @Summary Сохранить файл с датасетом
// @Tags file
// @Description Получаем файл с фронта с новым датасетом для расчета
// @Accept mpfd
// @Produce json
// @Param	file	 formData  string true	"form data with file" minlength(1)
// @Success 200 {object} model.File
// @Failure 422 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /v1/file [post]
func (h *Api) AddFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse request body as multipart form data with 32MB max memory
		err := c.Request().ParseMultipartForm(32 << 20)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// Get file uploaded via Form
		file, fileHeader, err := c.Request().FormFile("file")
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		defer file.Close()

		f := model.File{
			Name:     fileHeader.Filename,
			CreateAt: time.Now(),
			Size:     fileHeader.Size,
			Status:   "PROCESSED",
		}

		validate := f.Validate()
		if validate != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, validate.Error())
		}

		f, err = h.store.File().Create(c.Request().Context(), f)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		localFile, err := os.Create(f.Name)
		if err != nil {
			return err
		}
		defer localFile.Close()

		// Copy the uploaded file data to the newly created file on the filesystem
		if _, err := io.Copy(localFile, file); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		h.worker.Add(f)

		return c.JSON(http.StatusCreated, f)
	}
}

// GetLastFile Получить состояние текущего файла с датасетом
// GetLastFile godoc
// @Summary Получить состояние текущего файла с датасетом
// @Tags file
// @Description Если файл есть в расчете то получим его состояние
// @Produce json
// @Success 200 {object} model.File
// @Failure 204 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /v1/file [get]
func (h *Api) GetLastFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := h.store.File().GetLast(c.Request().Context())
		if err != nil {
			if err == sql.ErrNoRows {
				return echo.NewHTTPError(http.StatusNoContent, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, file)
	}
}
