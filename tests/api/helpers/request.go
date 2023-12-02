package helpers

import (
	http "net/http"
	httptest "net/http/httptest"
)

func Request(handler http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	return response
}
