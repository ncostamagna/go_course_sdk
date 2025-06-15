package mock

import (
	"testing"

	"github.com/ncostamagna/go_course_sdk/user"
)

func TestMock_User(t *testing.T) {

	t.Run("test user mock", func(t *testing.T) {
		var _ user.Transport = (*UserSdkMock)(nil)
	})
}