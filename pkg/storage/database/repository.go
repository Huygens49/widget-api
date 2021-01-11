package database

import (
	"database/sql"

	"github.com/Huygens49/widget-api/pkg/read"
	"github.com/Huygens49/widget-api/pkg/write"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

func (r Repository) GetAllWidgets() ([]read.Widget, error) {
	var widgets []read.Widget
	query := "select id, created_at, updated_at, description, owner, value from widgets where deleted_at is null"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var w read.Widget
		err := rows.Scan(&w.ID, &w.CreatedAt, &w.UpdatedAt, &w.Description, &w.Owner, &w.Value)

		if err != nil {
			return nil, err
		}

		widgets = append(widgets, w)
	}

	return widgets, nil
}

func (r Repository) GetWidget(id uint) (read.Widget, error) {
	var w read.Widget
	query := "select id, created_at, updated_at, description, owner, value from widgets where id = $1 and deleted_at is null"

	err := r.db.QueryRow(query, id).Scan(&w.ID, &w.CreatedAt, &w.UpdatedAt, &w.Description, &w.Owner, &w.Value)
	if err != nil {
		return read.Widget{}, err
	}

	return w, nil
}

func (r Repository) AddWidget(widget write.Widget) (read.Widget, error) {
	cmd := "insert into widgets(created_at, updated_at, description, owner, value) values(current_timestamp, current_timestamp, $1, $2, $3) returning id"
	id := uint(0)

	err := r.db.QueryRow(cmd, widget.Description, widget.Owner, widget.Value).Scan(&id)
	if err != nil {
		return read.Widget{}, err
	}

	return r.GetWidget(id)
}

func (r Repository) UpdateWidget(id uint, widget write.Widget) error {
	cmd := "update widgets set description = $1, owner = $2, value = $3, updated_at = current_timestamp where id = $4"
	_, err := r.db.Exec(cmd, widget.Description, widget.Owner, widget.Value, id)

	return err
}

func (r Repository) DeleteWidget(id uint) error {
	cmd := "update widgets set deleted_at = current_timestamp where id = $1"
	_, err := r.db.Exec(cmd, id)

	return err
}
