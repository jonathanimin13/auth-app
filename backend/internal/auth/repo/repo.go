package repo

import (
	"auth-app/internal/apperrors"
	"auth-app/internal/entity"
	"auth-app/pkg/customerror"
	"context"
	"database/sql"
)

type AuthRepo interface{
	IsEmailExists(ctx context.Context, email string) (bool, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type authRepoImpl struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) AuthRepo {
	return &authRepoImpl{
		db: db,
	}
}

func (r *authRepoImpl) IsEmailExists(ctx context.Context, email string) (bool, error) {
	query := `SELECT EXISTS (
							SELECT 
								1
							FROM 
								users
							WHERE 
								email = $1
							AND 
								deleted_at IS NULL
						)`

	var isExists bool
	err := r.db.QueryRowContext(ctx, query, email).Scan(&isExists)
	if err != nil {
		return false, customerror.NewInternalServerError(apperrors.FieldServer, apperrors.ErrInternalServer, err)
	}

	return isExists, nil
}

func (r *authRepoImpl) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT
							id,
							username,
							password
						FROM
							users
						WHERE 
							email = $1
						AND 
							deleted_at IS NULL`

	var user entity.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)
	if err != nil {
		return nil, customerror.NewInternalServerError(apperrors.FieldServer, apperrors.ErrInternalServer, err)
	}

	return &user, nil
}