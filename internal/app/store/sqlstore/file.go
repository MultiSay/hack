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

func (r *FileRepository) Create(ctx context.Context, p *model.File) error {
	return r.store.db.QueryRowContext(ctx,
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
}

func (r *FileRepository) Update(ctx context.Context, p *model.File) error {
	return r.store.db.QueryRow(
		`UPDATE files SET
			(
				send_at,
				receive_at
			) =
		 	(
				 $1,
				 $2
			)
		WHERE
			id=$3`,
		p.SendAt,
		p.ReceiveAt,
	).Err()
}

func (r *FileRepository) GetByID(ctx context.Context, paymentID int) (*model.File, error) {
	p := &model.File{}
	var send_at sql.NullTime
	var receive_at sql.NullTime
	if err := r.store.db.QueryRowContext(ctx,
		`SELECT 
			id,
  		name,
  		create_at,
 			send_at,
  		receive_at,
			size
		FROM 
			files
		WHERE
			id = $1`,
		paymentID,
	).Scan(
		&p.ID,
		&p.Name,
		&p.CreateAt,
		&send_at,
		&receive_at,
		&p.Size,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	p.SendAt = send_at.Time
	p.ReceiveAt = receive_at.Time
	return p, nil
}
