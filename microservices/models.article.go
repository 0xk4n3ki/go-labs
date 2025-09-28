package main

import "errors"

type article struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article Body 1"},
	article{ID: 2, Title: "Article 2", Content: "Article Body 2"},
}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}