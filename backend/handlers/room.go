package handlers

import (
	"encoding/json"
	"log"
	"meetsync/backend/models"
	"meetsync/backend/services"
	"net/http"
)

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateRoomHandler: Request received")

	var room models.Room
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		log.Println("CreateRoomHandler: Invalid request body", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Println("CreateRoomHandler: Creating room with name:", room.Name)

	roomID, err := services.CreateCloudflareRoom(room.Name)
	if err != nil {
		log.Println("CreateRoomHandler: Failed to create room", err)
		http.Error(w, "Failed to create room", http.StatusInternalServerError)
		return
	}

	log.Println("CreateRoomHandler: Room created with ID:", roomID)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"roomID": roomID})
}

func JoinRoomHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("JoinRoomHandler: Request received")

	var joinRequest models.JoinRequest
	err := json.NewDecoder(r.Body).Decode(&joinRequest)
	if err != nil {
		log.Println("JoinRoomHandler: Invalid request body", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Println("JoinRoomHandler: Joining room with ID:", joinRequest.RoomID, "and user:", joinRequest.UserName)

	token, err := services.JoinCloudflareRoom(joinRequest.RoomID, joinRequest.UserName)
	if err != nil {
		log.Println("JoinRoomHandler: Failed to join room", err)
		http.Error(w, "Failed to join room", http.StatusInternalServerError)
		return
	}

	log.Println("JoinRoomHandler: User joined with token:", token)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
