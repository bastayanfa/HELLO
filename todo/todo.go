package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ระวังเรื่อง graceful shutdown

type memDB map[int]*Task

// ห้ามใช้ตัวแปรที่อยู่ข้างนอกแบบนี้กับงาน API จะเกิด REST
// ถ้าต้องการใช้ต้องมี tool มาช่วย เช่น Sync new Text (?), pacge atomic ของ Go ก็ได้
var index int
var tasks memDB = make(memDB)

type Task struct {
	gorm.Model
	Title string
	Done  bool
}

type NewTasksTodo struct {
	Task string `json:"task"`
}

type Inserter interface {
	Insert(interface{}) error
}

type Insert *gorm.DB

func (gdb Insert) Insert(v interface{}) error {
	return gdb.Create(v).error
}

type nodb struct{}

func (db nodb) Insert(v interface{}) error {
	if cache, ok := v.(*Task); ok {
		tasks[index] = cache
	}
	return nil
}

type App struct {
	// db *gorm.DB
	db Inserter
}

func NewApp(db Inserter) *App {
	return &App{db: db}
}

func (app *App) AddTaskfunc(c *gin.Context) {
	var task NewTasksTodo
	if err := c.Bind(&task); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// New(task.Task)
	// app.db.Create(&Task{Title: task.Task, Done: false})
	app.db.Insert(&Task{Title: task.Task, Done: false})
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
