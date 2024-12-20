package common

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func Json(w http.ResponseWriter, code int, msg string, payload interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if payload == false {
		json.NewEncoder(w).Encode(bson.M{"code": code, "success": false, "msg": msg, "data": nil})
	}

	json.NewEncoder(w).Encode(bson.M{"code": code, "success": true, "msg": msg, "data": payload})
}
