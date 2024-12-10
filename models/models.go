package models

import (
	"time"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	PasswordHash string
	FirstName    string
	LastName     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Event struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	DateTime    time.Time
	Location    string
	CreatedBy   uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Registration struct {
	ID               uint `gorm:"primaryKey"`
	UserID           uint
	EventID          uint
	RegistrationDate time.Time
	TicketType       string
	NumberOfTickets  int
}

type Notification struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint
	EventID uint
	Message string
	SentAt  time.Time
}

type Payment struct {
	ID             uint `gorm:"primaryKey"`
	RegistrationID uint
	Amount         float64
	Status         string
	PaymentDate    time.Time
}

type Venue struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Address     string
	Capacity    int
	ContactInfo string
}

type Ticket struct {
	ID           uint `gorm:"primaryKey"`
	EventID      uint
	Type         string
	Price        float64
	Availability int
}

type Review struct {
	ID        uint `gorm:"primaryKey"`
	EventID   uint
	UserID    uint
	Rating    int
	Comment   string
	CreatedAt time.Time
}

type Category struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
}

type EventCategory struct {
	ID         uint `gorm:"primaryKey"`
	EventID    uint
	CategoryID uint
}
