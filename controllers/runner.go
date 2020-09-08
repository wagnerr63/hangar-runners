package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"../models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	runners := models.GetAllRunners()

	temp.ExecuteTemplate(w, "Index", runners)

}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		funcao := r.FormValue("funcao")
		aniversario := r.FormValue("aniversario")
		hangarRun := r.FormValue("hangarRun")

		aniversarioConvertido, err := time.Parse("2006-01-02", aniversario)
		if err != nil {
			log.Println("Erro ao converter o aniversário: ", err)
		}

		models.AddRunner(nome, funcao, aniversarioConvertido, hangarRun)

	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	idRunner := r.URL.Query().Get("id")
	models.DeleteRunner(idRunner)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idRunner := r.URL.Query().Get("id")
	runner := models.GetRunner(idRunner)

	temp.ExecuteTemplate(w, "Edit", runner)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		funcao := r.FormValue("funcao")
		aniversario := r.FormValue("aniversario")
		hangarRun := r.FormValue("hangarRun")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error na conversão do ID: ", err)
		}

		aniversarioConvertido, err := time.Parse("2006-01-02", aniversario)
		if err != nil {
			log.Println("Erro ao converter o aniversário: ", err)
		}

		models.UpdateRunner(idConv, nome, funcao, aniversarioConvertido, hangarRun)

	}
	http.Redirect(w, r, "/", 301)
}

func SendManifest(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("manifest.json")
	if err != nil {
		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}

func SendSW(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("service-worker.js")
	if err != nil {
		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Write(data)
}
