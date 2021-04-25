package infrastructure

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Item struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type Routing struct {
	DBConn *sql.DB
	Gin    *gin.Engine
	Port   string
}

func NewRouting(db *sql.DB) *Routing {
	r := &Routing{
		DBConn: db,
		Gin:    gin.Default(),
		Port:   "localhost:9090",
	}
	r.SetRouting()
	return r
}

func (r *Routing) SetRouting() {
	r.Gin.Use(cors.New(cors.Config{
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
	}))

	r.Gin.GET("/todo-items", func(c *gin.Context) {
		rows, err := r.DBConn.Query("SELECT id, title, content, date FROM todos")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		itemList := []Item{}
		for rows.Next() {
			var id int64
			var title string
			var content string
			var date string

			if err := rows.Scan(&id, &title, &content, &date); err != nil {
				log.Fatal(err)
			}
			itemList = append(itemList, Item{Id: id, Title: title, Content: content, Date: date})
		}
		c.JSON(http.StatusOK, itemList)
	})
	r.Gin.POST("/todo", func(c *gin.Context) {
		item := Item{}
		c.Bind(&item)

		newItem := Item{
			Title:   item.Title,
			Content: item.Content,
			Date:    time.Now().Format("2006-01-02"),
		}

		stmt, err := r.DBConn.Prepare("INSERT INTO todos (title, content, date) VALUES (?, ?, ?)")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "database error: prepare failed."})
		}
		defer stmt.Close()

		result, err := stmt.Exec(newItem.Title, newItem.Content, newItem.Date)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "database error: insert failed."})
		}

		id, err := result.LastInsertId()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "database error: get result last insert id."})
		}
		newItem.Id = id
		c.JSON(http.StatusOK, newItem)
	})
	r.Gin.DELETE("/todo", func(c *gin.Context) {
		item := Item{}
		c.Bind(&item)

		stmt, err := r.DBConn.Prepare("DELETE FROM todos WHERE id = ?")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "database error: prepare failed."})
		}

		_, err = stmt.Exec(item.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "database error: delete failed."})
		}

		c.JSON(http.StatusOK, Item{
			Id: item.Id,
		})
	})
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
