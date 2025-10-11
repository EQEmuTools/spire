package request

import (
	"github.com/EQEmuTools/spirerererere/internal/models"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) models.User {
	user, ok := c.Get("user").(models.User)
	if !ok {
		return models.User{}
	}

	return user
}
