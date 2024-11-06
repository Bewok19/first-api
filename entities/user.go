package entities

import "time"

type User struct {
	ID        uint64    `gorm:"primaryKey:autoIncrement"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Age       int       `gorm:"column:age"`
	CreatedAt time.Time `gorm:"column:created_at"`
	updatedAt time.Time `gorm:"column:updated_at"`
}
