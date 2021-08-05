package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ระวังเรื่อง graceful shutdown

// ห้ามใช้ตัวแปรที่อยู่ข้างนอกแบบนี้กับงาน API จะเกิด REST
// ถ้าต้องการใช้ต้องมี tool มาช่วย เช่น Sync new Text (?), pacge atomic ของ Go ก็ได้
var index int
var tasks map[int]*Task = make(map[int]*Task)

type Task struct {
	Title string
	Done  bool
}

type NewTasksTodo struct {
	Task string `json:"task"`
}

func AddTaskfunc(c *gin.Context) {
	var task NewTasksTodo
	if err := c.Bind(&task); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	New(task.Task)
}

func SetDonefunc(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	tasks[i].Done = true
}

func GetTaskfunc(c *gin.Context) {
	c.JSON(http.StatusOK, List())
}

func List() map[int]*Task {
	return tasks
}

func New(task string) {
	defer func() {
		index++
	}()

	tasks[index] = &Task{
		Title: task,
		Done:  false,
	}
}
