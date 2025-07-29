package controllers

import (
	"encoding/json"
	"go-api/config"
	"go-api/models"
	"go-api/repositories"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPessoas(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	pessoas, err := repositories.GetAll(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pessoas)
}

func GetPessoa(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	pessoa, err := repositories.GetByID(db, id)
	if err != nil {
		http.Error(w, "Pessoa não encontrada", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(pessoa)
}

func CreatePessoa(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	var p models.Pessoa
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := repositories.Create(db, &p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func UpdatePessoa(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var p models.Pessoa
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p.ID = id
	if err := repositories.Update(db, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func DeletePessoa(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := repositories.Delete(db, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
