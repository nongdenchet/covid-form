package main

import (
	"log"
	"net/http"

	"github.com/nongdenchet/covidform/repository"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/nongdenchet/covidform/handler"
	"github.com/nongdenchet/covidform/model"
	"github.com/nongdenchet/covidform/utils"
)

func main() {
	// Database
	db, err := gorm.Open("mysql", "quan:developer@/covid_form?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migration
	err = db.AutoMigrate(&model.Venue{}).Error
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

	// Authen api
	authen := router.PathPrefix(utils.ApiV1).Subrouter()
	authen.Use(am.Handler)
	authen.HandleFunc("/welcome", h.WelcomeHandler).Methods("POST")

	// Start server
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
