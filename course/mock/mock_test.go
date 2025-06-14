package mock

import (
	"testing"

	"github.com/ncostamagna/go_course_sdk/course"
)

func TestMock_Courses(t *testing.T) {
	t.Run("test course mock", func(t *testing.T) {
		var _ course.Transport = (*CourseSdkMock)(nil)
	})
}
