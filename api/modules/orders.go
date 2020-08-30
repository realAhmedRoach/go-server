package modules

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"sces/mgmt"
	"sces/store"
)

func OrderRoutes(app *mgmt.Application) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/sukuk", func(r chi.Router) {
		r.Get("/", listSukukOrders(app))
		r.Post("/", createSukukOrder(app))
		r.Get("/{uid}", getSukukOrder(app))
		r.Delete("/{uid}", deleteSukukOrder(app))
	})

	return router
}

func listSukukOrders(app *mgmt.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if res, err := app.SukukManager().(*store.DBSukukOrderService).List(); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err.Error()))
		} else {
			writer.Write([]byte(res))
		}
	}
}

func createSukukOrder(app *mgmt.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, _ := ioutil.ReadAll(request.Body)

		order := mgmt.SukukOrder{}

		if err := json.Unmarshal(body, &order); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(mgmt.JSONError{
				Msg: "invalid order input",
			}.Error()))
			return
		}

		order.FirmID = mgmt.TESTFIRMID

		// TODO: Validate order

		if uid, err := app.SukukManager().Put(order.FirmID, order.Sukuk, order.Price, order.Quantity, order.OrderSide, order.OrderType); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err.Error()))
		} else {
			writer.WriteHeader(http.StatusCreated)
			writer.Write([]byte(mgmt.JSONResult(uid)))
		}
	}
}

func getSukukOrder(app *mgmt.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		uid := chi.URLParam(request, "uid")

		if res, err := app.SukukManager().Get(uid); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err.Error()))
		} else {
			writer.Write([]byte(res))
		}
	}
}

func deleteSukukOrder(app *mgmt.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		uid := chi.URLParam(request, "uid")

		if err := app.SukukManager().Delete(uid); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err.Error()))
		} else {
			writer.WriteHeader(http.StatusOK)
		}
	}
}
