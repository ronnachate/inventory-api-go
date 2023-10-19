package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username  string     `gorm:"type:varchar(50);uniqueIndex;not null"`
	Title     string     `gorm:"type:varchar(50)"`
	Name      string     `gorm:"type:varchar(100);not null"`
	Lastname  string     `gorm:"type:varchar(100)"`
	StatusID  int        `gorm:"not null"`
	Status    UserStatus `gorm:"ForeignKey:StatusID"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}

type UserRepository interface {
	GetByID(c context.Context, id string) (User, error)
}

type UserUsecase interface {
	GetByID(c context.Context, userID string) (User, error)
}
