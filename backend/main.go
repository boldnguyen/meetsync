package main

import (
	"log"
	"meetsync/backend/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	// Đăng ký các route
	r.HandleFunc("/api/rooms", handlers.CreateRoomHandler).Methods("POST")
	r.HandleFunc("/api/rooms/{roomID}/join", handlers.JoinRoomHandler).Methods("POST")

	// Cấu hình CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081"}, // Cho phép frontend trên cổng 8081
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Khởi động server
	handler := c.Handler(r)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
