package delivery

import (
	"be12/mentutor/config"
	"be12/mentutor/features/forum"
	"be12/mentutor/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type postHandler struct {
	PostUsecase forum.UseCaseInterface
}

func New(e *echo.Echo, ps forum.UseCaseInterface) {
	handler := &postHandler{
		PostUsecase: ps,
	}

	e.POST("/forum/me", handler.AddPosting(), middleware.JWT([]byte(config.SECRET_JWT))) // REFERENSI
	e.POST("/forum", handler.AddStatus(), middleware.JWT([]byte(config.SECRET_JWT)))
	e.GET("/forum", handler.SelectAll(), middleware.JWT([]byte(config.SECRET_JWT)))

}

func (ph *postHandler) AddPosting() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input Request
		id, _, _ := middlewares.ExtractToken(c)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Input From Client",
			})
		}

		cnv := ToDomain(input)
		_, errposts := ph.PostUsecase.AddPost(cnv, id)
		if errposts != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Input From Client",
			})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success add status",
			// "data":    cnv,
		})
	}
}

func (ph *postHandler) AddStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input Request
		id, _, _ := middlewares.ExtractToken(c)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Input From Client",
			})
		}

		cnv := ToDomain(input)
		// resp := ToDomainRes(cnv)
		res, errposts := ph.PostUsecase.AddStatus(cnv, id)
		if errposts != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Input From Client",
			})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success Add Status",
			"data":    ToResponse(res),
		})
	}
}

func (ph *postHandler) SelectAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _, _ := middlewares.ExtractToken(c)
		if id < 1 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Invalid Input From Client",
			})
		}
		res, err := ph.PostUsecase.GetAllPosts()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Input From Client",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all status",
			"data":    res,
		})

	}
}
