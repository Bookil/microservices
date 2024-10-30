package grpc

import (
	"context"

	"product/internal/adapters/grpc/interceptor"
	"product/internal/application/core/domain"

	productv1 "github.com/Bookil/Bookil-Proto/gen/golang/product/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (a *Adapter) AddAuthor(ctx context.Context, request *productv1.AddAuthorRequest) (*productv1.AddAuthorResponse, error) {
	err := a.validator.ValidateAddAuthor(request.Name, request.About)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	author := domain.NewAuthor(request.Name, request.About)

	_, err = a.api.AddAuthor(ctx, author)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.AddAuthorResponse{}, nil
}

func (a *Adapter) DeleteAuthorByID(ctx context.Context, request *productv1.DeleteAuthorByIDRequest) (*productv1.DeleteAuthorByIDResponse, error) {
	err := a.api.DeleteAuthorByID(ctx, uint(request.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.DeleteAuthorByIDResponse{}, nil
}

func (a *Adapter) GetAllAuthors(request *productv1.GetAllAuthorsRequest, stream productv1.ProductService_GetAllAuthorsServer) error {
	authors, err := a.api.GetAllAuthors(stream.Context())
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	for _, author := range authors {

		response := &productv1.GetAllAuthorsResponse{
			Author: &productv1.Author{
				AuthorId: uint32(author.ID),
				Name:     author.Name,
				About:    author.About,
			},
		}

		err := stream.Send(response)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Adapter) AddGenre(ctx context.Context, request *productv1.AddGenreRequest) (*productv1.AddGenreResponse, error) {
	err := a.validator.ValidateAddGenre(request.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	genre := domain.NewGenre(request.Name)

	_, err = a.api.AddGenre(ctx, genre)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.AddGenreResponse{}, nil
}

func (a *Adapter) DeleteGenreByID(ctx context.Context, request *productv1.DeleteGenreByIDRequest) (*productv1.DeleteGenreByIDResponse, error) {
	err := a.api.DeleteGenreByID(ctx, uint(request.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.DeleteGenreByIDResponse{}, nil
}

func (a *Adapter) GetAllGenres(request *productv1.GetAllGenresRequest, stream productv1.ProductService_GetAllGenresServer) error {
	genres, err := a.api.GetAllGenres(stream.Context())
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	for _, genre := range genres {
		response := &productv1.GetAllGenresResponse{
			Genre: &productv1.Genre{
				GenreId: uint32(genre.ID),
				Name:    genre.Name,
			},
		}

		err := stream.Send(response)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Adapter) GetAllBooks(request *productv1.GetAllBooksRequest, stream productv1.ProductService_GetAllBooksServer) error {
	books, err := a.api.GetAllBooks(stream.Context())
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	for _, book := range books {
		var respAuthors []*productv1.Author
		var respGenres []*productv1.Genre

		for _, author := range book.Authors {
			respAuthors = append(respAuthors, &productv1.Author{
				AuthorId: uint32(author.ID),
				Name:     author.Name,
				About:    author.About,
			})
		}

		for _, genre := range book.Genres {
			respGenres = append(respGenres, &productv1.Genre{
				GenreId: uint32(genre.ID),
				Name:    genre.Name,
			})
		}

		response := &productv1.GetAllBooksResponse{
			Book: &productv1.Book{
				BookId:      uint32(book.ID),
				Title:       book.Title,
				Price:       float32(book.Price),
				Description: book.Description,
				Quantity:    uint32(book.Quantity),
				Year:        uint32(book.Year),
				Authors:     respAuthors,
				Genre:       respGenres,
				// Check
				CreatedAt: timestamppb.New(book.CreatedAt),
				UpdatedAt: timestamppb.New(book.UpdatedAt),
			},
		}

		err := stream.Send(response)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Adapter) AddBook(ctx context.Context, request *productv1.AddBookRequest) (*productv1.AddBookResponse, error) {
	err := a.validator.ValidateAddBook(request.Title, request.Description, request.Price, uint(request.Quantity), uint(request.Year))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	authors := convertProtoAuthorsToDomainAuthors(request.Authors)
	genres := convertProtoGenresToDomainGenres(request.Genre)

	book := domain.NewBook(request.Title, request.Description, authors, uint(request.Quantity), uint(request.Year), genres, float64(request.Price))

	_, err = a.api.AddBook(ctx, book)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.AddBookResponse{}, nil
}

func (a *Adapter) GetBookByID(ctx context.Context, request *productv1.GetBookByIDRequest) (*productv1.GetBookByIDResponse, error) {
	book, err := a.api.GetBookByID(ctx, uint(request.BookId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	respBook := convertDomainBookToProtoBook(book)

	return &productv1.GetBookByIDResponse{
		Book: respBook,
	}, nil
}

func (a *Adapter) GetBooksByTitle(request *productv1.GetBooksByTitleRequest, stream productv1.ProductService_GetBooksByTitleServer) error {
	err := a.validator.ValidateGetBooksByTitle(request.Title)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	books, err := a.api.GetBooksByTitle(stream.Context(), request.Title)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	for _, book := range books {
		respBook := convertDomainBookToProtoBook(&book)

		err := stream.Send(&productv1.GetBooksByTitleResponse{Book: respBook})
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

func (a *Adapter) GetBooksByGenre(request *productv1.GetBooksByGenreRequest, stream productv1.ProductService_GetBooksByGenreServer) error {
	err := a.validator.ValidateGetBooksByGenre(request.GenreName)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	books, err := a.api.GetBooksByGenre(stream.Context(), request.GenreName)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	for _, book := range books {
		respBook := convertDomainBookToProtoBook(&book)

		err := stream.Send(&productv1.GetBooksByGenreResponse{Book: respBook})
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

func (a *Adapter) GetBooksByAuthor(request *productv1.GetBooksByAuthorRequest, stream productv1.ProductService_GetBooksByAuthorServer) error {
	err := a.validator.ValidateGetBooksByAuthor(request.AuthorName)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	books, err := a.api.GetBooksByAuthor(stream.Context(), request.AuthorName)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	for _, book := range books {
		respBook := convertDomainBookToProtoBook(&book)

		err := stream.Send(&productv1.GetBooksByAuthorResponse{Book: respBook})
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

func (a *Adapter) DeleteBookByID(ctx context.Context, request *productv1.DeleteBookByIDRequest) (*productv1.DeleteBookByIDResponse, error) {
	err := a.api.DeleteBookByID(ctx, uint(request.BookId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.DeleteBookByIDResponse{}, nil
}

func (a *Adapter) ModifyBookByID(ctx context.Context, request *productv1.ModifyBookByIDRequest) (*productv1.ModifyBookByIDResponse, error) {
	err := a.validator.ValidateModifyBookByID(request.Title, request.Description, request.Price, uint(request.Quantity), uint(request.Year))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	authors := convertProtoAuthorsToDomainAuthors(request.Authors)
	genres := convertProtoGenresToDomainGenres(request.Genre)
	book := domain.NewBook(request.Title, request.Description, authors, uint(request.Quantity), uint(request.Year), genres, float64(request.Price))

	_, err = a.api.ModifyBookByID(ctx, uint(request.BookId), book)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.ModifyBookByIDResponse{}, nil
}

func (a *Adapter) AddBookToCart(ctx context.Context, request *productv1.AddBookToCartRequest) (*productv1.AddBookToCartResponse, error) {
	userID, ok := ctx.Value(interceptor.UserID{}).(string)
	if !ok {
		return nil, status.Error(codes.Internal, "an error occurred")
	}
	err := a.api.AddBookToCart(ctx, uint(request.BookId), userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.AddBookToCartResponse{}, nil
}

func (a *Adapter) DeleteBookFromCartByID(ctx context.Context, request *productv1.DeleteBookFromCartByIDRequest) (*productv1.DeleteBookFromCartByIDResponse, error) {
	err := a.api.DeleteBookFromCartByID(ctx, uint(request.BookId), uint(request.CartId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.DeleteBookFromCartByIDResponse{}, nil
}

func convertDomainBookToProtoBook(book *domain.Book) *productv1.Book {
	var respAuthors []*productv1.Author
	var respGenres []*productv1.Genre

	for _, author := range book.Authors {
		respAuthors = append(respAuthors, &productv1.Author{
			AuthorId: uint32(author.ID),
			Name:     author.Name,
			About:    author.About,
		})
	}

	for _, genre := range book.Genres {
		respGenres = append(respGenres, &productv1.Genre{
			GenreId: uint32(genre.ID),
			Name:    genre.Name,
		})
	}

	resBook := &productv1.Book{
		BookId:      uint32(book.ID),
		Title:       book.Title,
		Description: book.Title,
		Year:        uint32(book.Year),
		Quantity:    uint32(book.Quantity),
		Price:       float32(book.Price),
		Authors:     respAuthors,
		Genre:       respGenres,
		// Check
		CreatedAt: timestamppb.New(book.CreatedAt),
		UpdatedAt: timestamppb.New(book.UpdatedAt),
	}

	return resBook
}

func convertProtoAuthorsToDomainAuthors(inputAuthors []*productv1.Author) []*domain.Author {
	var authors []*domain.Author

	for _, author := range inputAuthors {
		authors = append(authors, &domain.Author{
			ID:    uint(author.AuthorId),
			Name:  author.Name,
			About: author.About,
		})
	}

	return authors
}

func convertProtoGenresToDomainGenres(inputGenres []*productv1.Genre) []*domain.Genre {
	var genres []*domain.Genre

	for _, genre := range inputGenres {
		genres = append(genres, &domain.Genre{
			ID:   uint(genre.GetGenreId()),
			Name: genre.Name,
		})
	}

	return genres
}
