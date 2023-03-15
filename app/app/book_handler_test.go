package app

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_service "myapp/mocks/service"
	mock_logger "myapp/mocks/util/logger"
	"myapp/model"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestApp_HandleCreateBook(t *testing.T) {
	type args struct {
		jsonStr []byte
		book    *model.BookDto
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		statusCode  int
		prepareMock func(mockSvc *mock_service.MockBookServiceInterface)
	}{
		{
			name: "success call",
			args: args{
				jsonStr: []byte(`{"title":"title", "author":"author", "published_date":"2006-01-02", "image_url":"image_url", "description":"description"}`),
				book:    mockBookDto(),
			},
			wantErr:    false,
			statusCode: http.StatusCreated,
			prepareMock: func(mockSvc *mock_service.MockBookServiceInterface) {
				mockSvc.EXPECT().CreateBook(gomock.Any(), gomock.Any()).Return(mockBookDto(), nil).AnyTimes()
			},
		},
		{
			name: "server error",
			args: args{
				jsonStr: []byte(`{"title":"title", "author":"author", "published_date":"2006-01-02", "image_url":"image_url", "description":"description"}`),
			},
			wantErr:    true,
			statusCode: http.StatusInternalServerError,
			prepareMock: func(mockSvc *mock_service.MockBookServiceInterface) {
				mockSvc.EXPECT().CreateBook(gomock.Any(), gomock.Any()).Return(nil, errors.New("data creation failure")).AnyTimes()
			},
		},
		{
			name: "bad request",
			args: args{
				jsonStr: []byte{},
			},
			wantErr:    true,
			statusCode: http.StatusUnprocessableEntity,
			prepareMock: func(mockSvc *mock_service.MockBookServiceInterface) {
				mockSvc.EXPECT().CreateBook(gomock.Any(), gomock.Any()).Return(nil, errors.New("data access failure")).AnyTimes()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockLogger := mock_logger.NewMockLoggerInterface(ctrl)
			mockLogger.EXPECT().Info().AnyTimes()
			mockLogger.EXPECT().Warn().AnyTimes()

			mockBookService := mock_service.NewMockBookServiceInterface(ctrl)

			if tt.prepareMock != nil {
				tt.prepareMock(mockBookService)
			}

			req, err := http.NewRequest("POST", "api/v1/books", bytes.NewBuffer(tt.args.jsonStr))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			a := &App{
				logger:  mockLogger,
				svcBook: mockBookService,
			}

			handler := http.HandlerFunc(a.HandleCreateBook)
			handler.ServeHTTP(rr, req)

			switch tt.statusCode {
			case http.StatusCreated:
				assert.Equal(t, rr.Code, tt.statusCode)

				b, err := getBody(tt.args.book)
				assert.NoError(t, err)

				str1 := bytes.NewBuffer(b).String()
				str2 := bytes.NewBuffer(rr.Body.Bytes()).String()

				assert.Contains(t, str2, str1)
			case http.StatusUnprocessableEntity:
				assert.Equal(t, rr.Code, tt.statusCode)
			case http.StatusInternalServerError:
				assert.Equal(t, rr.Code, tt.statusCode)
			default:
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Result().StatusCode, tt.statusCode)
			}
		})
	}
}

func mockBookDto() *model.BookDto {
	book := &model.BookDto{
		ID:            1,
		Title:         "title",
		Author:        "author",
		PublishedDate: "2006-01-02",
		ImageUrl:      "image_url",
		Description:   "description",
	}

	return book
}

func TestApp_HandleReadBook(t *testing.T) {
	type args struct {
		id   string
		book *model.BookDto
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		statusCode  int
		prepareMock func(mockSvc *mock_service.MockBookServiceInterface)
	}{
		{
			name: "success call",
			args: args{
				id:   "1",
				book: mockBookDto(),
			},
			wantErr:    false,
			statusCode: http.StatusOK,
			prepareMock: func(mockSvc *mock_service.MockBookServiceInterface) {
				mockSvc.EXPECT().GetBookByID(gomock.Any(), gomock.Any()).Return(mockBookDto(), nil).AnyTimes()

			},
		},
		{
			name: "bad request",
			args: args{
				id: "2",
			},
			wantErr:    true,
			statusCode: http.StatusNotFound,
			prepareMock: func(mockSvc *mock_service.MockBookServiceInterface) {
				mockSvc.EXPECT().GetBookByID(gomock.Any(), gomock.Any()).Return(nil, gorm.ErrRecordNotFound).AnyTimes()
			},
		},
		{
			name: "id request failure",
			args: args{
				id: "invalidParam",
			},
			wantErr:    true,
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "server error",
			args: args{
				id: "1",
			},
			wantErr:    true,
			statusCode: http.StatusInternalServerError,
			prepareMock: func(mockSvc *mock_service.MockBookServiceInterface) {
				mockSvc.EXPECT().GetBookByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("data access failure")).AnyTimes()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockLogger := mock_logger.NewMockLoggerInterface(ctrl)
			mockLogger.EXPECT().Info().AnyTimes()
			mockLogger.EXPECT().Warn().AnyTimes()

			mockBookService := mock_service.NewMockBookServiceInterface(ctrl)

			if tt.prepareMock != nil {
				tt.prepareMock(mockBookService)
			}

			path := fmt.Sprintf("api/v1/books/%s", tt.args.id)
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			a := &App{
				logger:  mockLogger,
				svcBook: mockBookService,
			}

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.args.id)

			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			handler := http.HandlerFunc(a.HandleReadBook)
			handler.ServeHTTP(rr, req)

			switch tt.statusCode {
			case http.StatusOK:
				assert.Equal(t, rr.Code, tt.statusCode)

				b, err := getBody(tt.args.book)
				assert.NoError(t, err)

				str1 := bytes.NewBuffer(b).String()
				str2 := bytes.NewBuffer(rr.Body.Bytes()).String()

				assert.Contains(t, str2, str1)
			case http.StatusNotFound:
				assert.Equal(t, rr.Code, tt.statusCode)
			case http.StatusUnprocessableEntity:
				assert.Equal(t, rr.Code, tt.statusCode)
			case http.StatusInternalServerError:
				assert.Equal(t, rr.Code, tt.statusCode)
			}
		})
	}
}

func getBody(b *model.BookDto) ([]byte, error) {
	result := &model.BookDto{
		ID:            b.ID,
		Title:         b.Title,
		Author:        b.Author,
		PublishedDate: b.PublishedDate,
		ImageUrl:      b.ImageUrl,
		Description:   b.Description,
	}

	book, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func mockListBookDto() []model.BookDto {
	books := []model.BookDto{
		{
			ID:            1,
			Title:         "title",
			Author:        "author",
			PublishedDate: "2006-01-02",
			ImageUrl:      "image_url",
			Description:   "description",
		},
	}

	return books
}

func getListBody(b []model.BookDto) ([]byte, error) {

	book, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func TestApp_ListBooks(t *testing.T) {
	type args struct {
		books []model.BookDto
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		statusCode  int
		prepareMock func(mockSvc *mock_service.MockBookServiceInterface)
	}{
		{
			name: "success call",
			args: args{
				books: mockListBookDto(),
			},
			wantErr:    false,
			statusCode: http.StatusOK,
			prepareMock: func(mockSvc *mock_service.MockBookServiceInterface) {
				mockSvc.EXPECT().GetListBook(gomock.Any()).Return(mockListBookDto(), nil).AnyTimes()
			},
		},
		{
			name:       "server error",
			wantErr:    true,
			statusCode: http.StatusInternalServerError,
			prepareMock: func(mockSvc *mock_service.MockBookServiceInterface) {
				mockSvc.EXPECT().GetListBook(gomock.Any()).Return(nil, errors.New("data access failure")).AnyTimes()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockLogger := mock_logger.NewMockLoggerInterface(ctrl)
			mockLogger.EXPECT().Info().AnyTimes()
			mockLogger.EXPECT().Warn().AnyTimes()

			mockBookService := mock_service.NewMockBookServiceInterface(ctrl)

			if tt.prepareMock != nil {
				tt.prepareMock(mockBookService)
			}

			req, err := http.NewRequest("GET", "api/v1/books", nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			a := &App{
				logger:  mockLogger,
				svcBook: mockBookService,
			}

			handler := http.HandlerFunc(a.HandleListBooks)
			handler.ServeHTTP(rr, req)

			switch tt.statusCode {
			case http.StatusOK:
				assert.Equal(t, rr.Code, tt.statusCode)

				b, err := getListBody(tt.args.books)
				assert.NoError(t, err)

				str1 := bytes.NewBuffer(b).String()
				str2 := bytes.NewBuffer(rr.Body.Bytes()).String()

				assert.Contains(t, str2, str1)
			case http.StatusNotFound:
				assert.Equal(t, rr.Code, tt.statusCode)
			// case http.StatusUnprocessableEntity:
			// 	assert.Equal(t, rr.Code, tt.statusCode)
			case http.StatusInternalServerError:
				assert.Equal(t, rr.Code, tt.statusCode)
			}
		})
	}
}
