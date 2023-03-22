package service

import (
	"context"
	"errors"
	"myapp/model"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	mock_repository "myapp/mocks/repository"
)

var bookForm = &model.BookForm{
	Title:         "title",
	Author:        "author",
	PublishedDate: "2006-01-02",
	ImageUrl:      "image_url",
	Description:   "description",
}

var bookDB = &model.Book{
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

var booksDB = model.Books{
	bookDB,
}

func TestBookService_CreateBook(t *testing.T) {
	type args struct {
		ctx  context.Context
		book *model.BookForm
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		prepareMock func(mockRepo *mock_repository.MockBookRepoInterface)
	}{
		{
			name: "success call",
			args: args{
				ctx:  context.Background(),
				book: bookForm,
			},
			wantErr: false,
			prepareMock: func(mockRepo *mock_repository.MockBookRepoInterface) {
				mockRepo.EXPECT().CreateBook(gomock.Any(), gomock.Any()).Return(bookDB, nil).AnyTimes()
			},
		},
		{
			name: "error from repo",
			args: args{
				ctx:  context.Background(),
				book: bookForm,
			},
			wantErr: true,
			prepareMock: func(mockRepo *mock_repository.MockBookRepoInterface) {
				mockRepo.EXPECT().CreateBook(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
		},
		{
			name: "error parse book",
			args: args{
				ctx:  context.Background(),
				book: &model.BookForm{},
			},
			wantErr: true,
			prepareMock: func(mockRepo *mock_repository.MockBookRepoInterface) {
				mockRepo.EXPECT().CreateBook(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockRepo := mock_repository.NewMockBookRepoInterface(ctrl)

			if tt.prepareMock != nil {
				tt.prepareMock(mockRepo)
			}

			svc := NewBookService(mockRepo)

			resp, err := svc.CreateBook(tt.args.ctx, tt.args.book)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.NotEmpty(t, resp)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestBookService_GetBookByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		prepareMock func(mockRepo *mock_repository.MockBookRepoInterface)
	}{
		{
			name: "success call",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantErr: false,
			prepareMock: func(mockRepo *mock_repository.MockBookRepoInterface) {
				mockRepo.EXPECT().ReadBook(gomock.Any(), gomock.Any()).Return(bookDB, nil).AnyTimes()
			},
		},
		{
			name: "error call",
			args: args{
				ctx: context.Background(),
				id:  2,
			},
			wantErr: true,
			prepareMock: func(mockRepo *mock_repository.MockBookRepoInterface) {
				mockRepo.EXPECT().ReadBook(gomock.Any(), gomock.Any()).Return(nil, errors.New("error")).AnyTimes()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockRepo := mock_repository.NewMockBookRepoInterface(ctrl)

			if tt.prepareMock != nil {
				tt.prepareMock(mockRepo)
			}

			svc := NewBookService(mockRepo)

			resp, err := svc.GetBookByID(tt.args.ctx, tt.args.id)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.NotEmpty(t, resp)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestBookService_GetListBook(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		prepareMock func(mockRepo *mock_repository.MockBookRepoInterface)
	}{
		{
			name: "success call",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
			prepareMock: func(mockRepo *mock_repository.MockBookRepoInterface) {
				mockRepo.EXPECT().ListBooks(gomock.Any()).Return(booksDB, nil).AnyTimes()
			},
		},
		{
			name: "error call",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			prepareMock: func(mockRepo *mock_repository.MockBookRepoInterface) {
				mockRepo.EXPECT().ListBooks(gomock.Any()).Return(nil, errors.New("error")).AnyTimes()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockRepo := mock_repository.NewMockBookRepoInterface(ctrl)

			if tt.prepareMock != nil {
				tt.prepareMock(mockRepo)
			}

			svc := NewBookService(mockRepo)

			resp, err := svc.GetListBook(tt.args.ctx)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.NotEmpty(t, resp)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestBookService_UpdateBook(t *testing.T) {
	type args struct {
		ctx  context.Context
		id   uint
		book *model.BookForm
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		prepareMock func(mockRepo *mock_repository.MockBookRepoInterface)
	}{
		{
			name: "success call",
			args: args{
				ctx:  context.Background(),
				id:   1,
				book: bookForm,
			},
			wantErr: false,
			prepareMock: func(mockRepo *mock_repository.MockBookRepoInterface) {
				mockRepo.EXPECT().UpdateBook(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			},
		},
		{
			name: "error call",
			args: args{
				ctx:  context.Background(),
				id:   1,
				book: bookForm,
			},
			wantErr: true,
			prepareMock: func(mockRepo *mock_repository.MockBookRepoInterface) {
				mockRepo.EXPECT().UpdateBook(gomock.Any(), gomock.Any()).Return(errors.New("error")).AnyTimes()
			},
		},
		{
			name: "error empty book",
			args: args{
				ctx:  context.Background(),
				id:   1,
				book: &model.BookForm{},
			},
			wantErr: true,
			prepareMock: func(mockRepo *mock_repository.MockBookRepoInterface) {
				mockRepo.EXPECT().UpdateBook(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockRepo := mock_repository.NewMockBookRepoInterface(ctrl)

			if tt.prepareMock != nil {
				tt.prepareMock(mockRepo)
			}

			svc := NewBookService(mockRepo)

			err := svc.UpdateBook(tt.args.ctx, tt.args.id, tt.args.book)
			if !tt.wantErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestBookService_DeleteBook(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uint
	}

	tests := []struct {
		name        string
		args        args
		wantErr     bool
		prepareMock func(mockRepo *mock_repository.MockBookRepoInterface)
	}{
		{
			name: "success call",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantErr: false,
			prepareMock: func(mockRepo *mock_repository.MockBookRepoInterface) {
				mockRepo.EXPECT().DeleteBook(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			},
		},
		{
			name: "error call",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantErr: true,
			prepareMock: func(mockRepo *mock_repository.MockBookRepoInterface) {
				mockRepo.EXPECT().DeleteBook(gomock.Any(), gomock.Any()).Return(errors.New("error")).AnyTimes()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockRepo := mock_repository.NewMockBookRepoInterface(ctrl)

			if tt.prepareMock != nil {
				tt.prepareMock(mockRepo)
			}

			svc := NewBookService(mockRepo)

			err := svc.DeleteBook(tt.args.ctx, tt.args.id)
			if !tt.wantErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
