package router

import (
	"encoding/json"
	"go-stocks/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthCheck(w http.ResponseWriter, r *http.Request)  {
	res:="Api is running fine"
	json.NewEncoder(w).Encode(res);
}

func Router() *mux.Router{
	router:=mux.NewRouter();
	router.HandleFunc("/health",HealthCheck).Methods("GET","OPTIONS")
	router.HandleFunc("/api/stock/{id}",middleware.GetStock).Methods("GET","OPTIONS");
	router.HandleFunc("/api/stock",middleware.GetAllStocks).Methods("GET","OPTIONS");
	router.HandleFunc("/api/newstock",middleware.CreateStock).Methods("POST","OPTIONS");
	router.HandleFunc("/api/stock/{id}",middleware.UpdateStock).Methods("PUT","OPTIONS");
	router.HandleFunc("/api/deletestock/{id}",middleware.DeleteStock).Methods("DELETE","OPTIONS");
	router.HandleFunc("/api/deleteall",middleware.DeleteAllStock).Methods("DELETE","OPTIONS")
	return router
}