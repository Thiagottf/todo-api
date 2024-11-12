// handlers/task_handlers.go
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/seu-usuario/todo-api/database"
	"github.com/seu-usuario/todo-api/models"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.Tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	task.ID = len(database.Tasks) + 1
	database.Tasks = append(database.Tasks, task)
	json.NewEncoder(w).Encode(task)
}

// Para UpdateTask e DeleteTask, você pode continuar criando handlers similares.
