package sqlstore

import (
	"context"
	"database/sql"
	"hack/internal/app/model"
	"hack/internal/app/store"
	"time"
)

type FileRepository struct {
	store *Store
}

func (r *FileRepository) Create(ctx context.Context, p model.File) (model.File, error) {
	err := r.store.db.QueryRowContext(ctx,
		`INSERT INTO files 
			(
  			"name",
 				"create_at",
				"size"
			)
		 VALUES 
		 	(
				$1,
				$2,
				$3
			) RETURNING id`,
		p.Name,
		time.Now(),
		p.Size,
	).Scan(&p.ID)
	return p, err
}

func (r *FileRepository) Update(ctx context.Context, p model.File) error {
	return r.store.db.QueryRow(
		`UPDATE files SET
			(
				send_at,
				received_at
			) =
		 	(
				 $1,
				 $2
			)
		WHERE
			id=$3`,
		p.SendAt,
		p.ReceivedAt,
	).Err()
}

func (r *FileRepository) GetByID(ctx context.Context, fileID int) (model.File, error) {
	p := model.File{}
	var send_at sql.NullTime
	var received_at sql.NullTime
	if err := r.store.db.QueryRowContext(ctx,
		`SELECT 
			id,
  		name,
  		create_at,
 			send_at,
  		received_at,
			size, 
			status
		FROM 
			files
		WHERE
			id = $1`,
		fileID,
	).Scan(
		&p.ID,
		&p.Name,
		&p.CreateAt,
		&send_at,
		&received_at,
		&p.Size,
		&p.Status,
	); err != nil {
		if err == sql.ErrNoRows {
			return p, store.ErrRecordNotFound
		}
		return p, err
	}
	if send_at.Valid {
		p.SendAt = send_at.Time
	}
	if received_at.Valid {
		p.ReceivedAt = received_at.Time
	}
	return p, nil
}

func (r *FileRepository) GetLast(ctx context.Context) (model.File, error) {
	p := model.File{}
	var send_at sql.NullTime
	var received_at sql.NullTime
	if err := r.store.db.QueryRowContext(ctx,
		`SELECT 
			id,
  		name,
  		create_at,
 			send_at,
  		received_at,
			size, 
			status
		FROM 
			files
		ORDER BY id DESC
		LIMIT 1
			`,
	).Scan(
		&p.ID,
		&p.Name,
		&p.CreateAt,
		&send_at,
		&received_at,
		&p.Size,
		&p.Status,
	); err != nil {
		if err == sql.ErrNoRows {
			return p, store.ErrRecordNotFound
		}
		return p, err
	}
	if send_at.Valid {
		p.SendAt = send_at.Time
	}
	if received_at.Valid {
		p.ReceivedAt = received_at.Time
	}
	return p, nil
}

func (r *FileRepository) GetList(ctx context.Context) ([]model.File, error) {
	list := []model.File{}
	var send_at sql.NullTime
	var received_at sql.NullTime
	rows, err := r.store.db.QueryContext(ctx,
		`SELECT
    id, name, create_at, send_at, received_at, size, status
  FROM 
		files
	`)

	if err != nil {
		return list, err
	}

	for rows.Next() {
		var p model.File
		err = rows.Scan(
			&p.ID,
			&p.Name,
			&p.CreateAt,
			&send_at,
			&received_at,
			&p.Size,
			&p.Status,
		)
		if err != nil {
			return list, err
		}
		if send_at.Valid {
			p.SendAt = send_at.Time
		}
		if received_at.Valid {
			p.ReceivedAt = received_at.Time
		}
		list = append(list, p)
	}

	return list, nil
}
