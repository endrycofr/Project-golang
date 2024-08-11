package transaction

import (
	"bwastartup/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       user.User `gorm:"foreignKey:UserID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
