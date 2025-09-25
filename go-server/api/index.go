package handler

import (
	"encoding/json"
	"net/http"
)

type Map map[string]any

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Map{
		"success": true,
		"message": "Hello World!",
	})
}
