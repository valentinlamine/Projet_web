package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func main() {
	//démarage serveur
	fmt.Println("Server is running on port 80 http://localhost")
	//gestion css
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	//gestion img
	images := http.FileServer(http.Dir("images"))
	http.Handle("/images/", http.StripPrefix("/images/", images))
	//gestion html
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/Achat", buyHandler)
	http.HandleFunc("/QuiSommesNous", whoareusHandler)
	http.HandleFunc("/MonCompte", accountHandler)
	http.HandleFunc("/Inscription", registerHandler)
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

func registerHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Email":    r.FormValue("email"),
		"Password": r.FormValue("password"),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	fmt.Printf("json data: %s\n", jsonData)
	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("info.json", file, 0644)

	tmpl5 := template.Must(template.ParseFiles("templates/Inscription.html"))
	tmpl5.Execute(w, nil)
}
