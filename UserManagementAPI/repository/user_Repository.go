package repository

import (
	"context"
	"database/sql"

	"github.com/joshua468/usermanagementapi/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	user := &model.User{}
	err := r.db.QueryRowContext(ctx, "SELECT id,name,age FROM users Where id =?", id).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users(name,age)VALUES(?, ?)", user.Name, user.Age)
	if err != nil {
		return err
	}
	return nil
}
