package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/nongdenchet/covidform/repository"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/nongdenchet/covidform/handler"
	"github.com/nongdenchet/covidform/model"
	"github.com/nongdenchet/covidform/utils"
)

func main() {
	isDebug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		panic(err)
	}

	// Database
	db, err := gorm.Open("mysql", os.Getenv("DB_URI"))
	db.LogMode(isDebug)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migration
	err = db.AutoMigrate(&model.Venue{}).Error
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Visit{}).Error
	if err != nil {
		panic(err)
	}

	// Handlers
	h := handler.NewHandler(db)
	router := mux.NewRouter()
	am := handler.AuthenMiddleware{Repo: repository.NewVenueRepo(db)}

	// Non authen api
	nonAuthen := router.PathPrefix(utils.ApiV1).Subrouter()
	nonAuthen.HandleFunc("/venues", h.RegisterHandler).Methods("POST")
	nonAuthen.HandleFunc("/sessions", h.LoginHandler).Methods("POST")
	nonAuthen.HandleFunc("/venues/{id}", h.GetVenueHandler).Methods("GET")
	nonAuthen.HandleFunc("/venues/{id}/visits", h.SubmitFormHandler).Methods("POST")

	// Authen api
	authen := router.PathPrefix(utils.ApiV1).Subrouter()
	authen.Use(am.Handler)
	authen.HandleFunc("/venues/self", h.UpdateVenueHandler).Methods("PATCH")
	authen.HandleFunc("/venues/self/visits", h.GetVisitsByVenueHandler).Methods("GET")

	// Start server
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
