package middleware_test

import (
	"fmt"
	"myapp/app/router/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	expRespBody    = "{\"message\":\"Hello World!\"}"
	expContentType = "application/json"
)

func sampleHandlerFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expRespBody)
	}
}

func TestContentTypeJson(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	middleware.ContentTypeJson(http.HandlerFunc(sampleHandlerFunc())).ServeHTTP(rr, r)
	response := rr.Result()

	if respBody := rr.Body.String(); respBody != expRespBody {
		t.Errorf("Wrong response body:  got %v want %v", respBody, expRespBody)
	}

	if status := response.StatusCode; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}

	if contentType := response.Header.Get("Content-type"); contentType != expContentType {
		t.Errorf("Wrong status code: got %v want %v", contentType, expContentType)
	}
}
