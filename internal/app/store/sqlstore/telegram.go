package sqlstore

import (
	"context"
	"hack/internal/app/model"
)

type TelegramRepository struct {
	store *Store
}

func (r *TelegramRepository) GetList(ctx context.Context) ([]model.Telegram, error) {
	list := []model.Telegram{}
	rows, err := r.store.db.QueryContext(ctx,
		`SELECT
    id, name_id, name, n_subscribers, category
  FROM 
		telegram
	`)

	if err != nil {
		return list, err
	}

	for rows.Next() {
		var p model.Telegram
		err = rows.Scan(
			&p.ID,
			&p.NameID,
			&p.Name,
			&p.NSubscribers,
			&p.Category,
		)
		if err != nil {
			return list, err
		}
		list = append(list, p)
	}

	return list, nil
}
