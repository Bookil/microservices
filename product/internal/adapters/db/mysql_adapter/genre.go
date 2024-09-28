package mysql_adapter

import (
	"context"
	"product/internal/application/core/domain"
)

func (a *Adapter) AddGenre(ctx context.Context, genre *domain.Genre) (*domain.Genre, error) {
	err := a.db.WithContext(ctx).Create(genre).Error
	if err != nil {
		return nil, err
	}

	return genre, nil
}

func (a *Adapter) GetAllGenres(ctx context.Context) ([]domain.Genre, error) {
	var genres []domain.Genre

	err := a.db.WithContext(ctx).Preload("Books").Find(&genres).Error
	if err != nil{
		return nil,err
	}

	return genres, nil
}
