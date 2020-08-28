package modules

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"sces/api"
)

func OrderRoutes(app *api.Application) *chi.Mux {
	router := chi.NewRouter()

	router.Post("/sukuk", createOrder(app))

	return router
}

func createOrder(app *api.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, _ := ioutil.ReadAll(request.Body)
		defer request.Body.Close()

		order := api.SukukOrder{}

		if err := json.Unmarshal(body, &order); err != nil {
			_, _ = writer.Write([]byte(api.JSONError(err.Error())))
		}

		order.FirmID = api.TESTFIRMID

		// TODO: Validate order

		uid := app.Sukuk.Put(order.FirmID, order.Sukuk, order.Price, order.Quantity, order.Side, order.OrderType)

		fmt.Fprint(writer, uid)
	}
}
