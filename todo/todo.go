package todo

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var index int
var tasks map[int]*Task = make(map[int]*Task)

type Task struct {
	Title string
	Done  bool
}

type NewTaskTodo struct {
	Task string `json:"task"`
}

type Serializer interface {
	Decode(io.ReadCloser, interface{}) error
	Encode(io.WriteCloser, interface{}) error
}

type JSONSerializer struct{}

func NewJSONSerializer() JSONSerializer {
	return JSONSerializer{}
}

func (JSONSerializer) Decode(body io.ReadCloser, v interface{}) error {
	return json.NewDecoder(body).Decode(v)
}

func (JSONSerializer) Encode(body io.WriteCloser, v interface{}) error {
	return json.NewEncoder(body).Encode(v)
}

type App struct {
	serialize Serializer
}

func NewApp(serializer Serializer) App {
	return App{serialize: serializer}
}

func (app *App) AddTask(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var task NewTaskTodo
	if err := app.serialize.Decode(r.Body, &task); err != nil {
		// if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	New(task.Task)
}

func SetDone(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	index := vars["index"]
	i, err := strconv.Atoi(index)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	tasks[i].Done = true

}

func GetTask(rw http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(rw).Encode(List()); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
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
