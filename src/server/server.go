package main

import (
	"net/http"
	"fmt"
	"html/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}
func biddingHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/bidding.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "bidding", nil)
}

func driversHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/drivers.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "drivers", nil)
}

func carsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/cars.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "cars", nil)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/bidding", biddingHandler)
	http.HandleFunc("/drivers", driversHandler)
	http.HandleFunc("/cars", carsHandler)
	http.ListenAndServe(":3000", nil)
}