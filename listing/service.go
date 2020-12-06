package listing

type repository interface {
	GetWidget(id int) (*Widget, error)
	GetAllWidgets() ([]Widget, error)
}

type Service interface {
	GetWidget(id int) (*Widget, error)
	GetAllWidgets() ([]Widget, error)
}

type service struct {
	r repository
}

func NewService(r repository) Service {
	return &service{r: r}
}

func (s *service) GetWidget(id int) (*Widget, error) {
	w, err := s.r.GetWidget(id)
	return w, err
}

func (s *service) GetAllWidgets() ([]Widget, error) {
	w, err := s.r.GetAllWidgets()
	return w, err
}
