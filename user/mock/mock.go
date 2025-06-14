package mock

import (
	"errors"

	"github.com/ncostamagna/gocourse_domain/domain"
)

type UserSdkMock struct {
	GetFunc func(id string) (*domain.User, error)
}

func (r *UserSdkMock) Get(id string) (*domain.User, error) {
	if r.GetFunc == nil {
		return nil, errors.New("GetFunc is not set")
	}
	return r.GetFunc(id)
}
