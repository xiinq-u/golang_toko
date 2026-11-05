package models

import "time"

type Category struct {
	ID        int
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description,omitempty"`
	Price        int       `json:"price"`
	Quantity     int       `json:"quantity"`
	CategoryID   *int      `json:"category_id,omitempty"`
	CategoryName string    `json:"category_name,omitempty"`
	Image        string    `json:"image,omitempty"`
	Unit         string    `json:"unit"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Profile struct {
	ID          int
	AuthUserID  string
	FullName    string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Transaction struct {
	ID        int
	UserID    string
	Total     int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransactionItem struct {
	ID            int
	TransactionID int
	ProductID     int
	Quantity      int
	Price         int
	CreatedAt     time.Time
}

type HeroSlider struct {
	ID        int
	Title     string
	Subtitle  string
	ImageURL  string
	Link      string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
