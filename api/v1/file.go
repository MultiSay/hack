package v1

import (
	"hack/internal/app/model"
	"net/http"
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
// @Router /v1/file/ [post]
func (h *Api) AddFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		f := &model.File{
			Name:     "NEW",
			CreateAt: time.Now(),
			Size:     0,
		}

		validate := f.Validate()
		if validate != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, validate.Error())
		}

		if err := h.store.File().Create(c.Request().Context(), f); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		//h.worker.Add(o)

		return c.JSON(http.StatusAccepted, f)
	}
}
