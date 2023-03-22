package service

import (
	"context"
	"myapp/model"
	"myapp/repository"
)

type BookService struct {
	bookRepo repository.BookRepoInterface
}

func NewBookService(bookRepo repository.BookRepoInterface) *BookService {
	return &BookService{bookRepo: bookRepo}
}

type BookServiceInterface interface {
	CreateBook(ctx context.Context, book *model.BookForm) (*model.BookDto, error)
	GetBookByID(ctx context.Context, id uint) (*model.BookDto, error)
	GetListBook(ctx context.Context) ([]model.BookDto, error)
	UpdateBook(ctx context.Context, id uint, book *model.BookForm) error
	DeleteBook(ctx context.Context, id uint) error
}

func (b *BookService) CreateBook(ctx context.Context, book *model.BookForm) (*model.BookDto, error) {
	bookModel, err := book.ToModel()
	if err != nil {
		return &model.BookDto{}, err
	}

	respBook, err := b.bookRepo.CreateBook(ctx, bookModel)
	if err != nil {
		return &model.BookDto{}, err
	}

	resp := respBook.ToDto()

	return resp, nil
}

func (b *BookService) GetBookByID(ctx context.Context, id uint) (*model.BookDto, error) {
	book, err := b.bookRepo.ReadBook(ctx, id)
	if err != nil {
		return &model.BookDto{}, err
	}

	bookDto := book.ToDto()

	return bookDto, nil
}

func (b *BookService) GetListBook(ctx context.Context) ([]model.BookDto, error) {
	books, err := b.bookRepo.ListBooks(ctx)
	if err != nil {
		return []model.BookDto{}, err
	}

	booksDto := books.ToDto()

	return booksDto, nil
}

func (b *BookService) UpdateBook(ctx context.Context, id uint, book *model.BookForm) error {
	bookModel, err := book.ToModel()
	if err != nil {
		return err
	}

	bookModel.ID = id
	err = b.bookRepo.UpdateBook(ctx, bookModel)
	if err != nil {
		return err
	}

	return nil
}

func (b *BookService) DeleteBook(ctx context.Context, id uint) error {
	err := b.bookRepo.DeleteBook(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
