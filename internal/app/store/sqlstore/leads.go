package sqlstore

import (
	"context"
	"database/sql"
	"hack/internal/app/model"
)

type LeadRepository struct {
	store *Store
}

func (r *LeadRepository) GetList(ctx context.Context) ([]model.Lead, error) {
	list := []model.Lead{}
	var utm_source sql.NullString
	var utm_content sql.NullString
	var utm_campaign sql.NullString
	rows, err := r.store.db.QueryContext(ctx,
		`SELECT
    id, client_id, product_category_name, utm_source, utm_content, utm_campaign, date, cpc
  FROM 
		leads
	`)

	if err != nil {
		return list, err
	}

	for rows.Next() {
		var p model.Lead
		err = rows.Scan(
			&p.ID,
			&p.ClientID,
			&p.ProductCategoryName,
			&utm_source,
			&utm_content,
			&utm_campaign,
			&p.Date,
			&p.CPC,
		)
		if err != nil {
			return list, err
		}
		if utm_source.Valid {
			p.UTMSource = utm_source.String
		}
		if utm_content.Valid {
			p.UTMContent = utm_content.String
		}
		if utm_campaign.Valid {
			p.UTMCampaing = utm_campaign.String
		}
		list = append(list, p)
	}

	return list, nil
}
