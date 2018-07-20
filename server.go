package main

import (
	"log"
	"net/http"

	"github.com/cobinix/01-login/routes/callback"
	"github.com/cobinix/01-login/routes/home"
	"github.com/cobinix/01-login/routes/login"
	"github.com/cobinix/01-login/routes/logout"
	"github.com/cobinix/01-login/routes/middlewares"
	"github.com/cobinix/01-login/routes/user"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", home.HomeHandler)
	r.HandleFunc("/login", login.LoginHandler)
	r.HandleFunc("/logout", logout.LogoutHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/", r)
	log.Print("Server listening on http://localhost:3000/")
	http.ListenAndServe("0.0.0.0:3000", nil)
}
