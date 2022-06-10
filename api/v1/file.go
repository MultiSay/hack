package v1

import (
	"hack/internal/app/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

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
