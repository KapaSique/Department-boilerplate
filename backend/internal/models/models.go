package models

import "time"

type News struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Document struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	FileURL     string    `json:"file_url"`
	Category    string    `json:"category"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type BudgetReport struct {
	ID        int       `json:"id"`
	Year      int       `json:"year"`
	Quarter   int       `json:"quarter"`
	Title     string    `json:"title"`
	Revenue   float64   `json:"revenue"`
	Expenses  float64   `json:"expenses"`
	Balance   float64   `json:"balance"`
	FileURL   string    `json:"file_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type DepartmentInfo struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Address     string   `json:"address"`
	Phone       string   `json:"phone"`
	Email       string   `json:"email"`
	WorkHours   string   `json:"work_hours"`
	Head        string   `json:"head"`
	Departments []string `json:"departments"`
}

type Contact struct {
	Department string `json:"department"`
	Head       string `json:"head"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}
