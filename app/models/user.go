package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"uniqueIndex"`
	DisplayName string
	Email       string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
