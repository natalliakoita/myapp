package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Books []*Book

func (b Books) ToDto() []BookDto {
	books := make([]BookDto, 0)

	for i := 0; i < len(b); i++ {
		book := BookDto{
			ID:            b[i].ID,
			Title:         b[i].Title,
			Author:        b[i].Author,
			PublishedDate: b[i].PublishedDate.Format("2006-01-02"),
			ImageUrl:      b[i].ImageUrl,
			Description:   b[i].Description,
		}

		books = append(books, book)
	}

	return books
}

type Book struct {
	gorm.Model
	Title         string
	Author        string
	PublishedDate time.Time
	ImageUrl      string
	Description   string
}

type BookDto struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
	ImageUrl      string `json:"image_url"`
	Description   string `json:"description"`
}

func (b Book) ToDto() *BookDto {
	return &BookDto{
		ID:            b.ID,
		Title:         b.Title,
		Author:        b.Author,
		PublishedDate: b.PublishedDate.Format("2006-01-02"),
		ImageUrl:      b.ImageUrl,
		Description:   b.Description,
	}
}

type BookForm struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
	ImageUrl      string `json:"image_url"`
	Description   string `json:"description"`
}

func (f *BookForm) ToModel() (*Book, error) {
	pubDate, err := time.Parse("2006-01-02", f.PublishedDate)
	if err != nil {
		return nil, err
	}

	return &Book{
		Title:         f.Title,
		Author:        f.Author,
		PublishedDate: pubDate,
		ImageUrl:      f.ImageUrl,
		Description:   f.Description,
	}, nil
}
