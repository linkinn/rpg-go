package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"rpg-go/internal/dto"
	"rpg-go/internal/helpers"
	"rpg-go/internal/service"

	"github.com/go-chi/render"
)

type WeaponHandler struct {
	ctx context.Context
	svc *service.WeaponsService
}

func NewWeaponHandler(ctx context.Context, svc *service.WeaponsService) *WeaponHandler {
	return &WeaponHandler{
		ctx: ctx,
		svc: svc,
	}
}

// -> Weapon
func (wh *WeaponHandler) CreateWeapon(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateWeaponInputDTO

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		wrapErr := helpers.WrapError(err, http.StatusBadRequest, "WeaponHandler.CreateWeapon")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, wrapErr)
		return
	}
	defer r.Body.Close()

	err = wh.svc.CreateWeapon(wh.ctx, input)
	if err != nil {
		wrapErr := helpers.WrapError(err, http.StatusBadRequest, "WeaponHandler.CreateWeapon")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, wrapErr)
		return
	}

	render.Status(r, http.StatusCreated)
}

func (wh *WeaponHandler) ListWeapon(w http.ResponseWriter, r *http.Request) {
	weapon, err := wh.svc.ListWeapon(wh.ctx)
	if err != nil {
		wrapErr := helpers.WrapError(err, http.StatusBadRequest, "WeaponHandler.ListWeapon")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, wrapErr)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, weapon)
}

// -> Weapon Type
func (wh *WeaponHandler) CreateWeaponType(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateWeaponTypeInputDTO

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		wrapErr := helpers.WrapError(err, http.StatusBadRequest, "WeaponHandler.CreateWeaponType")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, wrapErr)
		return
	}
	defer r.Body.Close()

	err = wh.svc.CreateWeaponType(wh.ctx, input)
	if err != nil {
		wrapErr := helpers.WrapError(err, http.StatusBadRequest, "WeaponHandler.CreateWeaponType")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, wrapErr.ErrorJSON())
		return
	}

	render.Status(r, http.StatusCreated)
}

func (wh *WeaponHandler) ListWeaponType(w http.ResponseWriter, r *http.Request) {
	wt, err := wh.svc.ListWeaponType(wh.ctx)
	if err != nil {
		wrapErr := helpers.WrapError(err, http.StatusBadRequest, "WeaponHandler.ListWeaponType")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, wrapErr)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, wt)
}

func (wh *WeaponHandler) GetWeaponTypeByID(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
}

func (wh *WeaponHandler) GetWeaponTypeByName(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
}

func (wh *WeaponHandler) UpdateWeaponType(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
}

func (wh *WeaponHandler) DeleteWeaponType(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
}
