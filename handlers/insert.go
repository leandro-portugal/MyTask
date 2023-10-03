package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/leandro-portugal/MyTask.git/models"
)

func Insert(w http.ResponseWriter, r *http.Request) {

	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {

		log.Printf("Erro ao fazer decode do json %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Create(task)

	var response map[string]any

	if err != nil {
		response = map[string]any{
			"error":   true,
			"message": fmt.Sprintf("Ocorreu um erro na criação de uma nova tarefa %v", err),
		}
	} else {
		response = map[string]any{
			"error":   false,
			"message": fmt.Sprintf("Sucesso na criação de uma nova tarefa, id: %d", id),
		}
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
