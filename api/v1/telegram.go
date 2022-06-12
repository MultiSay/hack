package v1

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetTelegramList Получить список телеграм каналов
// GetTelegramList godoc
// @Summary Получить список телеграм каналов
// @Tags telegram
// @Description Получить список телеграм каналов
// @Produce json
// @Success 200 {object} []model.Telegram
// @Failure 204 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /v1/lead [get]
func (h *Api) GetTelegramList() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := h.store.Telegram().GetList(c.Request().Context())
		if err != nil {
			if err == sql.ErrNoRows {
				return echo.NewHTTPError(http.StatusNoContent, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, file)
	}
}
