package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `db:"id"`
	Username  string `db:"username"`
	Email     string `db:"email"`
	Password  string `db:"-"`
	CreatedAt string `db:"created_at"`
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`

	err := s.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}
