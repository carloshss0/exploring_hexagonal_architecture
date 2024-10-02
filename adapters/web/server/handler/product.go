package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/carloshss0/exploring_hexagonal_architecture/adapters/dto"
	"github.com/carloshss0/exploring_hexagonal_architecture/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type UpdatePriceRequest struct {
	Price float64 `json:{price}`
}

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/product/{id}/enable", n.With(
		negroni.Wrap(changeProductStatus(service)),
	)).Methods("PUT", "OPTIONS")

	r.Handle("/product/{id}/disable", n.With(
		negroni.Wrap(changeProductStatus(service)),
	)).Methods("PUT", "OPTIONS")

	r.Handle("/product/{id}/price", n.With(
		negroni.Wrap(updateProductPrice(service)),
	)).Methods("PUT", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return 
		}

		err = json.NewEncoder(w).Encode(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return 
		}


	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var productDto dto.Product

		err := json.NewDecoder(r.Body).Decode(&productDto)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		product, err := service.Create(productDto.Name, productDto.Price)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

	})
}

func changeProductStatus(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return 
		}
		url := r.URL.String()
		var result interface{}

		if strings.Contains(url, "enable") {
			result, err = service.Enable(product)
		} else if strings.Contains(url, "disable") {
			result, err = service.Disable(product)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError("Unsupported Action!"))
			return
		}

		// result, err := service.Enable(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(result)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return 
		}


	})
}

func updateProductPrice(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return 
		}
		
		var req UpdatePriceRequest

		err = json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		result, err := service.UpdatePrice(product, req.Price)
		// result, err := service.Enable(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(result)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return 
		}


	})
}

