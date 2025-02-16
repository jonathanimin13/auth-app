package repo

import "database/sql"

type AuthRepo interface{}

type authRepoImpl struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) AuthRepo {
	return &authRepoImpl{
		db: db,
	}
}