package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hugovallada/go-clean-arch/client/model"
	errormodel "github.com/hugovallada/go-clean-arch/domain/error_model"
	"github.com/hugovallada/go-clean-arch/domain/usecase"
	"github.com/hugovallada/go-clean-arch/service/gateway"
	"github.com/hugovallada/go-clean-arch/service/translator"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var request model.ProductRequest
	w.Header().Add("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		HandlerError(err, 500, "Error ao decodificar uma mensagem", w)
		return
	}
	product := translator.FromProductRequestToProduct(request)
	saveGateway := gateway.SaveProductGatewayImpl{}
	getProductGateway := gateway.GetProductsGatewayImpl{}
	code, err := usecase.SaveProductUseCase(product, &saveGateway, &getProductGateway)
	if err != nil {
		if ue, ok := err.(*errormodel.UnprocessableEntityError); ok {
			HandlerError(ue, 422, ue.Message, w)
		} else {
			HandlerError(err, 500, "Error ao persistir o dado", w)
		}
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("code", code)
	w.WriteHeader(http.StatusCreated)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	getAllGateway := gateway.GetProductsGatewayImpl{}
	products, err := usecase.GetAllProducts(&getAllGateway)
	if err != nil {
		HandlerError(err, http.StatusInternalServerError, "Algo aconteceu", w)
		return
	}
	w.WriteHeader(http.StatusOK)
	productsResponse := translator.FromListOfProductsToListOfProductsReponse(products)
	json.NewEncoder(w).Encode(productsResponse)
}

func HandlerError(err error, statusCode int, errorMessage string, w http.ResponseWriter) {
	log.Println(err)
	errorResponse := model.ErrorMessage{Message: errorMessage}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResponse)
}
