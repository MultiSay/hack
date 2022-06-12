package v1

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetCompaignList Получить список Компаний
// GetCompaignList godoc
// @Summary Получить список Компаний
// @Tags compaign
// @Description Получить список Компаний
// @Produce json
// @Success 200 {object} []model.Compaign
// @Failure 204 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /v1/compaign [get]
func (h *Api) GetCompaignList() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := h.store.Compaign().GetList(c.Request().Context())
		if err != nil {
			if err == sql.ErrNoRows {
				return echo.NewHTTPError(http.StatusNoContent, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, file)
	}
}
