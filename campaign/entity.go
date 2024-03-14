package campaign

import "time"

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string // Huruf besar pada "Slug"
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage // Menggunakan foreignKey untuk menunjukkan kunci asing
}

type CampaignImage struct {
	ID         int
	CampaignID int // Huruf besar pada "CampaignID"
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
