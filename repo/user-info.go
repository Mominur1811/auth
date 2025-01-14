package repo

import (
	"auth-repo/authentication"
	"auth-repo/types"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type UserInfoRepo interface {
	authentication.UserInfoRepo
}
type userInfoRepo struct {
	tableName string
	db        *sqlx.DB
	psql      sq.StatementBuilderType
}

func NewUserInfoRepo(db *db) UserInfoRepo {
	return &userInfoRepo{
		tableName: `"user_info"`,
		db:        db.DB,
		psql:      db.psql,
	}
}

func (r *userInfoRepo) RegisterUser(ctx context.Context, user *types.UserInfo) (*types.UserInfo, error) {

	query, args, err := r.psql.Insert(r.tableName).Columns("email", "password", "role").Values(user.Email, user.Password, user.Role).Suffix("RETURNING *").ToSql()

	var userInfo types.UserInfo

	err = r.db.QueryRowxContext(ctx, query, args...).StructScan(&userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}
