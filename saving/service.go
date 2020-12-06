package saving

type repository interface {
	AddWidget(widget *Widget) error
	UpdateWidget(id int, widget *Widget) error
}

type Service interface {
	AddWidget(widget *Widget) error
	UpdateWidget(id int, widget *Widget) error
}

type service struct {
	r repository
}

func NewService(r repository) Service {
	return &service{r: r}
}

func (s *service) AddWidget(w *Widget) error {
	return s.r.AddWidget(w)
}

func (s *service) UpdateWidget(id int, w *Widget) error {
	return s.r.UpdateWidget(id, w)
}
