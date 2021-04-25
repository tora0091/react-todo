package controller

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/tora0091/react-todo/server/domain"
)

type TodoController struct {
	DBConn *sql.DB
}

func NewTodoController(dbConn *sql.DB) *TodoController {
	return &TodoController{
		DBConn: dbConn,
	}
}

func (t *TodoController) GetItems(c *gin.Context) {
	rows, err := t.DBConn.Query("SELECT id, title, content, date FROM todos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	itemList := domain.Items{}
	for rows.Next() {
		var id int64
		var title string
		var content string
		var date string

		if err := rows.Scan(&id, &title, &content, &date); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		itemList = append(itemList, domain.Item{Id: id, Title: title, Content: content, Date: date})
	}
	c.JSON(http.StatusOK, itemList)
}

func (t *TodoController) RegisterItem(c *gin.Context) {
	item := domain.Item{}
	c.Bind(&item)

	newItem := domain.Item{
		Title:   item.Title,
		Content: item.Content,
		Date:    time.Now().Format("2006-01-02"),
	}

	stmt, err := t.DBConn.Prepare("INSERT INTO todos (title, content, date) VALUES (?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error: prepare failed."})
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(newItem.Title, newItem.Content, newItem.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error: insert failed."})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error: get result last insert id."})
		return
	}
	newItem.Id = id
	c.JSON(http.StatusOK, newItem)
}

func (t *TodoController) DeleteItem(c *gin.Context) {
	item := domain.Item{}
	c.Bind(&item)

	stmt, err := t.DBConn.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error: prepare failed."})
		return
	}

	_, err = stmt.Exec(item.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "database error: delete failed."})
		return
	}

	c.JSON(http.StatusOK, domain.Item{
		Id: item.Id,
	})
}
