package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/leandro-portugal/MyTask.git/models"
)

func Update(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {

		log.Printf("Erro ao fazer parse do id %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var task models.Task

	err = json.NewDecoder(r.Body).Decode(&task)

	if err != nil {

		log.Printf("Erro ao fazer decode do json %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(int64(id), task)

	if err != nil {

		log.Printf("Erro ao atualizar registro %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return

	}
	if rows > 1 {
		log.Printf("Error: Foram atualizados %d registros", rows)
	}

	response := map[string]any{
		"error":   false,
		"message": "Dados atualizados com sucesso!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
