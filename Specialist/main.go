package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port string = "8090"
	db   []Pizza
)

type Pizza struct {
	ID       int     `json: "Id"`
	Diameter int     `json: "Diam"`
	Price    float64 `json: "Price"`
	Title    string  `json:"Title"`
}

func init() {
	pizza1 := Pizza{
		ID:       1,
		Diameter: 22,
		Price:    500.50,
		Title:    "Pepperoni",
	}

	pizza2 := Pizza{
		ID:       2,
		Diameter: 25,
		Price:    650.23,
		Title:    "BBQ",
	}
	pizza3 := Pizza{
		ID:       3,
		Diameter: 22,
		Price:    450,
		Title:    "Margaritta",
	}

	db = append(db, pizza1, pizza2, pizza3)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetGreet)
	router.HandleFunc("/{id}", GetByID)
	log.Fatal(http.ListenAndServe(":"+port, nil)) //запускает сервер и слушает указанный
}

//ResponseWriter место куда записываем ответ
//Request откуда брать запрос

//обработчик
func GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get by id")
}

func GetAllPizzas(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(db)
}

func GetPizzaByID(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
}
