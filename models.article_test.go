package main

import (
	"testing"
)

func testGetAllArticles(test *testing.T) {
	testArticles := getAllArticles()

	if len(testArticles) != len(articleList) {
		test.Fail()
	}

	for i, article := range testArticles {
		if 	article.Content != articleList[i].Content || 
			article.Title != articleList[i].Title || 
			article.ID != articleList[i].ID {
				test.Fail()
				break
		}
	}
}