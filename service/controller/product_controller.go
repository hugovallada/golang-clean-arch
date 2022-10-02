package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hugovallada/go-clean-arch/client/model"
	"github.com/hugovallada/go-clean-arch/domain/usecase"
	"github.com/hugovallada/go-clean-arch/service/gateway"
	"github.com/hugovallada/go-clean-arch/service/translator"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var request model.ProductRequest
	w.Header().Add("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		errorMessage := model.ErrorMessage{Message: "Error ao decodificar a mensagem"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorMessage)
	}

	product := translator.FromProductRequestToProduct(request)
	saveGateway := gateway.SaveProductGatewayImpl{}

	code, err := usecase.SaveProductUseCase(product, &saveGateway)

	if err != nil {
		log.Println(err)
		errorMessage := model.ErrorMessage{Message: "Error ao persistir os dados"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorMessage)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("code", code)
	w.WriteHeader(http.StatusCreated)
}
