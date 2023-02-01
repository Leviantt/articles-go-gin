package main 

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(test *testing.T) {
	router := getRouter(true)

	router.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(test, router, req, func(recorder *httptest.ResponseRecorder) bool {
		statusOK := recorder.Code == http.StatusOK

		page, err := ioutil.ReadAll(recorder.Body)
		pageOK := err != nil && strings.Index(string(page), "<title>Home Page</title>") > 0
		return statusOK && pageOK
	})
}
