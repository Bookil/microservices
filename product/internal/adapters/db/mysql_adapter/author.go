package mysql_adapter

import (
	"context"

	"product/internal/application/core/domain"
)

func (a *Adapter) AddAuthor(ctx context.Context, author *domain.Author) (*domain.Author, error) {
	err := a.db.WithContext(ctx).Create(author).Error
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (a *Adapter) GetAllAuthors(ctx context.Context) ([]domain.Author, error) {
	var authors []domain.Author

	err := a.db.WithContext(ctx).Preload("Books").Find(&authors).Error
	if err != nil {
		return nil, err
	} else {
		return authors, nil
	}
}
