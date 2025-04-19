package user

import "time"

type User struct {
	ID        uint    `gorm:"primaryKey:autoincrement"`
	Name      string  `gorm:"not null"`
	LastName  string  `gorm:"not null"`
	Email     string  `gorm:"uniqueIndex;not null"`
	Password  string  `gorm:"not null"`
	CPF       string  `gorm:"uniqueIndex; not null"`
	Phone     string  `gorm:"not null"`
	Role      string  `gorm:"not null"`
	Address   Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdadedAt time.Time
}

type Address struct {
	ID         uint   `gorm:"primaryKey:autoincrement"`
	Street     string `gorm:"not null"`
	Number     string `gorm:"not null"`
	City       string `gorm:"not null"`
	State      string `gorm:"not null"`
	Complement *string
	ZipCode    string `gorm:"not null"`
	UserID     uint   `gorm:"uniqueIndex;not null"`
	CreatedAt  time.Time
	UpdadedAt  time.Time
}
