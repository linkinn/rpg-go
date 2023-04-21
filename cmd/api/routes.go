package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"rpg-go/internal/handlers"
	"rpg-go/internal/service"
)

type HandlerSetup struct {
	WeaponHandler *handlers.WeaponHandler
}

var handlerSetup HandlerSetup

func routes() http.Handler {
	handlerStart()
	mux := chi.NewRouter()

	// default middleware
	mux.Use(middleware.Logger)

	// TODO: create routes to armors
	// -> armors routes

	// TODO: create routes to accounts
	// -> accounts routes

	// TODO: create routes to characters
	// -> characters routes

	// TODO: create routes to clans
	// -> clans routes

	// TODO: create routes to classes
	// -> classes routes

	// TODO: create routes to races
	// -> races routes
	mux.Post("/weapons", handlerSetup.WeaponHandler.CreateWeapon)
	mux.Get("/weapons", handlerSetup.WeaponHandler.ListWeapon)

	// TODO: create routes to weapons
	// -> weapons routes
	mux.Post("/weapons_types", handlerSetup.WeaponHandler.CreateWeaponType)
	mux.Get("/weapons_types", handlerSetup.WeaponHandler.ListWeaponType)
	mux.Get("/weapons_types/{id}", handlerSetup.WeaponHandler.GetWeaponTypeByID)
	mux.Get("/weapons_types/{name}", handlerSetup.WeaponHandler.GetWeaponTypeByName)
	mux.Patch("/weapons_types/{id}", handlerSetup.WeaponHandler.UpdateWeaponType)
	mux.Delete("/weapons_types/{id}", handlerSetup.WeaponHandler.DeleteWeaponType)

	return mux
}

func handlerStart() {
	ctx := context.Background()

	// -> weapons setup
	weaponService := service.NewWeaponsService(app.Querier)
	webPlayerHandler := handlers.NewWeaponHandler(ctx, weaponService)

	// -> set handlers
	handlerSetup = HandlerSetup{
		WeaponHandler: webPlayerHandler,
	}
}
