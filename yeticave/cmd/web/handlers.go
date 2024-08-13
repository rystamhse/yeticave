package main

import (
	"html/template"
	"log"
	"net/http"
)

// Обработчик главной странице.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/index" {
		//http.NotFound(w, r)
		ts, err := template.ParseFiles("./ui/html/404.html")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
		return
	}

	files := []string{
		"./ui/html/index.html",
		"./ui/html/partial.tmpl",
		"./ui/html/footer.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// Обработчик страницы добавления лота.
func AddLot(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/add-lot.html",
		"./ui/html/partial.tmpl",
		"./ui/html/footer.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// Обработчик страницы лота.
func AllLots(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/all-lots.html",
		"./ui/html/partial.tmpl",
		"./ui/html/footer.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// Обработчик страницы регистрации.
func LoginUser(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/login.html",
		"./ui/html/partial.tmpl",
		"./ui/html/footer.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// Обработчик страницы Лота.
func Lot(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/lot.html",
		"./ui/html/partial.tmpl",
		"./ui/html/footer.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
