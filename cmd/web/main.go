package main

import (
	"log"
	"net/http"
	"time"

	handler "github.com/ShankaranarayananBR/bookings/pkg/handlers"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	//m := pat.New()
	//m.Get("/", http.HandlerFunc(handler.Repo.Home))
	//m.Get("/about", http.HandlerFunc(handler.Repo.About))
	//http.Handle("/", m)
	//http.Handle("/about", m)

	m := chi.NewRouter()
	m.Use(middleware.Recoverer)
	m.Use(noserve)
	m.Get("/", handler.Repo.Home)
	m.Get("/about", handler.Repo.About)
	http.Handle("/", m)
	http.Handle("/about", m)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
