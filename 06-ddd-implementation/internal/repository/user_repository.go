package repository

import (
	"06-ddd-implementation/internal/entity"
	"context"
	"database/sql"
)

type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error)
}

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error) {
	// step 1: define query-nya
	query := `INSERT INTO users (username, fullname, email, password) VALUES ($1, $2, $3, $4) RETURNING userid, username, fullname, email, password, profileurl, createdat`

	// step 2: execute query-nya
	row := tx.QueryRowContext(ctx, query, user.Username, user.Fullname, user.Email, user.Password)

	// step 3: scan hasilnya ke *entity.User untuk di return
	var createdUser entity.User
	err := row.Scan(&createdUser.ID, &createdUser.Username, &createdUser.Fullname, &createdUser.Email, &createdUser.Password, &createdUser.ProfileUrl, &createdUser.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}
