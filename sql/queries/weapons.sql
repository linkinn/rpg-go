-- name: CreateWeapons :exec
INSERT INTO weapons (weapon_type_id, name, price, weight, acc, str, con, dex, men, wit, critical, magic_attack, magic_speed, power_attack, power_speed)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);

-- name: ListWeapons :many
SELECT *, wt.id, wt.name, wt.description FROM weapons w LEFT JOIN weapons_type wt ON w.weapon_type_id = wt.id;

-- name: CreateWeaponsType :exec
INSERT INTO weapons_type (name, description) VALUES ($1, $2);

-- name: ListWeaponsType :many
SELECT * FROM weapons_type;

-- name: GetWeaponsTypeByID :one
SELECT * FROM weapons_type WHERE id = $1;

-- name: GetWeaponsTypeByName :one
SELECT * FROM weapons_type WHERE name = $1;

-- name: UpdateWeaponsType :exec
UPDATE weapons_type SET description = $1 WHERE id = $2;

-- name: DeleteWeaponsType :exec
DELETE FROM weapons_type WHERE id = $1;
