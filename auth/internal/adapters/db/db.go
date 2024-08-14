package db

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Bookil/microservices/auth/config"
	"github.com/Bookil/microservices/auth/internal/application/core/domain"
	"github.com/Bookil/microservices/auth/internal/ports"
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

func NewDB(db *gorm.DB) ports.DBPort {
	return &Adapter{db: db}
}

func (a *Adapter) Create(ctx context.Context, auth *domain.Auth) error {
	err := a.db.WithContext(ctx).Create(auth).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) GetByID(ctx context.Context, userID domain.UserID) (*domain.Auth, error) {
	auth := &domain.Auth{}

	err := a.db.WithContext(ctx).Where("user_id = ?", userID).Find(auth).Error
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (a *Adapter) ChangePassword(ctx context.Context, userID domain.UserID, hashedPassword string) error {
	auth, err := a.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	auth.HashedPassword = hashedPassword

	err = a.Save(ctx, auth)
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) ClearFailedLoginAttempts(ctx context.Context, userID domain.UserID) error {
	auth, err := a.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	auth.FailedLoginAttempts = 0

	err = a.Save(ctx, auth)
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) LockAccount(ctx context.Context, userID domain.UserID, lockDuration time.Duration) error {
	auth, err := a.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	auth.AccountLockedUntil = time.Now().Add(lockDuration).Unix()

	err = a.Save(ctx, auth)
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) UnlockAccount(ctx context.Context, userID domain.UserID) error {
	auth, err := a.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	auth.AccountLockedUntil = 0

	err = a.Save(ctx, auth)
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) IncrementFailedLoginAttempts(ctx context.Context, userID domain.UserID) error {
	auth, err := a.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	auth.FailedLoginAttempts++

	err = a.Save(ctx, auth)
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) VerifyEmail(ctx context.Context, userID domain.UserID) error {
	auth, err := a.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	auth.IsEmailVerified = true

	err = a.db.Save(auth).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) DeleteByID(ctx context.Context, userID domain.UserID) error {
	auth, err := a.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	err = a.db.WithContext(ctx).Delete(auth).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) Save(ctx context.Context, auth *domain.Auth) error {
	err := a.db.WithContext(ctx).Save(auth).Error
	if err != nil {
		return err
	}

	return nil
}
