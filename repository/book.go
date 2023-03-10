package repository

import (
	"context"
	"myapp/model"

	"github.com/jinzhu/gorm"
)

type BookRepo struct {
	repo *gorm.DB
}

func NewBookRepo(conn *gorm.DB) *BookRepo {
	return &BookRepo{
		repo: conn,
	}
}

func (r *BookRepo) ListBooks(ctx context.Context) (model.Books, error) {
	books := make([]*model.Book, 0)
	if err := r.repo.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepo) ReadBook(ctx context.Context, id uint) (*model.Book, error) {
	book := &model.Book{}
	if err := r.repo.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookRepo) DeleteBook(ctx context.Context, id uint) error {
	book := &model.Book{}
	if err := r.repo.Where("id = ?", id).Delete(&book).Error; err != nil {
		return err
	}

	return nil
}

func (r *BookRepo) CreateBook(ctx context.Context, book *model.Book) (*model.Book, error) {
	if err := r.repo.Create(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookRepo) UpdateBook(ctx context.Context, book *model.Book) error {
	if err := r.repo.First(&model.Book{}, book.ID).Update(book).Error; err != nil {
		return err
	}

	return nil
}

type BookRepoInterface interface {
	ListBooks(ctx context.Context) (model.Books, error)
	ReadBook(ctx context.Context, id uint) (*model.Book, error)
	DeleteBook(ctx context.Context, id uint) error
	CreateBook(ctx context.Context, book *model.Book) (*model.Book, error)
	UpdateBook(ctx context.Context, book *model.Book) error
}
