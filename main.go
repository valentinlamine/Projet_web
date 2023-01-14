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
	//gestion js
	script := http.FileServer(http.Dir("script"))
	http.Handle("/script/", http.StripPrefix("/script/", script))
	//gestion html
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/Achat", AchatHandler)
	http.HandleFunc("/QuiSommesNous", QuiSommesNousHandler)
	http.HandleFunc("/MonCompte", MonCompteHandler)
	http.HandleFunc("/Inscription", InscriptionHandler)
	http.HandleFunc("/Resultats", ResultatHandler)
	http.HandleFunc("/Erreur", ErreurHandler)
	http.ListenAndServe(":80", nil)
}
func alreadyregistered(email string) bool {
	jsonData, err := ioutil.ReadFile("json/users.json")
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

	jsonData, err := ioutil.ReadFile("json/users.json")
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
	err = ioutil.WriteFile("json/users.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("could not write json: %s\n", err)
		return
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("templates/index.html"))
	if r.Method != http.MethodPost {
		tmpl1.Execute(w, nil)
		return
	}
	tmpl1.Execute(w, nil)
}

func AchatHandler(w http.ResponseWriter, r *http.Request) {
	tmpl2 := template.Must(template.ParseFiles("templates/Achat.html"))
	tmpl2.Execute(w, cars)
}

func QuiSommesNousHandler(w http.ResponseWriter, r *http.Request) {
	tmpl3 := template.Must(template.ParseFiles("templates/quisommesnous.html"))
	tmpl3.Execute(w, nil)
}

func MonCompteHandler(w http.ResponseWriter, r *http.Request) {
	tmpl4 := template.Must(template.ParseFiles("templates/MonCompte.html"))
	tmpl4.Execute(w, nil)
}

func InscriptionHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("email") != "" {
		if alreadyregistered(r.FormValue("email")) {
			http.Redirect(w, r, "/Erreur", http.StatusSeeOther)
		}
		register(r.FormValue("email"), r.FormValue("password"))
	}
	tmpl5 := template.Must(template.ParseFiles("templates/Inscription.html"))
	tmpl5.Execute(w, nil)
}

func ResultatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("search"))
	tmpl6 := template.Must(template.ParseFiles("templates/Resultats.html"))
	tmpl6.Execute(w, nil)
}

func ErreurHandler(w http.ResponseWriter, r *http.Request) {
	tmpl7 := template.Must(template.ParseFiles("templates/Erreur.html"))
	tmpl7.Execute(w, nil)
}

type User struct {
	Email    string
	Password string
}

type Car struct {
	Image      string
	Brand      string
	Engine     string
	Kilometers int
	MaxSpeed   int
	Year       int
	Price      int
}

// declaration des voitures
var cars = []Car{
	{
		Image:      "/images/car_1.png",
		Brand:      "Mercedes-Benz S-Class",
		Engine:     "300 hybride",
		Kilometers: 100000,
		MaxSpeed:   250,
		Year:       2022,
		Price:      150000,
	},
	{
		Image:      "/images/car_1.png",
		Brand:      "Audi2",
		Engine:     "Diesel",
		Kilometers: 100000,
		MaxSpeed:   200,
		Year:       2010,
		Price:      50,
	},
	{
		Image:      "/images/car_1.png",
		Brand:      "Audi3",
		Engine:     "Diesel",
		Kilometers: 100000,
		MaxSpeed:   200,
		Year:       2010,
		Price:      5000,
	},
	{
		Image:      "/images/car_1.png",
		Brand:      "Audi4",
		Engine:     "Diesel",
		Kilometers: 100000,
		MaxSpeed:   200,
		Year:       2010,
		Price:      5000,
	},
	{
		Image:      "/images/car_1.png",
		Brand:      "Audi5",
		Engine:     "Diesel",
		Kilometers: 100000,
		MaxSpeed:   200,
		Year:       2010,
		Price:      5000,
	},
	{
		Image:      "/images/car_1.png",
		Brand:      "Audi6",
		Engine:     "Diesel",
		Kilometers: 100000,
		MaxSpeed:   200,
		Year:       2010,
		Price:      5000,
	},
	{
		Image:      "/images/car_1.png",
		Brand:      "Audi7",
		Engine:     "Diesel",
		Kilometers: 100000,
		MaxSpeed:   200,
		Year:       2010,
		Price:      5000,
	},
	{
		Image:      "/images/car_1.png",
		Brand:      "Audi8",
		Engine:     "Diesel",
		Kilometers: 100000,
		MaxSpeed:   200,
		Year:       2010,
		Price:      5000,
	},
	{
		Image:      "/images/car_1.png",
		Brand:      "Audi9",
		Engine:     "Diesel",
		Kilometers: 100000,
		MaxSpeed:   200,
		Year:       2010,
		Price:      5000,
	},
	{
		Image:      "/images/car_1.png",
		Brand:      "Audi10",
		Engine:     "Diesel",
		Kilometers: 100000,
		MaxSpeed:   200,
		Year:       2010,
		Price:      5000,
	},
	{
		Image:      "/images/car_1.png",
		Brand:      "Audi11",
		Engine:     "Diesel",
		Kilometers: 100000,
		MaxSpeed:   200,
		Year:       2010,
		Price:      5000,
	},
	{
		Image:      "/images/car_1.png",
		Brand:      "Audi12",
		Engine:     "Diesel",
		Kilometers: 100000,
		MaxSpeed:   200,
		Year:       2010,
		Price:      5000,
	},
	{
		Image:      "/images/car_1.png",
		Brand:      "Audi13",
		Engine:     "Diesel",
		Kilometers: 100000,
		MaxSpeed:   200,
		Year:       2010,
		Price:      5000,
	},
	{
		Image:      "/images/car_1.png",
		Brand:      "Audi14",
		Engine:     "Diesel",
		Kilometers: 100000,
		MaxSpeed:   200,
		Year:       2010,
		Price:      5000,
	},
	{
		Image:      "/images/car_1.png",
		Brand:      "Audi15",
		Engine:     "Diesel",
		Kilometers: 100000,
		MaxSpeed:   200,
		Year:       2010,
		Price:      5000,
	},
}
