package models

import "time"

type Patient struct {
    ID           uint          `gorm:"primaryKey"`
    Name         string        `gorm:"not null"`
    Email        string        `gorm:"uniqueIndex;not null"`
    PasswordHash string        `gorm:"not null"`
    Appointments []Appointment `gorm:"constraint:OnDelete:CASCADE"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
