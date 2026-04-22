package persistence

import (
	"config-service/internal/domain/model"
	"config-service/internal/domain/ports"
	"database/sql"
	"fmt"
	"time"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.OutputPort {
	return &repository{db: db}
}

func (r *repository) Save(site model.MonitoredURL) error {
	q := `INSERT INTO config_urls (url, review_time, creation_date, modify_date) VALUES ($1, $2, $3, $4)` // PostgreSQL

	_, err := r.db.Exec(q, site.URL, int(site.ReviewTime.Seconds()), site.CreationDate, site.ModifyDate)

	return err
}

func (r *repository) FindByID(id int) (model.MonitoredURL, error) {
	q := "SELECT id, url, review_time, creation_date, modify_date FROM config_urls WHERE id = $1" // PostgreSQL

	var site model.MonitoredURL
	var reviewTime int
	err := r.db.QueryRow(q, id).Scan(&site.ID, &site.URL, &reviewTime, &site.CreationDate, &site.ModifyDate)
	if err != nil {
		return model.MonitoredURL{}, err
	}
	site.ReviewTime = time.Duration(reviewTime) * time.Second

	if err == sql.ErrNoRows {
		return model.MonitoredURL{}, fmt.Errorf("site no encontrado")
	}

	return site, nil
}

func (r *repository) FindAll() ([]model.MonitoredURL, error) {
	q := "SELECT id, url, review_time, creation_date, modify_date FROM config_urls" // PostgreSQL

	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sites []model.MonitoredURL
	for rows.Next() {
		var site model.MonitoredURL
		var reviewTime int
		if err := rows.Scan(&site.ID, &site.URL, &reviewTime, &site.CreationDate, &site.ModifyDate); err != nil {
			return nil, err
		}
		site.ReviewTime = time.Duration(reviewTime) * time.Second
		sites = append(sites, site)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sites, nil
}

func (r *repository) Update(site model.MonitoredURL) error {
	q := `UPDATE config_urls SET url = $1, review_time = $2, modify_date = $3 WHERE id = $4` // PostgreSQL

	_, err := r.db.Exec(q, site.URL, int(site.ReviewTime.Seconds()), site.ModifyDate, site.ID)
	return err
}

func (r *repository) Delete(id int) error {
	q := "DELETE FROM config_urls WHERE id = $1" // PostgreSQL
	_, err := r.db.Exec(q, id)
	return err
}
