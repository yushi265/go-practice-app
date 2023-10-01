package main

import (
	"example.com/todo-server/model"
	"example.com/todo-server/router"
	"github.com/labstack/echo/v4"
)

func main() {
	sqlDB := model.DBConnection()
	defer sqlDB.Close()
	e := echo.New()
	router.SetRouter(e)
}
