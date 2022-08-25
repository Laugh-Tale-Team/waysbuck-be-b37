package routes

import (
	"waysbuck/handlers"
	"waysbuck/pkg/middleware"
	"waysbuck/pkg/mysql"
	"waysbuck/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryTopping(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/carts",h.FindCarts).Methods("GET")
	r.HandleFunc("/cart/{id}",h.GetCart).Methods("GET")
	r.HandleFunc("/carts-id", middleware.Auth(h.FindCartsById)).Methods("GET")
	r.HandleFunc("/cart", middleware.Auth(h.CreateCart)).Methods("POST")
	r.HandleFunc("/cart_id", middleware.Auth(h.UpdateCart)).Methods("PATCH")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.DeleteCart)).Methods("DELETE")
}