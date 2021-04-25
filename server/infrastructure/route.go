package infrastructure

import (
	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/tora0091/react-todo/server/config"
	"github.com/tora0091/react-todo/server/controller"
)

type Routing struct {
	DBConn *sql.DB
	Gin    *gin.Engine
	Port   string
}

func NewRouting(db *sql.DB) *Routing {
	r := &Routing{
		DBConn: db,
		Gin:    gin.Default(),
		Port:   config.ConfigureString("server.address"),
	}
	r.SetRouting()
	return r
}

func (r *Routing) SetRouting() {
	r.Gin.Use(corsSetting())

	r.Gin.GET("/todo-items", controller.NewTodoController(r.DBConn).GetItems)
	r.Gin.POST("/todo", controller.NewTodoController(r.DBConn).RegisterItem)
	r.Gin.DELETE("/todo", controller.NewTodoController(r.DBConn).DeleteItem)
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}

func corsSetting() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST", "GET", "PUT", "DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
	})
}
