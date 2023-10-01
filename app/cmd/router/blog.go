package router

import (
	"net/http"

	"example.com/todo-server/model"
	"github.com/labstack/echo/v4"
)

type PostBlogRequest struct {
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func PostBlogHandler(c echo.Context) error {
	var req PostBlogRequest

	err := c.Bind(&req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	var user *model.User

	user, err = model.GetUser(req.UserID)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User Not Found")
	}

	var blog *model.Blog

	params := model.CreateBlogParams{
		UserID:  user.ID,
		Title:   req.Title,
		Content: req.Content,
	}

	blog, err = model.CreateBlog(params)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusOK, blog)
}
