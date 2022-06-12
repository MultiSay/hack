package sqlstore

import (
	"context"
	"hack/internal/app/model"
)

type CompaignRepository struct {
	store *Store
}

func (r *CompaignRepository) GetList(ctx context.Context) ([]model.Compaign, error) {
	list := []model.Compaign{}
	rows, err := r.store.db.QueryContext(ctx,
		`SELECT
    id, utm_campaign, gender, age_from, age_to, city, theme
  FROM 
		leads
	`)

	if err != nil {
		return list, err
	}

	for rows.Next() {
		var p model.Compaign
		err = rows.Scan(
			&p.ID,
			&p.UTMCampaign,
			&p.Gender,
			&p.AgeFrom,
			&p.AgeTo,
			&p.City,
			&p.Theme,
		)
		if err != nil {
			return list, err
		}
		list = append(list, p)
	}

	return list, nil
}
