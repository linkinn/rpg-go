// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: accounts.sql

package sql

import (
	"context"
)

const ListAccounts = `-- name: ListAccounts :many
SELECT id, nickname, password, email, last_active, last_ip, updated_at, created_at FROM accounts
`

func (q *Queries) ListAccounts(ctx context.Context) ([]Accounts, error) {
	rows, err := q.db.Query(ctx, ListAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Accounts{}
	for rows.Next() {
		var i Accounts
		if err := rows.Scan(
			&i.ID,
			&i.Nickname,
			&i.Password,
			&i.Email,
			&i.LastActive,
			&i.LastIp,
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
