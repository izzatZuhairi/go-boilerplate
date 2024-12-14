package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Handler struct {
	db *mongo.Client
}

func NewHandler(db *mongo.Client) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	userCollection := h.db.Database("user").Collection("users")

	userCursor, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch", http.StatusInternalServerError)
		return
	}

	defer userCursor.Close(context.Background())

	var users []bson.M
	if err := userCursor.All(context.Background(), &users); err != nil {
		http.Error(w, "Failed to decode users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
