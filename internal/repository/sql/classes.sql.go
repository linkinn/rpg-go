// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: classes.sql

package sql

import (
	"context"
)

const ListClasses = `-- name: ListClasses :many
SELECT id, race_id, name, acc, str, con, dex, men, wit, critical, evasion, magic_attack, magic_defense, magic_speed, power_attack, power_defense, power_speed, updated_at, created_at FROM classes
`

func (q *Queries) ListClasses(ctx context.Context) ([]Classes, error) {
	rows, err := q.db.Query(ctx, ListClasses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Classes{}
	for rows.Next() {
		var i Classes
		if err := rows.Scan(
			&i.ID,
			&i.RaceID,
			&i.Name,
			&i.Acc,
			&i.Str,
			&i.Con,
			&i.Dex,
			&i.Men,
			&i.Wit,
			&i.Critical,
			&i.Evasion,
			&i.MagicAttack,
			&i.MagicDefense,
			&i.MagicSpeed,
			&i.PowerAttack,
			&i.PowerDefense,
			&i.PowerSpeed,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}