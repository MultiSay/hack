package v1

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetLeadList Получить список Лидов
// GetLeadList godoc
// @Summary Получить список Лидов
// @Tags file
// @Description Получить список Лидов
// @Produce json
// @Success 200 {object} []model.Lead
// @Failure 204 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /v1/lead [get]
func (h *Api) GetLeadList() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := h.store.Lead().GetList(c.Request().Context())
		if err != nil {
			if err == sql.ErrNoRows {
				return echo.NewHTTPError(http.StatusNoContent, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, file)
	}
}
