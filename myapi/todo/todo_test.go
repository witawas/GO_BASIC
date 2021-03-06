package todo

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTodosHandler(t *testing.T) {
	t.Run("GET /todos", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/todos", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(todosHandler)

		handler.ServeHTTP(rr, req)

		status := rr.Code
		if http.StatusOK != status {
			t.Errorf("should response status code %v but got %v \n", http.StatusOK, status)
		}

		resp := `hello GET todos`

		if rr.Body.String() != resp {
			t.Errorf("should response body %q but got %q \n", resp, rr.Body.String())
		}
	})
}
