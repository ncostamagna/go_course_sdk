package course_test

import (
	"errors"
	"testing"
	"net/http"
	"fmt"
	"os"
	c "github.com/ncostamagna/go_http_client/client"
	courseSdk "github.com/ncostamagna/go_course_sdk/course"
)

var header http.Header
var sdk courseSdk.Transport

func TestMain(m *testing.M) {
	header = http.Header{}
	header.Set("Content-Type", "application/json")
	sdk = courseSdk.NewHttpClient("base-url", "")
	os.Exit(m.Run())
}

// use the mock in the client
func TestGet_ResponseError(t *testing.T) {
	expectedErr := courseSdk.ErrNotFound{
		Message: "course '1' doesn't exist",
	}

	err := c.AddMockups(&c.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/courses/1",
		RespHTTPCode: 404,
		RespBody: fmt.Sprintf(`{
						"status": 404,
						"message": "%s"
					}`, expectedErr.Error()),
	})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	course, err := sdk.Get("1")

	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
	if course != nil {
		t.Errorf("expected nil, got %v", course)
	}
}