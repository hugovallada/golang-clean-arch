package client

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hugovallada/go-clean-arch/service/controller"
)

func Routes() http.Handler {
	router := chi.NewRouter()

	router.Post("/products", controller.CreateProduct)
	router.Get("/products", controller.GetAll)

	return router
}
