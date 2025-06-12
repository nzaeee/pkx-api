package persistence

import (
	"context"
	"database/sql"
	"pkx-api/internal/domain/entity"
	"pkx-api/internal/domain/repository"
	"time"

	"pkx-api/internal/infrastructure/persistence/db"
)

type userRepository struct {
	q *db.Queries
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{
		q: db.New(db),
	}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	params := db.CreateUserParams{
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	result, err := r.q.CreateUser(ctx, params)
	if err != nil {
		return err
	}

	user.ID = result.ID
	return nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	result, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrUserNotFound
		}
		return nil, err
	}

	return &entity.User{
		ID:        result.ID,
		Email:     result.Email,
		Password:  result.Password,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	result, err := r.q.GetUserByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrUserNotFound
		}
		return nil, err
	}

	return &entity.User{
		ID:        result.ID,
		Email:     result.Email,
		Password:  result.Password,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	params := db.UpdateUserParams{
		ID:        user.ID,
		Email:     sql.NullString{String: user.Email, Valid: true},
		Password:  sql.NullString{String: user.Password, Valid: true},
		UpdatedAt: time.Now(),
	}

	_, err := r.q.UpdateUser(ctx, params)
	return err
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteUser(ctx, id)
} 
