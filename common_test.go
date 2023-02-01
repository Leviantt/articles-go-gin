package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine {
	router := gin.Default()
	if withTemplates {
		router.LoadHTMLGlob("templates/*")
	}
	return router
}

func testHTTPResponse(	test *testing.T, 
						router *gin.Engine, 
						req *http.Request, 
						isEqual func(recorder *httptest.ResponseRecorder) bool) {
	// create a response recorder
	recorder := httptest.NewRecorder()

	// create the service and process the above request
	router.ServeHTTP(recorder, req)

	if !isEqual(recorder) {
		test.Fail()
	}
}

func saveList() {
	tmpArticleList = articleList
}

func restoreList() {
	articleList = tmpArticleList
}