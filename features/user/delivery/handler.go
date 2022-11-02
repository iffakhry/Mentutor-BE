package delivery

import (
	"be12/mentutor/features/user"
	"be12/mentutor/middlewares"
	"net/http"
	"strconv"

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
	e.PUT("/users/:id", handler.UpdateUser, middlewares.JWTMiddleware())
	e.GET("/profile", handler.GetByTokenJWT, middlewares.JWTMiddleware())
	e.GET("/users", handler.GetUser, middlewares.JWTMiddleware())
	e.GET("/users/:id", handler.GetByIdWithJWT, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", handler.DeleteMyAccount, middlewares.JWTMiddleware())

}
func (delivery *UserDelivery) DeleteMyAccount(c echo.Context) error {
	_, role := middlewares.ExtractToken(c)
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "param must be number",
		})
	}
	if role != "admin" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "you dont have access",
		})
	}
	row, err := delivery.userUsecase.Delete(id)
	if err != nil || row != 1 {
		return c.JSON(500, map[string]interface{}{
			"message": "failed wrong token",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"message": "success delete account",
	})
}
func (delivery *UserDelivery) UpdateUser(c echo.Context) error {

	var dataUpdate UpdateFormat
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error Bind data",
		})
	}

	var add user.Core
	if dataUpdate.Email != "" {
		add.Email = dataUpdate.Email
	}
	if dataUpdate.Name != "" {
		add.Name = dataUpdate.Name
	}
	if dataUpdate.Password != "" {
		add.Password = dataUpdate.Password
	}

	if dataUpdate.Images != "" {
		add.Images = dataUpdate.Images
	}
	id, _ := strconv.Atoi(c.Param("id"))

	token, role := middlewares.ExtractToken(c)
	add.ID = uint(id)
	if id == -1 {
		return c.JSON(400, map[string]interface{}{
			"message": "failed id not found",
		})
	}

	if role != "admin" && id != token {
		return c.JSON(400, map[string]interface{}{
			"message": "dont have access",
		})
	}

	row, err := delivery.userUsecase.PutDataId(add)
	if err != nil || row < 1 {
		return c.JSON(400, map[string]interface{}{
			"message": "failed not found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success update",
	})

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

func (delivery *UserDelivery) GetByTokenJWT(c echo.Context) error {
	idToken, _ := middlewares.ExtractToken(c)

	res, err := delivery.userUsecase.GetProfile(idToken)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "failed get profile",
		})
	}

	respon := FromCore(res)

	return c.JSON(200, map[string]interface{}{
		"message": "success get my profile",
		"data":    respon,
	})
}

func (delivery *UserDelivery) GetUser(c echo.Context) error {
	idToken, _ := middlewares.ExtractToken(c)
	if idToken < 1 {
		return c.JSON(400, map[string]interface{}{
			"message": "failed get all user",
		})

	}
	res, err := delivery.userUsecase.GetAlluser()
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "failed get profile",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success get my profile",
		"data":    toResponList(res),
	})
}

func (delivery *UserDelivery) GetByIdWithJWT(c echo.Context) error {

	idToken, _ := middlewares.ExtractToken(c)
	id, _ := strconv.Atoi(c.Param("id"))
	if id == -1 {
		return c.JSON(400, map[string]interface{}{
			"message": "failed id not found",
		})
	}

	res, err := delivery.userUsecase.GetDataId(id, idToken)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "failed get profile",
		})
	}

	respon := FromCore(res)

	return c.JSON(200, map[string]interface{}{
		"message": "success",
		"data":    respon,
	})

}
