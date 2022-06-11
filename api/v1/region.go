package v1

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetRegionPredictList Получить список городов с придиктивной аналитикой
// GetRegionPredictList godoc
// @Summary Получить список городов с придиктивной аналитикой
// @Tags region
// @Description Результат работы модели загружаем в этом методе
// @Produce json
// @Success 200 {object} []model.RegionPredict
// @Success 204 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /v1/region/predict [get]
func (h *Api) GetRegionPredictList() echo.HandlerFunc {
	return func(c echo.Context) error {

		r, err := h.store.Region().PredictList(c.Request().Context())
		if err != nil {
			if err == sql.ErrNoRows {
				return echo.NewHTTPError(http.StatusNoContent, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, r)
	}
}
