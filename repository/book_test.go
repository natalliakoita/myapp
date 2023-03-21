package repository_test

import (
	"context"
	"database/sql/driver"
	"errors"
	"log"
	"myapp/model"
	"myapp/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	conn, err := gorm.Open("mysql", db)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return conn, mock
}

var book = &model.Book{
	Model: gorm.Model{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: &time.Time{},
	},
	Title:         "title",
	Author:        "author",
	PublishedDate: time.Now(),
	ImageUrl:      "image_url",
	Description:   "description",
}

func TestBookRepo_ReadBook(t *testing.T) {
	db, mock := NewMock()

	defer db.Close()

	repo := repository.NewBookRepo(db)

	query := "SELECT * FROM `books` WHERE `books`.`deleted_at` IS NULL AND ((id = ?)) ORDER BY `books`.`id` ASC LIMIT 1"

	t.Run("Success call", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "title", "author", "published_date", "image_url", "description"}).
			AddRow(
				book.ID,
				book.CreatedAt,
				book.UpdatedAt,
				book.DeletedAt,
				book.Title,
				book.Author,
				book.PublishedDate,
				book.ImageUrl,
				book.Description)

		mock.ExpectQuery(query).
			WithArgs(book.ID).
			WillReturnRows(rows)

		resp, err := repo.ReadBook(context.Background(), book.ID)
		assert.NoError(t, err)
		assert.NotEmpty(t, resp)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Error call", func(t *testing.T) {
		mock.ExpectQuery(query).
			WillReturnError(errors.New("error"))

		resp, err := repo.ReadBook(context.Background(), book.ID)
		assert.Empty(t, resp)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestBookRepo_ListBook(t *testing.T) {
	db, mock := NewMock()

	defer db.Close()

	repo := repository.NewBookRepo(db)

	query := "SELECT * FROM `books` WHERE `books`.`deleted_at` IS NULL"

	t.Run("Success call", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "title", "author", "published_date", "image_url", "description"}).
			AddRow(
				book.ID,
				book.CreatedAt,
				book.UpdatedAt,
				book.DeletedAt,
				book.Title,
				book.Author,
				book.PublishedDate,
				book.ImageUrl,
				book.Description)

		mock.ExpectQuery(query).
			WillReturnRows(rows)

		resp, err := repo.ListBooks(context.Background())
		assert.NoError(t, err)
		assert.NotEmpty(t, resp)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Error call", func(t *testing.T) {
		mock.ExpectQuery(query).
			WillReturnError(errors.New("error"))

		resp, err := repo.ListBooks(context.Background())
		assert.Empty(t, resp)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestBookRepo_DeleteBook(t *testing.T) {
	db, mock := NewMock()

	defer db.Close()

	repo := repository.NewBookRepo(db)

	query := "UPDATE `books` SET `deleted_at`=? WHERE `books`.`deleted_at` IS NULL AND ((id = ?))"

	t.Run("Success call", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(query).
			WithArgs(
				AnyTime{},
				book.ID,
			).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.DeleteBook(context.Background(), book.ID)
		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Error call", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(query).
			WithArgs(
				AnyTime{},
				book.ID,
			).WillReturnError(errors.New("error"))
		mock.ExpectRollback()

		err := repo.DeleteBook(context.Background(), book.ID)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestBookRepo_CreateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repo := repository.NewBookRepo(gormDB)

	t.Run("Success call", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `books`").WithArgs(
			book.ID,
			book.CreatedAt,
			book.UpdatedAt,
			book.DeletedAt,
			book.Title,
			book.Author,
			book.PublishedDate,
			book.ImageUrl,
			book.Description,
		).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		resp, err := repo.CreateBook(context.Background(), book)
		assert.NoError(t, err)
		assert.NotEmpty(t, resp)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Error call", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `books`").WithArgs(
			book.ID,
			book.CreatedAt,
			book.UpdatedAt,
			book.DeletedAt,
			book.Title,
			book.Author,
			book.PublishedDate,
			book.ImageUrl,
			book.Description,
		).WillReturnError(errors.New("error"))
		mock.ExpectRollback()

		resp, err := repo.CreateBook(context.Background(), book)
		assert.Empty(t, resp)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestBookRepo_UpdateBook(t *testing.T) {
	db, mock := NewMock()

	defer db.Close()

	repo := repository.NewBookRepo(db)

	query := "UPDATE `books` SET `author` = ?, `description` = ?, `image_url` = ?, `published_date` = ?, `title` = ?, `updated_at` = ? WHERE `books`.`deleted_at` IS NULL AND ((id = ?))"

	t.Run("Success call", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(query).WithArgs(
			book.Author,
			book.Description,
			book.ImageUrl,
			book.PublishedDate,
			book.Title,
			AnyTime{},
			book.ID,
		).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.UpdateBook(context.Background(), book)
		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Error call", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(query).WithArgs(
			book.Author,
			book.Description,
			book.ImageUrl,
			book.PublishedDate,
			book.Title,
			AnyTime{},
			book.ID,
		).WillReturnError(errors.New("error"))
		mock.ExpectRollback()

		err := repo.UpdateBook(context.Background(), book)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
