package dto

type CreateWeaponInputDTO struct {
	Name         string `json:"name"`
	Price        int32  `json:"price"`
	Weight       int32  `json:"weight"`
	Acc          int32  `json:"acc"`
	Str          int32  `json:"str"`
	Con          int32  `json:"con"`
	Dex          int32  `json:"dex"`
	Men          int32  `json:"men"`
	Wit          int32  `json:"wit"`
	Critical     int32  `json:"critical"`
	MagicAttack  int32  `json:"m_atk"`
	MagicSpeed   int32  `json:"m_spd"`
	PowerAttack  int32  `json:"p_atk"`
	PowerSpeed   int32  `json:"p_spd"`
	WeaponTypeID int32  `json:"weapon_type_id"`
}

type WeaponOutputDTO struct {
	Name        string              `json:"name"`
	Price       int32               `json:"price"`
	Weight      int32               `json:"weight"`
	Acc         int32               `json:"acc"`
	Str         int32               `json:"str"`
	Con         int32               `json:"con"`
	Dex         int32               `json:"dex"`
	Men         int32               `json:"men"`
	Wit         int32               `json:"wit"`
	Critical    int32               `json:"critical"`
	MagicAttack int32               `json:"m_atk"`
	MagicSpeed  int32               `json:"m_spd"`
	PowerAttack int32               `json:"p_atk"`
	PowerSpeed  int32               `json:"p_spd"`
	WeaponType  WeaponTypeOutputDTO `json:"weapon_type"`
}

type CreateWeaponTypeInputDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type WeaponTypeOutputDTO struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateWeaponTypeInputDTO struct {
	Description string `json:"description"`
}
