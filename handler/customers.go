package handler

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gitlab.com/idoko/bucketeer/db"
	"gitlab.com/idoko/bucketeer/models"
	"net/http"
	"strconv"
)

var customerIDKey = "customerID"

func items(router chi.Router) {
	router.Get("/", getAllCustomers)
	router.Post("/", createCustomer)
	router.Route("/{customerId}", func(router chi.Router) {
		router.Use(ItemContext)
		router.Get("/", getCustomer)
		router.Put("/", updateCustomer)
		router.Delete("/", deleteCustomer)
	})
}
func CustomersContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customerId := chi.URLParam(r, "customerId")
		if customerId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("customer ID is required")))
			return
		}
		id, err := strconv.Atoi(customerId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid item ID")))
		}
		ctx := context.WithValue(r.Context(), customerIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func createCustomer(w http.ResponseWriter, r *http.Request) {
	customer := &models.Customer{}
	if err := render.Bind(r, customer); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddItem(customer); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, customer); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := dbInstance.getAllCustomers()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, customers); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := r.Context().Value(customerIDKey).(int)
	customer, err := dbInstance.GetCustomerById(CustomerID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &customer); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	customerId := r.Context().Value(customerIDKey).(int)
	err := dbInstance.deleteCustomer(customerId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	customerId := r.Context().Value(customerIDKey).(int)
	customerData := models.Customer{}
	if err := render.Bind(r, &customerData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	customer, err := dbInstance.updateCustomer(customerId, customerData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &customer); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
func saveBalance(w http.ResponseWriter, r *http.Request) {
	customerID := r.Context().Value(customerIDKey).(int)
	customer, err := dbInstance.GetCustomerById(CustomerID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &customer); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
