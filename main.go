package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

var TESTFIRMID = "baf78936-5986-4f24-8a40-5e11aef970c6"

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
