package main

type BookList []struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Wikipedia string `json:"wikipedia"`
	Year      string `json:"year"`
	BookURL   string `json:"book_url"`
	Thumbnail string `json:"thumbnail"`
}
