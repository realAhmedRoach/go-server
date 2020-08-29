package main

import (
	"github.com/joho/godotenv"
	"net/http"
	"sces/api"
	"sces/mgmt"
	"sces/store"
)

func main() {
	if err := godotenv.Load(".env.dev"); err != nil {
		panic(err.Error())
	}

	conn := store.Connect()
	var app = mgmt.Application{Sukuk: &store.DBSukukOrderService{Conn: conn}}

	defer conn.Close()

	router := api.Routes(&app)

	if err := http.ListenAndServe(":8090", router); err != nil {
		panic(err.Error())
	}
}
