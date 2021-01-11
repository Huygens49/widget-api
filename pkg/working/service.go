package working

import (
	"github.com/Huygens49/widget-api/pkg/read"
	"github.com/Huygens49/widget-api/pkg/write"
)

type repository interface {
	GetWidget(id uint) (read.Widget, error)
	UpdateWidget(id uint, widget write.Widget) error
}

type Service interface {
	WorkOnWidget(id uint) error
}

type service struct {
	r repository
}

func NewService(r repository) Service {
	return service{r: r}
}

func (s service) WorkOnWidget(id uint) error {
	rw, err := s.r.GetWidget(id)

	if err != nil {
		return err
	}

	widget := Widget{Value: rw.Value}
	widget.Work()

	sw := write.Widget{
		Description: rw.Description,
		Owner:       rw.Owner,
		Value:       widget.Value,
	}

	return s.r.UpdateWidget(rw.ID, sw)
}
