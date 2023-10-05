package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/leandro-portugal/MyTask.git/models"
)

func List(w http.ResponseWriter, r *http.Request) {

	tasks, err := models.Read_All()

	if err != nil {
		log.Printf("Erro ao obter registros %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
