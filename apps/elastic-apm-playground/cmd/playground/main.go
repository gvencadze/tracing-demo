package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	apm_playground "github.com/gvencadze/apm-playground"
	"github.com/gvencadze/apm-playground/internal/profile"
	"github.com/gvencadze/apm-playground/pkg/database"
	"github.com/gvencadze/apm-playground/pkg/middleware"
	"github.com/pressly/goose/v3"
	"go.elastic.co/apm/module/apmgorilla/v2"
	"log"
	"net/http"
	"strconv"
)

func migrate(db *sql.DB) error {
	goose.SetBaseFS(apm_playground.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("goose set dialect: %w", err)
	}

	log.Println(goose.Status(db, "migrations"))

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("goose up: %w", err)
	}

	return nil
}

func main() {
	conn, db := database.Connect()

	if err := migrate(db); err != nil {
		log.Println(err)
		return
	}

	profileService := profile.NewIProfileWithTracing(
		profile.New(conn),
		"internal",
	)

	r := mux.NewRouter()
	apmgorilla.Instrument(r)

	r.Use(middleware.RequestIDMiddleware())
	r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Println(err)
			return
		}

		res, err := profileService.Get(r.Context(), id)
		if err != nil {
			log.Println(err)
			return
		}

		_ = json.NewEncoder(w).Encode(res)
		return
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
