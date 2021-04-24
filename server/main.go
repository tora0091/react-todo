package main

import (
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

var ItemList []Item

func init() {
	ItemList = []Item{
		Item{Id: 1, Title: "Good hey hey hey", Content: "Hello i am very happy!", Date: "2020-10-04"},
		Item{Id: 2, Title: "Oh my cat dog", Content: "i love cat, and dog", Date: "2020-12-14"},
		Item{Id: 3, Title: "see you later", Content: "see you next monday. OK, you are good", Date: "2021-02-22"},
		Item{Id: 4, Title: "Mornings", Content: "i am cow. i am very hungry", Date: "2021-04-10"},
	}
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
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

	r.GET("/todo-items", func(c *gin.Context) {
		c.JSON(http.StatusOK, ItemList)
	})
	r.POST("/todo", func(c *gin.Context) {
		item := Item{}
		c.Bind(&item)

		now := time.Now()
		unixTime := now.Unix()

		newItem := Item{
			Id:      unixTime,
			Title:   item.Title,
			Content: item.Content,
			Date:    now.Format("2006-01-02"),
		}

		ItemList = append(ItemList, newItem)
		c.JSON(http.StatusOK, newItem)
	})
	r.DELETE("/todo", func(c *gin.Context) {
		item := Item{}
		c.Bind(&item)

		newItems := []Item{}
		for _, itm := range ItemList {
			if itm.Id != item.Id {
				newItems = append(newItems, itm)
			}
		}

		ItemList = newItems
		c.JSON(http.StatusOK, Item{
			Id: item.Id,
		})
	})
	r.Run(":9090")
}
