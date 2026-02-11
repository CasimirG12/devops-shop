package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CasimirG12/devops-shop/internal/handlers"
)

func main() {

	mux := http.NewServeMux()
	// Auth Service
	mux.HandleFunc("/auth/login", handlers.AuthLoginHandler)
	mux.HandleFunc("/auth/logout", handlers.AuthLogoutHandler)
	// Product Service
	mux.HandleFunc("/products", handlers.ProductListHandler)
	mux.HandleFunc("/products/{id}", handlers.ProductDetailHandler)
	// Checkout Service
	mux.HandleFunc("/checkout/placeorder", handlers.CheckoutPlaceOrderHandler)
	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
