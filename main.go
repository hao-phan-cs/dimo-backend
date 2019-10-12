package main

import (
	"dimo-backend/config"
	"dimo-backend/handlers"
	"dimo-backend/middleware"
	"fmt"
	"github.com/didip/tollbooth"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/user/register", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/login", handlers.AuthenticateUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", handlers.GetUserInfo).Methods("GET")
	router.Handle("/", tollbooth.LimitHandler(tollbooth.NewLimiter(1, nil),
		http.HandlerFunc(handlers.Default))).Methods("GET")

	router.Handle("/api/store/id={id}&lat={lat}&long={long}",
		tollbooth.LimitHandler(tollbooth.NewLimiter(2, nil),
			http.HandlerFunc(handlers.GetStoreById))).Methods("GET")
	router.Handle("/api/store/user_id={user_id}&lat={lat}&long={long}&km_limit={km_limit}",
		tollbooth.LimitHandler(tollbooth.NewLimiter(2, nil),
			http.HandlerFunc(handlers.GetStoresByDistLimit))).Methods("GET")

	router.HandleFunc("/api/review/create", handlers.CreateReview).Methods("POST")
	router.HandleFunc("/api/review/delete", handlers.DeleteReview).Methods("DELETE")

	router.Use(middleware.Recovery)

	port := config.ApiPort
	fmt.Println("Listening on port:", port)
	err := http.ListenAndServe(":"+strconv.FormatInt(int64(port), 10), router)
	if err != nil {
		fmt.Print(err)
	}
}
