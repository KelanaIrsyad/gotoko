package fakers

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/kelanairsyad/gotoko/app/models"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	return &models.User{
		ID:            uuid.New().String(),
		FirstName:     faker.FirstName(),
		LastName:      faker.LastName(),
		Email:         faker.Email(),
		Password:      "Z7PEn.YK$8qUef:gw_0bkZS.)C9R3FCrDeyGu*@-1u@tTdPt.u",
		RememberToken: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}
}
