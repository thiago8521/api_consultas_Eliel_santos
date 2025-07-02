package models

import "time"

type Appointment struct {
    ID        uint      `gorm:"primaryKey"`
    PatientID uint      `gorm:"index;not null"`
    Datetime  time.Time `gorm:"not null;index"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
