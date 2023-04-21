package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"rpg-go/internal/dto"
	"rpg-go/internal/service"
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = wh.svc.CreateWeapon(wh.ctx, input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (wh *WeaponHandler) ListWeapon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	wt, err := wh.svc.ListWeapon(wh.ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wt)
}

// -> Weapon Type
func (wh *WeaponHandler) CreateWeaponType(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateWeaponTypeInputDTO

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = wh.svc.CreateWeaponType(wh.ctx, input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (wh *WeaponHandler) ListWeaponType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	wt, err := wh.svc.ListWeaponType(wh.ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wt)
}

func (wh *WeaponHandler) GetWeaponTypeByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (wh *WeaponHandler) GetWeaponTypeByName(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (wh *WeaponHandler) UpdateWeaponType(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (wh *WeaponHandler) DeleteWeaponType(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
