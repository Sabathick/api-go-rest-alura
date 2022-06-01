package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sabathick/api-go-rest-alura/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
