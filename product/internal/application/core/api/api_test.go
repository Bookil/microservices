package api_test

import (
	"context"
	"errors"
	"testing"

	"product/internal/application/core/api"
	"product/internal/application/core/domain"
	"product/mocks"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type ApplicationTestSuit struct {
	suite.Suite

	api  *api.Application
	cart *mocks.MockCartPort
	DB   *mocks.MockDBPort
}

var ErrUnknownError = errors.New("random error")

func TestApplicationTestSuit(t *testing.T) {
	suite.Run(t, new(ApplicationTestSuit))
}

func (a *ApplicationTestSuit) SetupSuite() {
	ctrl := gomock.NewController(a.T())

	mockedCart := mocks.NewMockCartPort(ctrl)
	mockedDB := mocks.NewMockDBPort(ctrl)

	api := api.NewApplication(mockedCart, mockedDB)

	a.api = api
	a.cart = mockedCart
	a.DB = mockedDB
}

func (a *ApplicationTestSuit) TestGetAllBooks_Success() {
	ctx := context.TODO()
	expectedBooks := []domain.Book{
		{ID: 1, Title: "Book 1"},
		{ID: 2, Title: "Book 2"},
	}

	a.DB.EXPECT().GetAllBooks(ctx).Return(expectedBooks, nil)

	books, err := a.api.GetAllBooks(ctx)
	a.NoError(err)
	a.Equal(expectedBooks, books)
}

func (a *ApplicationTestSuit) TestGetAllBooks_Error() {
	ctx := context.TODO()

	a.DB.EXPECT().GetAllBooks(ctx).Return(nil, ErrUnknownError)

	books, err := a.api.GetAllBooks(ctx)
	a.Error(err)
	a.Nil(books)
}

func (a *ApplicationTestSuit) TestGetBookByID_Success() {
	ctx := context.TODO()
	expectedBook := &domain.Book{ID: 1, Title: "Book 1"}

	a.DB.EXPECT().GetBookByID(ctx, domain.BookID(1)).Return(expectedBook, nil)

	book, err := a.api.GetBookByID(ctx, domain.BookID(1))
	a.NoError(err)
	a.Equal(expectedBook, book)
}

func (a *ApplicationTestSuit) TestGetBookByID_Error() {
	ctx := context.TODO()

	a.DB.EXPECT().GetBookByID(ctx, domain.BookID(1)).Return(nil, ErrUnknownError)

	book, err := a.api.GetBookByID(ctx, domain.BookID(1))
	a.Error(err)
	a.Nil(book)
}

func (a *ApplicationTestSuit) TestGetBooksByTitle_Success() {
	ctx := context.TODO()
	title := "Book 1"
	expectedBooks := []domain.Book{
		{ID: 1, Title: "Book 1"},
	}

	a.DB.EXPECT().GetBooksByTitle(ctx, title).Return(expectedBooks, nil)

	books, err := a.api.GetBooksByTitle(ctx, title)
	a.NoError(err)
	a.Equal(expectedBooks, books)
}

func (a *ApplicationTestSuit) TestGetBooksByTitle_Error() {
	ctx := context.TODO()
	title := "Book 1"

	a.DB.EXPECT().GetBooksByTitle(ctx, title).Return(nil, ErrUnknownError)

	books, err := a.api.GetBooksByTitle(ctx, title)
	a.Error(err)
	a.Nil(books)
}

func (a *ApplicationTestSuit) TestGetBooksByAuthor_Success() {
	ctx := context.TODO()
	authorName := "Author 1"
	expectedBooks := []domain.Book{
		{ID: 1, Title: "Book 1"},
	}

	a.DB.EXPECT().GetBooksByAuthor(ctx, authorName).Return(expectedBooks, nil)

	books, err := a.api.GetBooksByAuthor(ctx, authorName)
	a.NoError(err)
	a.Equal(expectedBooks, books)
}

func (a *ApplicationTestSuit) TestGetBooksByAuthor_Error() {
	ctx := context.TODO()
	authorName := "Author 1"

	a.DB.EXPECT().GetBooksByAuthor(ctx, authorName).Return(nil, ErrUnknownError)

	books, err := a.api.GetBooksByAuthor(ctx, authorName)
	a.Error(err)
	a.Nil(books)
}

func (a *ApplicationTestSuit) TestGetBooksByGenre_Success() {
	ctx := context.TODO()
	genreName := "Genre 1"
	expectedBooks := []domain.Book{
		{ID: 1, Title: "Book 1"},
	}

	a.DB.EXPECT().GetBooksByGenre(ctx, genreName).Return(expectedBooks, nil)

	books, err := a.api.GetBooksByGenre(ctx, genreName)
	a.NoError(err)
	a.Equal(expectedBooks, books)
}

func (a *ApplicationTestSuit) TestGetBooksByGenre_Error() {
	ctx := context.TODO()
	genreName := "Genre 1"

	a.DB.EXPECT().GetBooksByGenre(ctx, genreName).Return(nil, ErrUnknownError)

	books, err := a.api.GetBooksByGenre(ctx, genreName)
	a.Error(err)
	a.Nil(books)
}

func (a *ApplicationTestSuit) TestAddBook_Success() {
	ctx := context.TODO()
	book := &domain.Book{ID: 1, Title: "Book 1"}

	a.DB.EXPECT().AddBook(ctx, book).Return(book, nil)

	newBook, err := a.api.AddBook(ctx, book)
	a.NoError(err)
	a.Equal(book, newBook)
}

func (a *ApplicationTestSuit) TestAddBook_Error() {
	ctx := context.TODO()
	book := &domain.Book{ID: 1, Title: "Book 1"}

	a.DB.EXPECT().AddBook(ctx, book).Return(nil, ErrUnknownError)

	newBook, err := a.api.AddBook(ctx, book)
	a.Error(err)
	a.Nil(newBook)
}

func (a *ApplicationTestSuit) TestModifyBookByID_Success() {
	ctx := context.TODO()
	bookID := domain.BookID(1)
	book := &domain.Book{ID: 1, Title: "Book 1"}

	a.DB.EXPECT().ModifyBookByID(ctx, bookID, book).Return(book, nil)

	modifiedBook, err := a.api.ModifyBookByID(ctx, bookID, book)
	a.NoError(err)
	a.Equal(book, modifiedBook)
}

func (a *ApplicationTestSuit) TestModifyBookByID_Error() {
	ctx := context.TODO()
	bookID := domain.BookID(1)
	book := &domain.Book{ID: 1, Title: "Book 1"}

	a.DB.EXPECT().ModifyBookByID(ctx, bookID, book).Return(nil, ErrUnknownError)

	modifiedBook, err := a.api.ModifyBookByID(ctx, bookID, book)
	a.Error(err)
	a.Nil(modifiedBook)
}

func (a *ApplicationTestSuit) TestDeleteBookByID_Success() {
	ctx := context.TODO()
	bookID := domain.BookID(1)

	a.DB.EXPECT().DeleteBookByID(ctx, bookID).Return(nil)

	err := a.api.DeleteBookByID(ctx, bookID)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestDeleteBookByID_Error() {
	ctx := context.TODO()
	bookID := domain.BookID(1)

	a.DB.EXPECT().DeleteBookByID(ctx, bookID).Return(ErrUnknownError)

	err := a.api.DeleteBookByID(ctx, bookID)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestAddBookToCart_Success() {
	ctx := context.TODO()
	bookID := domain.BookID(1)
	userID := "user1"
	book := &domain.Book{ID: 1, Title: "Book 1"}

	a.DB.EXPECT().GetBookByID(ctx, bookID).Return(book, nil)
	a.cart.EXPECT().AddBookToCart(ctx, bookID, userID).Return(nil)

	err := a.api.AddBookToCart(ctx, bookID, userID)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestAddBookToCart_GetBookError() {
	ctx := context.TODO()
	bookID := domain.BookID(1)
	userID := "user1"

	a.DB.EXPECT().GetBookByID(ctx, bookID).Return(nil, ErrUnknownError)

	err := a.api.AddBookToCart(ctx, bookID, userID)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestAddBookToCart_AddToCartError() {
	ctx := context.TODO()
	bookID := domain.BookID(1)
	userID := "user1"
	book := &domain.Book{ID: 1, Title: "Book 1"}

	a.DB.EXPECT().GetBookByID(ctx, bookID).Return(book, nil)
	a.cart.EXPECT().AddBookToCart(ctx, bookID, userID).Return(ErrUnknownError)

	err := a.api.AddBookToCart(ctx, bookID, userID)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestDeleteBookFromCartByID_Success() {
	ctx := context.TODO()
	bookID := domain.BookID(1)
	cartID := uint(1)

	a.cart.EXPECT().DeleteBookFromCartByID(ctx, bookID, cartID).Return(nil)

	err := a.api.DeleteBookFromCartByID(ctx, bookID, cartID)
	a.NoError(err)
}

func (a *ApplicationTestSuit) TestDeleteBookFromCartByID_Error() {
	ctx := context.TODO()
	bookID := domain.BookID(1)
	cartID := uint(1)

	a.cart.EXPECT().DeleteBookFromCartByID(ctx, bookID, cartID).Return(ErrUnknownError)

	err := a.api.DeleteBookFromCartByID(ctx, bookID, cartID)
	a.Error(err)
}

func (a *ApplicationTestSuit) TestAddAuthor_Success() {
	ctx := context.TODO()
	author := &domain.Author{ID: 1, Name: "Author 1"}

	a.DB.EXPECT().AddAuthor(ctx, author).Return(author, nil)

	newAuthor, err := a.api.AddAuthor(ctx, author)
	a.NoError(err)
	a.Equal(author, newAuthor)
}

func (a *ApplicationTestSuit) TestAddAuthor_Error() {
	ctx := context.TODO()
	author := &domain.Author{ID: 1, Name: "Author 1"}

	a.DB.EXPECT().AddAuthor(ctx, author).Return(nil, ErrUnknownError)

	newAuthor, err := a.api.AddAuthor(ctx, author)
	a.Error(err)
	a.Nil(newAuthor)
}

func (a *ApplicationTestSuit) TestGetAuthorByID_Success() {
	ctx := context.TODO()
	ID := 1
	author := &domain.Author{ID: 1, Name: "Author 1"}

	a.DB.EXPECT().GetAuthorByID(ctx,uint(ID) ).Return(author,nil)

	storedAuthor, err := a.api.GetAuthorByID(ctx, uint(ID))
	a.NoError(err)
	a.Equal(author,storedAuthor)
}


func (a *ApplicationTestSuit) TestGetAuthorByID_Error() {
	ctx := context.TODO()
	ID := 1

	a.DB.EXPECT().GetAuthorByID(ctx,uint(ID)).Return(nil,ErrUnknownError)

	storedAuthor, err := a.api.GetAuthorByID(ctx, uint(ID))	
	a.Error(err)
	a.Nil(storedAuthor)
}

func (a *ApplicationTestSuit) TestDeleteAuthorByID_Success() {
	ctx := context.TODO()
	ID := 1

	a.DB.EXPECT().DeleteAuthorByID(ctx,uint(ID)).Return(nil)

	err := a.api.DeleteAuthorByID(ctx, uint(ID))
	a.NoError(err)
}
func (a *ApplicationTestSuit) TestDeleteAuthorByID_Error() {
	ctx := context.TODO()
	ID := 1

	a.DB.EXPECT().DeleteAuthorByID(ctx,uint(ID)).Return(ErrUnknownError)

	err := a.api.DeleteAuthorByID(ctx, uint(ID))
	a.Error(err)
}

func (a *ApplicationTestSuit) TestGetAllAuthors_Success() {
	ctx := context.TODO()
	expectedAuthors := []domain.Author{
		{ID: 1, Name: "Author 1"},
		{ID: 2, Name: "Author 2"},
	}

	a.DB.EXPECT().GetAllAuthors(ctx).Return(expectedAuthors, nil)

	authors, err := a.api.GetAllAuthors(ctx)
	a.NoError(err)
	a.Equal(expectedAuthors, authors)
}

func (a *ApplicationTestSuit) TestGetAllAuthors_Error() {
	ctx := context.TODO()

	a.DB.EXPECT().GetAllAuthors(ctx).Return(nil, ErrUnknownError)

	authors, err := a.api.GetAllAuthors(ctx)
	a.Error(err)
	a.Nil(authors)
}

func (a *ApplicationTestSuit) TestAddGenre_Success() {
	ctx := context.TODO()
	genre := &domain.Genre{ID: 1, Name: "Genre 1"}

	a.DB.EXPECT().AddGenre(ctx, genre).Return(genre, nil)

	newGenre, err := a.api.AddGenre(ctx, genre)
	a.NoError(err)
	a.Equal(genre, newGenre)
}

func (a *ApplicationTestSuit) TestAddGenre_Error() {
	ctx := context.TODO()
	genre := &domain.Genre{ID: 1, Name: "Genre 1"}

	a.DB.EXPECT().AddGenre(ctx, genre).Return(nil, ErrUnknownError)

	newGenre, err := a.api.AddGenre(ctx, genre)
	a.Error(err)
	a.Nil(newGenre)
}

func (a *ApplicationTestSuit) TestGetGenreByID_Success() {
	ctx := context.TODO()
	ID := 1
	genre := &domain.Genre{ID: 1, Name: "Genre 1"}

	a.DB.EXPECT().GetGenreByID(ctx,uint(ID) ).Return(genre,nil)

	storedGenre, err := a.api.GetGenreByID(ctx, uint(ID))
	a.NoError(err)
	a.Equal(genre,storedGenre)
}


func (a *ApplicationTestSuit) TestGetGenreByID_Error() {
	ctx := context.TODO()
	ID := 1

	a.DB.EXPECT().GetGenreByID(ctx,uint(ID)).Return(nil,ErrUnknownError)

	storedGenre, err := a.api.GetGenreByID(ctx, uint(ID))	
	a.Error(err)
	a.Nil(storedGenre)
}

func (a *ApplicationTestSuit) TestDeleteGenreByID_Success() {
	ctx := context.TODO()
	ID := 1

	a.DB.EXPECT().DeleteGenreByID(ctx,uint(ID)).Return(nil)

	err := a.api.DeleteGenreByID(ctx, uint(ID))
	a.NoError(err)
}
func (a *ApplicationTestSuit) TestDeleteGenreByID_Error() {
	ctx := context.TODO()
	ID := 1

	a.DB.EXPECT().DeleteGenreByID(ctx,uint(ID)).Return(ErrUnknownError)

	err := a.api.DeleteGenreByID(ctx, uint(ID))
	a.Error(err)
}
func (a *ApplicationTestSuit) TestGetAllGenres_Success() {
	ctx := context.TODO()
	expectedGenres := []domain.Genre{
		{ID: 1, Name: "Genre 1"},
		{ID: 2, Name: "Genre 2"},
	}

	a.DB.EXPECT().GetAllGenres(ctx).Return(expectedGenres, nil)

	genres, err := a.api.GetAllGenres(ctx)
	a.NoError(err)
	a.Equal(expectedGenres, genres)
}

func (a *ApplicationTestSuit) TestGetAllGenres_Error() {
	ctx := context.TODO()

	a.DB.EXPECT().GetAllGenres(ctx).Return(nil, ErrUnknownError)

	genres, err := a.api.GetAllGenres(ctx)
	a.Error(err)
	a.Nil(genres)
}
