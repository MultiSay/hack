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
    id, position, city, current_client_index, predict_client_index, predict_arpu, predict_score
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
			&p.PredictScore,
		)

		if err != nil {
			return list, err
		}
		list = append(list, p)
	}

	return list, nil
}

func (r *RegionRepository) PredictListUpdate(ctx context.Context, list []model.RegionPredict) error {
	_, err := r.store.db.ExecContext(ctx, `TRUNCATE region_predict`)
	if err != nil {
		return err
	}
	for _, v := range list {
		_, err := r.store.db.ExecContext(ctx,
			`INSERT INTO "public"."region_predict" 
				("position", "city", "predict_score", "current_client_index", "predict_client_index") 
			VALUES 
				($1, $2, $3, $4, $5);
		`, v.Position, v.City, v.PredictScore, v.CurrentClientIndex, v.PredictClientIndex)
		if err != nil {
			return err
		}
	}

	return nil
}
