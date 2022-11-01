package delivery

import (
	"be12/mentutor/features/user"
	"be12/mentutor/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, usecase user.UsecaseInterface) {
	handler := &UserDelivery{
		userUsecase: usecase,
	}
	e.POST("/users", handler.PostData, middlewares.JWTMiddleware())
	// e.PUT("/users", handler.UpdateUser, middlewares.JWTMiddleware())
	// e.GET("/users", handler.GetByTokenJWT, middlewares.JWTMiddleware())
	// e.GET("/users/:id", handler.GetByIdWithJWT, middlewares.JWTMiddleware())
	// e.DELETE("/users", handler.DeleteMyAccount, middlewares.JWTMiddleware())

}

func (delivery *UserDelivery) PostData(c echo.Context) error {
	var dataRequest UserRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed To Bind",
		})
	}
	_, role := middlewares.ExtractToken(c)
	if role != "Admin" && role != "admin" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "you not have access",
		})
	}
	row, err := delivery.userUsecase.InsertData(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Register",
		})
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Register New User",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success Registers",
	})
}
