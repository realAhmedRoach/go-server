package modules

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"sces/api"
)

func OrderRoutes(app *api.Application) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/sukuk", func(r chi.Router) {
		router.Post("/", createSukukOrder(app))
		router.Get("/{uid}", getSukukOrder(app))
		router.Delete("/{uid}", deleteSukukOrder(app))
	})

	return router
}

func createSukukOrder(app *api.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, _ := ioutil.ReadAll(request.Body)
		defer request.Body.Close()

		order := api.SukukOrder{}

		if err := json.Unmarshal(body, &order); err != nil {
			writer.Write([]byte(api.JSONError{Msg: err.Error()}.Error()))
			return
		}

		order.FirmID = api.TESTFIRMID

		// TODO: Validate order

		if uid, err := app.Sukuk.Put(order.FirmID, order.Sukuk, order.Price, order.Quantity, order.Side, order.OrderType); err != nil {
			writer.Write([]byte(err.Error()))
		} else {
			writer.Write([]byte(api.JSONResult(uid)))
		}
	}
}

func getSukukOrder(app *api.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		uid := chi.URLParam(request, "uid")

		if res, err := app.Sukuk.Get(uid); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
		} else {
			writer.Write([]byte(res))
		}
	}
}

func deleteSukukOrder(app *api.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		uid := chi.URLParam(request, "uid")

		if err := app.Sukuk.Delete(uid); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		} else {
			writer.WriteHeader(http.StatusOK)
		}
	}
}
