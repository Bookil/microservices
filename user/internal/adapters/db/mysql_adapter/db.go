package mysql_adapter

import (
	"context"

	"github.com/Bookil/microservices/user/config"
	"github.com/Bookil/microservices/user/internal/adapters/db"
	"github.com/Bookil/microservices/user/internal/application/core/domain"
	"gorm.io/gorm"
)

type (
	Adapter struct {
		db *gorm.DB
	}
)

func NewAdapter(config *config.Mysql) (*Adapter, error) {
	db, err := db.NewDB(config)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		db: db,
	}, nil
}

func (a *Adapter) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := a.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, err
}

func (a *Adapter) GetUserByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	user := &domain.User{}

	err := a.db.WithContext(ctx).Where("user_id = ?", userID).Find(user).Error

	if user.UserID == "" {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *Adapter) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}

	err := a.db.WithContext(ctx).Where("email = ?", email).Find(user).Error

	if user.UserID == "" {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *Adapter) Update(ctx context.Context, userID domain.UserID, firstName, lastName string) (*domain.User, error) {
	user, err := a.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	user.FirstName = firstName
	user.LastName = lastName

	err = a.db.WithContext(ctx).Where("user_id = ?", user.UserID).Save(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *Adapter) DeleteByID(ctx context.Context, userID domain.UserID) error {
	user, err := a.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	err = a.db.WithContext(ctx).Where("user_id = ?", userID).Delete(user).Error
	if err != nil {
		return err
	}

	return nil
}
