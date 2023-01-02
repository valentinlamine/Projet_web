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
	http.HandleFunc("/Resultats", resultHandler)
	http.HandleFunc("/Erreur", errorHandler)
	http.ListenAndServe(":80", nil)
}
func alreadyregistered(email string) bool {
	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return false
	}
	//decode les données du fichier json dans une map
	var users []User
	err = json.Unmarshal(jsonData, &users)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return false
	}
	for _, user := range users {
		if user.Email == email {
			return true
		}
	}
	return false
}
func register(email string, password string) {
	data := map[string]interface{}{
		"Email":    email,
		"Password": password,
	}

	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	//decode les données du fichier json dans une map
	var users []User
	err = json.Unmarshal(jsonData, &users)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}
	//ajouter les données du formulaire dans la map
	users = append(users, User{
		Email:    data["Email"].(string),
		Password: data["Password"].(string),
	})
	//encoder les données de la map dans un fichier json
	jsonData, err = json.Marshal(users)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	//écrire les données dans le fichier json
	err = ioutil.WriteFile("users.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("could not write json: %s\n", err)
		return
	}
}

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
	if r.FormValue("email") != "" {
		if alreadyregistered(r.FormValue("email")) {
			http.Redirect(w, r, "/Erreur", http.StatusSeeOther)
		}
		register(r.FormValue("email"), r.FormValue("password"))
	}
	tmpl5 := template.Must(template.ParseFiles("templates/Inscription.html"))
	tmpl5.Execute(w, nil)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("search"))
	tmpl6 := template.Must(template.ParseFiles("templates/Resultats.html"))
	tmpl6.Execute(w, nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl7 := template.Must(template.ParseFiles("templates/Erreur.html"))
	tmpl7.Execute(w, nil)
}

type User struct {
	Email    string
	Password string
}
