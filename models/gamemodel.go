package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Title                 string
	Franchise             string
	Platform              string
	MainQuests            int
	SideQuests            int
	CompletedMainQuests   int
	CompletedSideQuests   int
	Collectibles          int
	CollectedCollectibles int
}
