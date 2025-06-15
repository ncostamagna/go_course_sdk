package mock

import (
	"errors"

	"github.com/ncostamagna/gocourse_domain/domain"
)

type UserSdkMock struct {
	GetMock func(id string) (*domain.User, error)
}

func (r *UserSdkMock) Get(id string) (*domain.User, error) {
	if r.GetMock == nil {
		return nil, errors.New("GetMock is not set")
	}
	return r.GetMock(id)
}