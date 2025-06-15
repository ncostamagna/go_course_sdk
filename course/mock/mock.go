package mock

import (
	"errors"

	"github.com/ncostamagna/gocourse_domain/domain"
)

type CourseSdkMock struct {
	GetMock func(id string) (*domain.Course, error)
}

func (r *CourseSdkMock) Get(id string) (*domain.Course, error) {
	if r.GetMock == nil {
		return nil, errors.New("GetMock is not set")
	}
	return r.GetMock(id)
}