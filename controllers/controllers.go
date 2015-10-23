package controllers

import (
	"encoding/json"
	"net/http"

	"labix.org/v2/mgo/bson"
)

func Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(bson.M{"status": "OK"})
}
