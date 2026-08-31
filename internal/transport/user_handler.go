package transport

import (
	"GOrm/internal/model"
	"GOrm/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler struct {
	service *service.UserService
}

func New(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (u *UserHandler) HandleAllUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users, err := u.service.GetAllUsers()
		if err != nil {
			http.Error(w, "No se encontraron usuarios", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Usuario inválido", http.StatusBadRequest)
		}

		created, err := u.service.CreateUser(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(created)

	default:
		http.Error(w, "No implementado", http.StatusNotImplemented)
		return
	}
}

func (u *UserHandler) HandleUserById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID no encontrado", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		user, err := u.service.GetUserById(id)
		if err != nil {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

	case http.MethodPut:
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Usuario inválido", http.StatusBadRequest)
			return
		}

		updated, err := u.service.UpdateUser(id, &user)
		if err != nil {
			http.Error(w, "Problema actualizando el usuario", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updated)

	case http.MethodDelete:
		if err := u.service.DeleteUser(id); err != nil {
			http.Error(w, "No se pudo borrar al usuario", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

	default:
		http.Error(w, "No implementado", http.StatusNotImplemented)
	}

}
