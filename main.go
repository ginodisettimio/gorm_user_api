package main

import (
	"GOrm/internal/model"
	"GOrm/internal/service"
	"GOrm/internal/store"
	"GOrm/internal/transport"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// Conexión Base de Datos
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.User{})

	// Inyección de dependencias
	UserStore := store.New(db)
	UserService := service.New(UserStore)
	UserHandler := transport.New(UserService)

	// Rutas
	http.HandleFunc("/users", UserHandler.HandleAllUsers)
	http.HandleFunc("/users/", UserHandler.HandleUserById)

	// Servidor
	fmt.Println("Endpoints:")
	fmt.Println("GET /users 		-> Retorna todos los usuarios")
	fmt.Println("POST /users 		-> Crea un usuario nuevo")
	fmt.Println("GET /users/{id}	-> Retorna usuario por ID")
	fmt.Println("PUT /users/{id} 	-> Actualiza al usuario si es válido")
	fmt.Println("DELETE /users/{id} -> Borra al usuario si existe")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
