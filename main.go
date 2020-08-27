package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func main() {
	conn := Connect()
	defer conn.Close()

	router := httprouter.New()

	router.GET("/firm", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		writer.Header().Add("Content-Type", "application/json")
		fmt.Fprint(writer, List(DB_FIRM, conn))
	})

	router.GET("/firm/:uid", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		writer.Header().Add("Content-Type", "application/json")
		fmt.Fprint(writer, Retrieve(params.ByName("uid"), DB_FIRM, conn))
	})

	router.POST("/sukuk", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		body, _ := ioutil.ReadAll(request.Body)
		println(string(body))
	})

	if err := http.ListenAndServe(":8090", router); err != nil {
		return
	}
}
