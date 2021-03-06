package main

import (
	"log"
	"net/http"

	"github.com/noobs9/calico-server/pkg/auth"
	"github.com/noobs9/calico-server/pkg/controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/user/{id}", auth.JwtMiddleware.Handler(controller.UserGetByID)).Methods("GET")
	r.Handle("/user", auth.JwtMiddleware.Handler(controller.UserGet)).Methods("GET")
	r.Handle("/user", controller.UserPost).Methods("POST")
	r.Handle("/user/{id}", auth.JwtMiddleware.Handler(auth.OnlyPersonMiddleware(controller.UserPut))).Methods("PUT")
	r.Handle("/user/{id}", auth.JwtMiddleware.Handler(auth.OnlyPersonMiddleware(controller.UserDelete))).Methods("DELETE")
	r.Handle("/todo/{id}", auth.JwtMiddleware.Handler(controller.TodoGetByID)).Methods("GET")
	r.Handle("/todo", auth.JwtMiddleware.Handler(controller.TodoGet)).Methods("GET")
	r.Handle("/todo", auth.JwtMiddleware.Handler(controller.TodoPost)).Methods("POST")
	r.Handle("/todo/{id}", auth.JwtMiddleware.Handler(controller.TodoPut)).Methods("PUT")
	r.Handle("/todo/{id}", auth.JwtMiddleware.Handler(controller.TodoDelete)).Methods("DELETE")

	r.HandleFunc("/auth", controller.GetTokenHandler).Methods("POST")
	r.Handle("/auth/test", auth.JwtMiddleware.Handler(auth.AuthTest))

	r.HandleFunc("/ping", pingHandler)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe failed: ", err)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}
