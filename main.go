package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	driver "agendamiento/common"
	ac "agendamiento/controllers/agendadas_controller"
	hc "agendamiento/controllers/horario_controller"
)

func main() {
	/*
		dbName := os.Getenv("DB_NAME")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
	*/

	//connection, err := driver.ConnectSQL(dbHost, dbPort, "Fernando", dbPass, dbName)
	connection, err := driver.ConnectSQL("localhost", "3003", "Fernando", "2123", "agendamiento")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	horarioHandler := hc.NewHorarioHandler(connection)
	agendadasHandler := ac.NewAgendadasHandler(connection)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/horario", horarioRouter(horarioHandler))
		rt.Mount("/agendado", agendadasRouter(agendadasHandler))
	})

	fmt.Println("Server listen at :5003")
	http.ListenAndServe(":5003", r)
}

// A completely separate router for posts routes
func horarioRouter(horarioHandler *hc.Horario) http.Handler {
	r := chi.NewRouter()
	r.Get("/", horarioHandler.Fetch)
	r.Get("/{id:[0-9]+}", horarioHandler.GetByID)
	r.Post("/", horarioHandler.Create)
	r.Put("/{id:[0-9]+}", horarioHandler.Update)
	r.Delete("/{id:[0-9]+}", horarioHandler.Delete)

	return r
}

func agendadasRouter(agendadasHandler *ac.Agendadas) http.Handler {
	r := chi.NewRouter()
	r.Get("/", agendadasHandler.Fetch)
	r.Get("/{id:[0-9]+}", agendadasHandler.GetByID)
	r.Post("/", agendadasHandler.Create)
	r.Put("/{id:[0-9]+}", agendadasHandler.Update)
	r.Delete("/{id:[0-9]+}", agendadasHandler.Delete)

	return r
}
