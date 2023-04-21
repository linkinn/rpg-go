package service

import (
	"context"
	db "database/sql"
	"rpg-go/internal/dto"
	"rpg-go/internal/repository/sql"
)

type WeaponsInterface interface {
	CreateWeapon(ctx context.Context, in dto.CreateWeaponInputDTO) error
	ListWeapon(ctx context.Context) ([]dto.WeaponOutputDTO, error)

	CreateWeaponType(ctx context.Context, in dto.CreateWeaponTypeInputDTO) error
	ListWeaponType(ctx context.Context) ([]dto.WeaponTypeOutputDTO, error)
	GetWeaponTypeByID(ctx context.Context, id int32) (dto.WeaponTypeOutputDTO, error)
	GetWeaponTypeByName(ctx context.Context, name string) (dto.WeaponTypeOutputDTO, error)
	UpdateWeaponType(ctx context.Context, id int32, in dto.UpdateWeaponTypeInputDTO) error
	DeleteWeaponType(ctx context.Context, id int32) error
}

type WeaponsService struct {
	repo sql.Querier
}

func NewWeaponsService(repo sql.Querier) *WeaponsService {
	return &WeaponsService{repo: repo}
}

// ----------
// -> Weapons
// ----------
func (ws *WeaponsService) CreateWeapon(ctx context.Context, in dto.CreateWeaponInputDTO) error {
	args := sql.CreateWeaponsParams{
		Name:        in.Name,
		Price:       in.Price,
		Weight:      in.Weight,
		Acc:         in.Acc,
		Str:         in.Str,
		Con:         in.Con,
		Dex:         in.Dex,
		Men:         in.Men,
		Wit:         in.Wit,
		Critical:    in.Critical,
		MagicAttack: in.MagicAttack,
		MagicSpeed:  in.MagicSpeed,
		PowerAttack: in.PowerAttack,
		PowerSpeed:  in.PowerSpeed,
	}

	if in.WeaponTypeID > 0 {
		wt, err := ws.GetWeaponTypeByID(ctx, in.WeaponTypeID)
		if err != nil {
			return err
		}

		args.WeaponTypeID = db.NullInt32{Int32: wt.ID, Valid: true}
	}

	err := ws.repo.CreateWeapons(ctx, args)
	return err
}

func (ws *WeaponsService) ListWeapon(ctx context.Context) ([]dto.WeaponOutputDTO, error) {
	w, err := ws.repo.ListWeapons(ctx)
	if err != nil {
		return []dto.WeaponOutputDTO{}, err
	}

	var out []dto.WeaponOutputDTO
	for _, v := range w {
		var wt dto.WeaponTypeOutputDTO

		if v.ID_2.Int32 > 0 {
			wt = dto.WeaponTypeOutputDTO{
				ID:          v.ID_2.Int32,
				Name:        v.Name_2.String,
				Description: v.Description_2.String,
			}
		}
		out = append(out, dto.WeaponOutputDTO{
			Name:        v.Name,
			Price:       v.Price,
			Weight:      v.Weight,
			Acc:         v.Acc,
			Str:         v.Str,
			Con:         v.Con,
			Dex:         v.Dex,
			Men:         v.Men,
			Wit:         v.Wit,
			Critical:    v.Critical,
			MagicAttack: v.MagicAttack,
			MagicSpeed:  v.MagicSpeed,
			PowerAttack: v.PowerAttack,
			PowerSpeed:  v.PowerSpeed,
			WeaponType:  wt,
		})
	}
	return out, nil
}

// ---------------
// -> Weapons Type
// ---------------
func (ws *WeaponsService) CreateWeaponType(ctx context.Context, in dto.CreateWeaponTypeInputDTO) error {
	args := sql.CreateWeaponsTypeParams{
		Name:        in.Name,
		Description: db.NullString{String: in.Description, Valid: true},
	}

	err := ws.repo.CreateWeaponsType(ctx, args)
	return err
}

func (ws *WeaponsService) ListWeaponType(ctx context.Context) ([]dto.WeaponTypeOutputDTO, error) {
	wt, err := ws.repo.ListWeaponsType(ctx)
	if err != nil {
		return []dto.WeaponTypeOutputDTO{}, err
	}

	var out []dto.WeaponTypeOutputDTO
	for _, v := range wt {
		out = append(out, dto.WeaponTypeOutputDTO{ID: v.ID, Name: v.Name, Description: v.Description.String})
	}

	return out, nil
}

func (ws *WeaponsService) GetWeaponTypeByID(ctx context.Context, id int32) (dto.WeaponTypeOutputDTO, error) {
	wt, err := ws.repo.GetWeaponsTypeByID(ctx, id)
	if err != nil {
		return dto.WeaponTypeOutputDTO{}, err
	}

	out := dto.WeaponTypeOutputDTO{
		ID:          wt.ID,
		Name:        wt.Name,
		Description: wt.Description.String,
	}
	return out, nil
}

func (ws *WeaponsService) GetWeaponTypeByName(ctx context.Context, name string) (dto.WeaponTypeOutputDTO, error) {
	wt, err := ws.repo.GetWeaponsTypeByName(ctx, name)
	if err != nil {
		return dto.WeaponTypeOutputDTO{}, err
	}

	out := dto.WeaponTypeOutputDTO{
		ID:          wt.ID,
		Name:        wt.Name,
		Description: wt.Description.String,
	}
	return out, nil
}

func (ws *WeaponsService) UpdateWeaponType(ctx context.Context, id int32, in dto.UpdateWeaponTypeInputDTO) error {
	wt, err := ws.repo.GetWeaponsTypeByID(ctx, id)
	if err != nil {
		return err
	}

	arg := sql.UpdateWeaponsTypeParams{
		Description: db.NullString{String: in.Description, Valid: true},
		ID:          wt.ID,
	}

	err = ws.repo.UpdateWeaponsType(ctx, arg)
	return err
}

func (ws *WeaponsService) DeleteWeaponType(ctx context.Context, id int32) error {
	wt, err := ws.repo.GetWeaponsTypeByID(ctx, id)
	if err != nil {
		return err
	}

	err = ws.repo.DeleteWeaponsType(ctx, wt.ID)
	return err
}
