package router

import (
	"github.com/VineetBavniya/PostgreSQL-GO/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET","OPTIONS")
	router.HandleFunc("/api/stock", middleware.GetAllStocks).Methods("GET","OPTIONS")
	router.HandleFunc("/api/newstock", middleware.CreateNewStock).Methods("POST","OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.DeleteStock).Methods("DELETE","OPTIONS")

	return router
}