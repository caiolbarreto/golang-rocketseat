package models

import "time"

type Loan struct {
	ID         string    `json:"id"`
	BookID     string    `json:"bookId"`
	UserID     string    `json:"userId"`
	Status     string    `json:"status"`
	BorrowedAt time.Time `json:"borrowedAt"`
	ReturnedAt time.Time `json:"returnedAt"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
