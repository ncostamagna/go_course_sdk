package mock

import (
	"errors"

	"github.com/ncostamagna/gocourse_domain/domain"
)


type CourseSdkMock struct {
	GetFunc func(id string) (*domain.Course, error)
}

func (r *CourseSdkMock) Get(id string) (*domain.Course, error) {
	if r.GetFunc == nil {
		return nil, errors.New("GetFunc is not set")
	}
	return r.GetFunc(id)
}