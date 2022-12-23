package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	//démarage serveur
	fmt.Println("Server is running on port 80 http://localhost")
	//gestion css
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	//gestion img
	img := http.FileServer(http.Dir("img"))
	http.Handle("/img/", http.StripPrefix("/img/", img))
	//gestion html
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/Achat", buyHandler)
	http.HandleFunc("/QuiSommesNous", whoareusHandler)
	http.HandleFunc("/MonCompte", accountHandler)
	http.ListenAndServe(":80", nil)
}

// func OnlyLetter(s string) bool {
// 	for _, r := range s {
// 		if !unicode.IsLetter(r) {
// 			return false
// 		}
// 	}
// 	return true
// }

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("templates/index.html"))
	if r.Method != http.MethodPost {
		tmpl1.Execute(w, nil)
		return
	}
	tmpl1.Execute(w, nil)
}

func buyHandler(w http.ResponseWriter, r *http.Request) {
	tmpl2 := template.Must(template.ParseFiles("templates/Achat.html"))
	tmpl2.Execute(w, nil)
}

func whoareusHandler(w http.ResponseWriter, r *http.Request) {
	tmpl3 := template.Must(template.ParseFiles("templates/quisommesnous.html"))
	tmpl3.Execute(w, nil)
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	tmpl4 := template.Must(template.ParseFiles("templates/MonCompte.html"))
	tmpl4.Execute(w, nil)
}
