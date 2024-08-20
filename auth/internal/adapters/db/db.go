package db

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Bookil/microservices/auth/config"
	"github.com/Bookil/microservices/auth/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	Adapter struct {
		db *gorm.DB
	}
)

var (
	dbInc *Adapter
	mutex = new(sync.Mutex)
)

func generateURL(url *config.Mysql) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", url.Username, url.Password, url.Host, url.Port, url.DBName)
}

func NewAdapter(url *config.Mysql) (*Adapter, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if dbInc == nil {
		genUrl := generateURL(url)

		log.Println("URL:", genUrl)

		db, openErr := gorm.Open(mysql.Open((genUrl)), &gorm.Config{})
		if openErr != nil {
			return nil, fmt.Errorf("db connection error: %v", openErr)
		}

		err := db.AutoMigrate(&domain.Auth{})
		if err != nil {
			return nil, fmt.Errorf("db migration error: %v", err)
		}
		dbInc = &Adapter{db: db}
	}
	return dbInc, nil
}

func (a *Adapter) Create(ctx context.Context, auth *domain.Auth) (*domain.Auth, error) {
	err := a.db.WithContext(ctx).Create(auth).Error
	if err != nil {
		return nil, err
	}

	return auth, err
}

func (a *Adapter) GetByID(ctx context.Context, userID domain.UserID) (*domain.Auth, error) {
	auth := &domain.Auth{}

	err := a.db.WithContext(ctx).Where("user_id = ?", userID).Find(auth).Error

	if auth.UserID == "" {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (a *Adapter) ChangePassword(ctx context.Context, userID domain.UserID, hashedPassword string) (*domain.Auth, error) {
	auth, err := a.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	auth.HashedPassword = hashedPassword

	savedAuth, err := a.Save(ctx, auth)
	if err != nil {
		return nil, err
	}

	return savedAuth, nil
}

func (a *Adapter) ClearFailedLoginAttempts(ctx context.Context, userID domain.UserID) (*domain.Auth, error) {
	auth, err := a.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	auth.FailedLoginAttempts = 0

	savedAuth, err := a.Save(ctx, auth)
	if err != nil {
		return nil, err
	}

	return savedAuth, nil
}

func (a *Adapter) LockAccount(ctx context.Context, userID domain.UserID, lockDuration time.Duration) (*domain.Auth, error) {
	auth, err := a.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	auth.AccountLockedUntil = time.Now().Add(lockDuration).Unix()

	savedAuth, err := a.Save(ctx, auth)
	if err != nil {
		return nil, err
	}

	return savedAuth, nil
}

func (a *Adapter) UnlockAccount(ctx context.Context, userID domain.UserID) (*domain.Auth, error) {
	auth, err := a.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	auth.AccountLockedUntil = 0

	savedAuth, err := a.Save(ctx, auth)
	if err != nil {
		return nil, err
	}

	return savedAuth, err
}

func (a *Adapter) IncrementFailedLoginAttempts(ctx context.Context, userID domain.UserID) (*domain.Auth, error) {
	auth, err := a.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	auth.FailedLoginAttempts++

	savedAuth, err := a.Save(ctx, auth)
	if err != nil {
		return nil, err
	}

	return savedAuth, err
}

func (a *Adapter) VerifyEmail(ctx context.Context, userID domain.UserID) (*domain.Auth, error) {
	auth, err := a.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	auth.IsEmailVerified = true

	savedAuth, err := a.Save(ctx, auth)
	if err != nil {
		return nil, err
	}

	return savedAuth, nil
}

func (a *Adapter) DeleteByID(ctx context.Context, userID domain.UserID) error {
	auth, err := a.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	err = a.db.WithContext(ctx).Where("user_id = ?", userID).Delete(auth).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) Save(ctx context.Context, auth *domain.Auth) (*domain.Auth, error) {
	err := a.db.WithContext(ctx).Where("user_id = ?", auth.UserID).Save(auth).Error
	if err != nil {
		return nil, err
	}

	return auth, nil
}
