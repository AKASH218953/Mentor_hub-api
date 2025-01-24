package rotues

import (
	"mentor_hub-api/controllers"

	"github.com/labstack/echo"
)

func Login(e *echo.Echo) {

	e.POST("/login", controllers.Login)

}
