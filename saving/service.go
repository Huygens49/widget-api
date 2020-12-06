package saving

import "github.com/Huygens49/widget-api/listing"

type repository interface {
	AddWidget(widget *Widget) (*listing.Widget, error)
	UpdateWidget(id uint, widget *Widget) error
}

type Service interface {
	AddWidget(widget *Widget) (*listing.Widget, error)
	UpdateWidget(id uint, widget *Widget) error
}

type service struct {
	r repository
}

func NewService(r repository) Service {
	return &service{r: r}
}

func (s *service) AddWidget(w *Widget) (*listing.Widget, error) {
	return s.r.AddWidget(w)
}

func (s *service) UpdateWidget(id uint, w *Widget) error {
	return s.r.UpdateWidget(id, w)
}
