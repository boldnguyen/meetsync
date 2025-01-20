package models

type Room struct {
	Name string `json:"name"`
}

type JoinRequest struct {
	RoomID   string `json:"roomID"`
	UserName string `json:"userName"`
}
