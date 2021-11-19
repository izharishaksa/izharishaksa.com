package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func main() {
	http.Handle("/asset/",
		http.StripPrefix("/asset/",
			http.FileServer(http.Dir("./internal/asset"))))
	http.Handle("/image/",
		http.StripPrefix("/image/",
			http.FileServer(http.Dir("./internal/image"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("./internal/view", "index.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "Learning Golang Web",
			"name":  "Batman",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Start server at localhost:900")
	err := http.ListenAndServe(":900", nil)
	if err != nil {
		panic(err)
	}
}
