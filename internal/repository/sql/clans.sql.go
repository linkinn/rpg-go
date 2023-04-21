// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: clans.sql

package sql

import (
	"context"
)

const ListClans = `-- name: ListClans :many
SELECT id, name, motto, level, reputation_score, leader_id, max_online_member, prev_max_online_member, updated_at, created_at FROM clans
`

func (q *Queries) ListClans(ctx context.Context) ([]Clans, error) {
	rows, err := q.db.Query(ctx, ListClans)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Clans{}
	for rows.Next() {
		var i Clans
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Motto,
			&i.Level,
			&i.ReputationScore,
			&i.LeaderID,
			&i.MaxOnlineMember,
			&i.PrevMaxOnlineMember,
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