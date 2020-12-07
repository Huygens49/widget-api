package database

import (
	"gorm.io/gorm"

	"github.com/Huygens49/widget-api/reading"
	"github.com/Huygens49/widget-api/saving"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllWidgets() ([]reading.Widget, error) {
	var widgetEntities []WidgetEntity
	result := r.db.Find(&widgetEntities)

	if result.Error != nil {
		return nil, result.Error
	}

	widgets := make([]reading.Widget, result.RowsAffected)
	for i, we := range widgetEntities {
		widget := reading.Widget{
			ID:          we.ID,
			Description: we.Description,
			Owner:       we.Owner,
			Value:       we.Value,
			CreatedAt:   we.CreatedAt,
			UpdatedAt:   we.UpdatedAt,
			DeletedAt:   we.DeletedAt.Time,
		}

		widgets[i] = widget
	}

	return widgets, nil
}

func (r *Repository) GetWidget(id uint) (*reading.Widget, error) {
	var we WidgetEntity
	result := r.db.First(&we, id)

	if result.Error != nil {
		return nil, result.Error
	}

	widget := reading.Widget{
		ID:          we.ID,
		Description: we.Description,
		Owner:       we.Owner,
		Value:       we.Value,
		CreatedAt:   we.CreatedAt,
		UpdatedAt:   we.UpdatedAt,
		DeletedAt:   we.DeletedAt.Time,
	}

	return &widget, nil
}

func (r *Repository) AddWidget(widget *saving.Widget) (*reading.Widget, error) {
	we := &WidgetEntity{
		Description: widget.Description,
		Owner:       widget.Owner,
	}

	result := r.db.Create(we)

	if result.Error != nil {
		return nil, result.Error
	}

	return r.GetWidget(we.ID)
}

func (r *Repository) UpdateWidget(id uint, widget *saving.Widget) error {
	var we WidgetEntity
	result := r.db.First(&we, id)

	if result.Error != nil {
		return result.Error
	}

	result = r.db.Model(&we).Updates(WidgetEntity{Description: widget.Description, Owner: widget.Owner, Value: widget.Value})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
