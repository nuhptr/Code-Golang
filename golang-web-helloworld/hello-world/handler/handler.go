package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, stay calm", http.StatusInternalServerError)
		return
	}

	// menggunakan map
	// data := map[string]interface{}{
	// 	"title":   "Golang - web",
	// 	"content": "Still learning with mas Agung Setiawan",
	// }

	// menggunakan struct
	// data := entity.Product{
	// 	ID:    1,
	// 	Name:  "Adi Nugraha Putra",
	// 	Price: 2200000,
	// 	Stock: 3,
	// }

	// menggunakan slice of struct
	data := []entity.Product{
		{ID: 1, Name: "Xpander", Price: 20000, Stock: 11},
		{ID: 2, Name: "Pajero", Price: 30000, Stock: 1},
		{ID: 3, Name: "Pajero Sport", Price: 40000, Stock: 8},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is coming, keep calm", http.StatusInternalServerError)
		return
	}
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario From Nintendo"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, stay calm", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"content": idNumb,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is coming, keep calm", http.StatusInternalServerError)
		return
	}
}
