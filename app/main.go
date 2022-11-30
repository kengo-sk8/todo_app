package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

)

type Todo struct {
	ID          int    `from:"ida"`
	User       	string `form:"user"`
	Task        string `form:"task"`
	Status 		int	   `from:"status"`
	CreateDate 	string
}

var todo []Todo
var idMax = 0

func AutoId() int {
	return idMax + 1
}

func GetDataTodo(c *gin.Context) {
	var b Todo
	if err := c.Bind(&b); err != nil {
		fmt.Errorf("%#v", err)
	}
	b.ID = AutoId()
	b.Status = 0
	b.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	todo = append(todo, b)
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todo": todo,
	})
}

func GetDoneTodo(c *gin.Context) {
	var b Todo
	if err := c.Bind(&b); err != nil {
		fmt.Errorf("%#v", err)
	}
	var s int

	if b.Status == 0 {
		s = 1
	} else {
		s = 0
	}

	for idx, t := range todo {
		if t.ID == b.ID {
			todo[idx].Status = s
		}
	}
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todo": todo,
	})
}


func main() {

	todo := []Todo{
		Todo {
			User:       "test1",
			Task:       "掃除",
			Status:		0,
			CreateDate:	time.Now().Format("2006-01-02 15:04:05"),
		},
		Todo {
			User:       "test2",
			Task:       "あいう",
			Status:		1,
			CreateDate:	time.Now().Format("2006-01-02 15:04:05"),
		},
		Todo {
			User:       "test3",
			Task:       "ジョンソン",
			Status:		1,
			CreateDate:	time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	router := gin.Default()
	// viewファイルの呼び出し
    router.LoadHTMLFiles("./views/index.html")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/todo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{} {
			"todo": todo,
		})
	})
	router.GET("/conduct", GetDataTodo)
	router.GET("/done", GetDoneTodo)
    router.Run(":8080")
}
