package entities

import "github.com/google/uuid"

type User struct {
	Id              uuid.UUID `db:"id"`
	UserId          string    `db:"user_id"`
	Name            string    `db:"full_name"`
	Phone           string    `db:"phone"`
	Email           string    `db:"email"`
	Username        string    `db:"username"`
	Password        string    `db:"password"`
	CreatedAt       string    `db:"created_at"`
	UpdatedAt       string    `db:"updated_at"`
	LastSessionDate string    `db:"last_session_date"`
}
