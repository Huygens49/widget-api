package database

import (
	"github.com/Huygens49/widget-api/pkg/reading"
	"github.com/Huygens49/widget-api/pkg/saving"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r Repository) GetAllWidgets() ([]reading.Widget, error) {
	var widgetEntities []WidgetEntity
	result := r.db.Find(&widgetEntities)

	if result.Error != nil {
		return nil, result.Error
	}

	var widgets []reading.Widget
	for _, we := range widgetEntities {
		widget := reading.Widget{
			ID:          we.ID,
			Description: we.Description,
			Owner:       we.Owner,
			Value:       we.Value,
			CreatedAt:   we.CreatedAt,
			UpdatedAt:   we.UpdatedAt,
		}

		widgets = append(widgets, widget)
	}

	return widgets, nil
}

func (r Repository) GetWidget(id uint) (reading.Widget, error) {
	var we WidgetEntity
	result := r.db.First(&we, id)

	if result.Error != nil {
		return reading.Widget{}, result.Error
	}

	widget := reading.Widget{
		ID:          we.ID,
		Description: we.Description,
		Owner:       we.Owner,
		Value:       we.Value,
		CreatedAt:   we.CreatedAt,
		UpdatedAt:   we.UpdatedAt,
	}

	return widget, nil
}

func (r Repository) AddWidget(widget saving.Widget) (reading.Widget, error) {
	we := WidgetEntity{
		Description: widget.Description,
		Owner:       widget.Owner,
		Value:       widget.Value,
	}

	result := r.db.Create(&we)

	if result.Error != nil {
		return reading.Widget{}, result.Error
	}

	return r.GetWidget(we.ID)
}

func (r Repository) UpdateWidget(id uint, widget saving.Widget) error {
	var we WidgetEntity
	result := r.db.First(&we, id)

	if result.Error != nil {
		return result.Error
	}

	ue := WidgetEntity{
		Description: widget.Description,
		Owner:       widget.Owner,
		Value:       widget.Value,
	}

	result = r.db.Model(&we).Updates(ue)

	return result.Error
}
