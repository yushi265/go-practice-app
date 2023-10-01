package router

import (
	"net/http"
	"strconv"

	"example.com/todo-server/model"
	"github.com/labstack/echo/v4"
)

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
}

func GetUsersHandler(c echo.Context) error {
	users, err := model.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusOK, users)
}

func CreateUserHandler(c echo.Context) error {
	var req CreateUserRequest

	err := c.Bind(&req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	var user *model.User

	user, err = model.CreateUser(req.Name, req.Email)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusOK, user)
}

func GetUserHandler(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	user, err := model.GetUser(userID)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUserHandler(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	var req UpdateUserRequest

	err = c.Bind(&req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	user, err := model.UpdateUser(userID, req.Name)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUserHandler(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	err = model.DeleteUser(userID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.NoContent(http.StatusOK)
}
