package reading

type repository interface {
	GetWidget(id uint) (*Widget, error)
	GetAllWidgets() ([]Widget, error)
}

type Service interface {
	GetWidget(id uint) (*Widget, error)
	GetAllWidgets() ([]Widget, error)
}

type service struct {
	r repository
}

func NewService(r repository) Service {
	return &service{r: r}
}

func (s *service) GetWidget(id uint) (*Widget, error) {
	w, err := s.r.GetWidget(id)
	return w, err
}

func (s *service) GetAllWidgets() ([]Widget, error) {
	w, err := s.r.GetAllWidgets()
	return w, err
}
