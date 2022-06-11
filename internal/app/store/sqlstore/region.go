package sqlstore

import (
	"context"
	"hack/internal/app/model"
)

type RegionRepository struct {
	store *Store
}

func (r *RegionRepository) PredictList(ctx context.Context) ([]model.RegionPredict, error) {
	list := []model.RegionPredict{}

	rows, err := r.store.db.QueryContext(ctx,
		`SELECT
    id, position, city, current_client_index, predict_client_index, predict_arpu
  FROM 
		region_predict
	`)

	if err != nil {
		return list, err
	}

	for rows.Next() {
		var p model.RegionPredict
		err = rows.Scan(&p.ID,
			&p.Position,
			&p.City,
			&p.CurrentClientIndex,
			&p.PredictClientIndex,
			&p.PredictArpu,
		)

		if err != nil {
			return list, err
		}
		list = append(list, p)
	}

	return list, nil
}
