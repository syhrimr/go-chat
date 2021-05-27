package acc

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/lolmourne/go-accounts/model"
)

type RedisResource struct {
	rdb  *redis.Client
	next DBItf
}

type DBResource struct {
	db *sqlx.DB
}

type DBItf interface {
	Register(username string, password string, salt string) error
	GetUserByUserID(userID int64) (model.User, error)
	GetUserByUserName(userName string) (model.User, error)
	UpdateUserProfpic(userID int64, newProfpic string) error
	UpdateUserName(userID int64, newUsername string) error
	UpdateUserPassword(userID int64, newPassword string) error
}

func NewRedisResource(rdb *redis.Client, next DBItf) DBItf {
	return &RedisResource{
		rdb:  rdb,
		next: next,
	}
}

func NewDBResource(dbParam *sqlx.DB) DBItf {
	return &DBResource{
		db: dbParam,
	}
}
