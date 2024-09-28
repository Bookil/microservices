package integration

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"product/config"
	"product/internal/adapters/db"
	adapter "product/internal/adapters/db/mysql_adapter"
	"product/internal/application/core/domain"

	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var genres = []*domain.Genre{
	{ID: 1, Name: "Fiction"},
	{ID: 2, Name: "Non-Fiction"},
	{ID: 3, Name: "Science Fiction"},
	{ID: 4, Name: "Fantasy"},
}

var authors = []*domain.Author{
	{ID: 1, Name: "author 1", About: "this is author 1"},
	{ID: 2, Name: "author 2", About: "this is author 2"},
	{ID: 3, Name: "author 3", About: "this is author 3"},
}

type ProductDatabaseTestSuite struct {
	suite.Suite
	adapter *adapter.Adapter
	book    *domain.Book
}

func TestProductDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(ProductDatabaseTestSuite))
}

func (o *ProductDatabaseTestSuite) SetupSuite() {
	ctx := context.TODO()

	err := os.Setenv("PRODUCT_ENV", "test")
	if err != nil {
		log.Fatalf("Could not set the environment variable to test: %s", err)
	}

	mysqlConfig := &config.Read().Mysql

	port := fmt.Sprintf("%d/tcp", mysqlConfig.Port)

	req := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{port},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": mysqlConfig.Password,
			"MYSQL_DATABASE":      mysqlConfig.DBName,
		},
		WaitingFor: wait.ForListeningPort(nat.Port(port)).WithStartupTimeout(3 * time.Minute),
	}

	mysqlContainer, connectErr := testcontainers.GenericContainer(ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})

	if connectErr != nil {
		log.Fatal("Failed to start Mysql:", connectErr)
	}

	endpoint, err := mysqlContainer.Endpoint(ctx, "")
	if err != nil {
		log.Fatal(err)
	}

	endPort, err := strconv.Atoi(strings.Split(endpoint, ":")[1])
	if err != nil {
		log.Fatal(err)
	}

	mysqlConfig.Port = endPort

	db, err := db.NewDB(mysqlConfig)
	if err != nil {
		log.Fatal(err)
	}

	adapter := adapter.NewAdapter(db)

	for _, genre := range genres {
		_, err := adapter.AddGenre(ctx, genre)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, author := range authors {
		_, err := adapter.AddAuthor(ctx, author)
		if err != nil {
			log.Fatal(err)
		}
	}

	o.adapter = adapter
}

func (o *ProductDatabaseTestSuite) TestA_CreateBook() {
	ctx := context.TODO()

	testCases := []struct {
		book  *domain.Book
		Valid bool
	}{
		{
			book:  domain.NewBook("Book 5", "Bad Book For You", []*domain.Author{authors[0]}, 5, 2024, genres[:2], 3600),
			Valid: true,
		},
		{
			book:  domain.NewBook("Book 6", "Bad Book For You", authors[:2], 5, 2024, genres[:2], 3600),
			Valid: true,
		},
		{
			book:  domain.NewBook("Book 6", "Bad Book For You", []*domain.Author{authors[1]}, 5, 2024, []*domain.Genre{genres[1]}, 3600),
			Valid: true,
		},
	}

	for _, tc := range testCases {
		book, err := o.adapter.AddBook(ctx, tc.book)
		if tc.Valid {
			o.NoError(err)
			o.book = book
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(book)
		}
	}
}

func (o *ProductDatabaseTestSuite) TestB_GetAllBooks() {
	ctx := context.TODO()

	books, err := o.adapter.GetAllBooks(ctx)
	o.NoError(err)
	o.NotNil(books)
}

func (o *ProductDatabaseTestSuite) TestC_GetBooksByTitle() {
	ctx := context.TODO()

	testCases := []struct {
		Title string
		Valid bool
	}{
		{
			Title: o.book.Title,
			Valid: true,
		},
		{
			Title: "Invalid",
			Valid: false,
		},
	}

	for _, tc := range testCases {
		books, err := o.adapter.GetBooksByTitle(ctx, tc.Title)
		if tc.Valid {
			o.NoError(err)
			o.NotNil(books)
			for _, book := range books {
				o.NotNil(book.Genres)
			}
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(books)
		}
	}
}

func (o *ProductDatabaseTestSuite) TestD_GetBooksByID() {
	ctx := context.TODO()

	testCases := []struct {
		ID    uint
		Valid bool
	}{
		{
			ID:    o.book.ID,
			Valid: true,
		},
		{
			ID:    55,
			Valid: false,
		},
	}

	for _, tc := range testCases {
		book, err := o.adapter.GetBookByID(ctx, tc.ID)
		if tc.Valid {
			o.NoError(err)
			o.NotNil(book)
			o.NotNil(book.Genres)
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(book)
		}
	}
}

func (o *ProductDatabaseTestSuite) TestE_GetBooksByGenreName() {
	ctx := context.TODO()

	testCases := []struct {
		GenreName string
		Valid     bool
	}{
		{
			GenreName: genres[0].Name,
			Valid:     true,
		},
		{
			GenreName: "Invalid",
			Valid:     false,
		},
	}

	for _, tc := range testCases {
		books, err := o.adapter.GetBooksByGenre(ctx, tc.GenreName)
		if tc.Valid {
			o.NoError(err)
			o.NotNil(books)
			isGenreIncludes := func() bool {
				for _, book := range books {
					for _, genre := range book.Genres {
						if genre.Name == tc.GenreName {
							return true
						}
					}
				}

				return false
			}()

			o.True(isGenreIncludes)
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(books)
		}
	}
}

func (o *ProductDatabaseTestSuite) TestF_GetBooksByAuthorName() {
	ctx := context.TODO()

	testCases := []struct {
		AuthorName string
		Valid      bool
	}{
		{
			AuthorName: authors[0].Name,
			Valid:      true,
		},
		{
			AuthorName: authors[1].Name,
			Valid:      true,
		},
		{
			AuthorName: "Invalid",
			Valid:      false,
		},
	}

	for _, tc := range testCases {
		books, err := o.adapter.GetBooksByAuthor(ctx, tc.AuthorName)
		if tc.Valid {
			o.NoError(err)
			o.NotNil(books)
			isAuthorIncludes := func() bool {
				for _, book := range books {
					for _, author := range book.Authors {
						if author.Name == tc.AuthorName {
							return true
						}
					}
				}

				return false
			}()

			o.True(isAuthorIncludes)
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(books)
		}
	}
}

func (o *ProductDatabaseTestSuite) TestG_ModifyBookByID() {
	ctx := context.TODO()

	testCases := []struct {
		bookID domain.BookID
		book   *domain.Book
		Valid  bool
	}{
		{
			bookID: o.book.ID,
			book:   domain.NewBook("book 7", "my fav book", []*domain.Author{authors[0]}, 15, 1917, genres, 35),
			Valid:  true,
		},

		{
			bookID: 55,
			book:   domain.NewBook("book 7", "my fav book", []*domain.Author{authors[1]}, 15, 1917, genres, 35),
			Valid:  false,
		},
	}

	for _, tc := range testCases {
		savedBook, err := o.adapter.ModifyBookByID(ctx, tc.bookID, tc.book)
		if tc.Valid {
			o.NoError(err)
			o.NotNil(savedBook)
			o.NotEmpty(savedBook.UpdatedAt)
		} else if !tc.Valid {
			o.Error(err)
		}
	}
}

func (o *ProductDatabaseTestSuite) TestH_DeleteBookByID() {
	ctx := context.TODO()

	testCases := []struct {
		bookID domain.BookID
		Valid  bool
	}{
		{
			bookID: o.book.ID,
			Valid:  true,
		},

		{
			bookID: 55,
			Valid:  false,
		},
	}

	for _, tc := range testCases {
		err := o.adapter.DeleteBookByID(ctx, tc.bookID)
		if tc.Valid {
			o.NoError(err)
		} else if !tc.Valid {
			o.Error(err)
		}
	}
}

func (o *ProductDatabaseTestSuite) TestI_GetAllGenres() {
	ctx := context.TODO()

	genres, err := o.adapter.GetAllGenres(ctx)
	o.NoError(err)
	o.NotNil(genres)
}

func (o *ProductDatabaseTestSuite) TestJ_GetAllAuthors() {
	ctx := context.TODO()

	authors, err := o.adapter.GetAllAuthors(ctx)
	o.NoError(err)
	o.NotNil(authors)
}
