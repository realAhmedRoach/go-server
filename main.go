package main

import (
	"github.com/joho/godotenv"
	"net/http"
	"sces/api"
)

func main() {
	if err := godotenv.Load(".env.dev"); err != nil {
		panic(err.Error())
	}

	conn := api.Connect()
	var app = api.Application{
		Sukuk: &api.DBSukukOrderService{
			Conn: conn,
		},
	}

	defer conn.Close()

	router := api.Routes(&app)

	if err := http.ListenAndServe(":8090", router); err != nil {
		panic(err.Error())
	}
}
