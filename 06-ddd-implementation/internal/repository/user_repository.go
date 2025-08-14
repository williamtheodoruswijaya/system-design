package repository

import (
	"06-ddd-implementation/internal/entity"
	"context"
	"database/sql"
)

type UserRepository interface {
	// Auth
	Register(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error)

	// Find User
	FindByUserID(ctx context.Context, db *sql.DB, userID int) (*entity.User, error)
	FindByUsername(ctx context.Context, db *sql.DB, username string) (*entity.User, error)
	FindByEmail(ctx context.Context, db *sql.DB, email string) (*entity.User, error)
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

func (r *UserRepositoryImpl) FindByUserID(ctx context.Context, db *sql.DB, userID int) (*entity.User, error) {
	// step 1: define query
	query := `SELECT userid, username, fullname, email, password, profileurl, createdat FROM users WHERE userid = $1`

	// step 2: execute query
	row := db.QueryRowContext(ctx, query, userID)

	// step 3: convert to entity
	var user entity.User
	err := row.Scan(&user.ID, &user.Username, &user.Fullname, &user.Email, &user.Password, &user.ProfileUrl, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	// step 4: return
	return &user, nil
}

func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, db *sql.DB, username string) (*entity.User, error) {
	// step 1: define query
	query := `SELECT userid, username, fullname, email, password, profileurl, createdat FROM users WHERE username = $1`

	// step 2: execute query
	row := db.QueryRowContext(ctx, query, username)

	// step 3: convert to entity
	var user entity.User
	err := row.Scan(&user.ID, &user.Username, &user.Fullname, &user.Email, &user.Password, &user.ProfileUrl, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	// step 4: return entity
	return &user, nil
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, db *sql.DB, email string) (*entity.User, error) {
	// step 1: define query
	query := `SELECT userid, username, fullname, email, password, profileurl, createdat FROM users WHERE email = $1`

	// step 2: execute query
	row := db.QueryRowContext(ctx, query, email)

	// step 3: convert to entity
	var user entity.User
	err := row.Scan(&user.ID, &user.Username, &user.Fullname, &user.Email, &user.Password, &user.ProfileUrl, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	// step 4: return entity
	return &user, nil
}
