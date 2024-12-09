package models

import (
	"time"
)

// User entity
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

// Event entity
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

// Registration entity
type Registration struct {
	ID               uint `gorm:"primaryKey"`
	UserID           uint
	EventID          uint
	RegistrationDate time.Time
	TicketType       string
	NumberOfTickets  int
}

// Notification entity
type Notification struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint
	EventID uint
	Message string
	SentAt  time.Time
}

// Payment entity
type Payment struct {
	ID             uint `gorm:"primaryKey"`
	RegistrationID uint
	Amount         float64
	Status         string
	PaymentDate    time.Time
}

// Venue entity
type Venue struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Address     string
	Capacity    int
	ContactInfo string
}

// Ticket entity
type Ticket struct {
	ID           uint `gorm:"primaryKey"`
	EventID      uint
	Type         string
	Price        float64
	Availability int
}

// Review entity
type Review struct {
	ID        uint `gorm:"primaryKey"`
	EventID   uint
	UserID    uint
	Rating    int
	Comment   string
	CreatedAt time.Time
}

// Category entity
type Category struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
}

// EventCategory entity
type EventCategory struct {
	ID         uint `gorm:"primaryKey"`
	EventID    uint
	CategoryID uint
}
