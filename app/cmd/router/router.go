package router

import (
	"os"

	"github.com/labstack/echo/v4/middleware"

	_ "net/http"

	"github.com/labstack/echo/v4"
)

//ルーティングを設定する関数 引数はecho.echo型のc であり、戻り値はerror型である
func SetRouter(e *echo.Echo) error {

	// 諸々の設定(*1)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header}\n",
		Output: os.Stdout,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// APIを書く場所
	api := e.Group("/api")
	{
		apiTasks := api.Group("/tasks")
		{
			apiTasks.GET("", GetTasksHandler)
			apiTasks.POST("", AddTaskHandler)
			apiTasks.PUT("/:taskID", ChangeFinishedTaskHandler)
			apiTasks.DELETE("/:taskID", DeleteTaskHandler)
		}
		apiUsers := api.Group("/users")
		{
			apiUsers.GET("", GetUsersHandler)
			apiUsers.POST("", CreateUserHandler)
			apiUsers.GET("/:userID", GetUserHandler)
			apiUsers.PUT("/:userID", UpdateUserHandler)
			apiUsers.DELETE("/:userID", DeleteUserHandler)
		}
	}

	// 8000番のポートを開く(*2)
	err := e.Start(":8080")
	return err
}
