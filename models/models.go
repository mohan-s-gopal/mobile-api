package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GenerateISOString generates a time string equivalent to Date.now().toISOString in JavaScript
func GenerateISOString() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.999Z07:00")
}

// Base contains common columns for all tables
type UserBase struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UUID      uuid.UUID `json:"uuid" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

// BeforeCreate will set Base struct before every insert
func (userBase *UserBase) BeforeCreate(tx *gorm.DB) error {
	// generate a uuid and save that as a string
	uuid := uuid.New()
	userBase.UUID = uuid

	// generate timestamps
	t := GenerateISOString()
	userBase.CreatedAt, userBase.UpdatedAt = t, t

	return nil
}

// AfterUpdate will update the Base struct after every update
func (base *Base) AfterUpdate(tx *gorm.DB) error {
	// update timestamps
	base.UpdatedAt = GenerateISOString()
	return nil
}

type Base struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// BeforeCreate will set Base struct before every insert
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	// generate timestamps
	t := GenerateISOString()
	base.CreatedAt, base.UpdatedAt = t, t

	return nil
}
