package router

import (
	"net/http"
	"strconv"

	"example.com/todo-server/model"
	"github.com/labstack/echo/v4"
)

type PostBlogRequest struct {
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PutBlogRequest struct {
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

	blog, err = model.CreateBlog(params, *user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusOK, blog)
}

func GetBlogsHandler(c echo.Context) error {
	blogs, err := model.GetBlogs(c.QueryParam("user_id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusOK, blogs)
}

func PutBlogHandler(c echo.Context) error {
	var req PutBlogRequest

	err := c.Bind(&req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	var blog *model.Blog

	params := model.UpdateBlogParams{
		Title: req.Title,
		Content: req.Content,
	}

	var blogID int

	blogID, err = strconv.Atoi(c.Param("blogID"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	blog, err = model.PutBlog(blogID, params)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusOK, blog)
}

func DeleteBlogHandler(c echo.Context) error {
	blogID, err := strconv.Atoi(c.Param("blogID"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	err = model.DeleteBlog(blogID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.NoContent(http.StatusOK)
}

func GetBlogHandler(c echo.Context) error {
	blogID, err := strconv.Atoi(c.Param("blogID"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	var blog *model.Blog

	blog, err = model.GetBlog(blogID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusOK, blog)
}
