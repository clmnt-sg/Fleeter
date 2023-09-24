package models

import (
	"github.com/google/uuid"
	"time"
)

type Fleet struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Content   string
	Upvote    uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}
