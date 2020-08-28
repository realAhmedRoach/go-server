package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var TESTFIRMID = "baf78936-5986-4f24-8a40-5e11aef970c6"

func main() {
	conn := Connect()
	defer conn.Close()

	router := chi.NewRouter()

	router.Use(middleware.SetHeader("Content-Type", "application/json"))

	router.Get("/firm", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		fmt.Fprint(writer, List(DB_FIRM, conn))
	})

	router.Get("/firm/{uid}", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		fmt.Fprint(writer, Retrieve(chi.URLParam(request, "uid"), DB_FIRM, conn))
	})

	router.Post("/sukuk", func(writer http.ResponseWriter, request *http.Request) {
		body, _ := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		println(string(body))

		order := SukukOrder{}

		if err := json.Unmarshal(body, &order); err != nil {
			_, _ = writer.Write([]byte(JSONError(err.Error())))
		}

		order.FirmID = TESTFIRMID

		uid := Insert(SUKUKORDERSCHEMA, conn,
			order.FirmID, order.Sukuk, order.Price, order.Quantity, order.Side, order.OrderType,
		)

		fmt.Fprint(writer, uid)
	})

	if err := http.ListenAndServe(":8090", router); err != nil {
		return
	}
}
