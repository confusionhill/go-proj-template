package hello

import (
	"database/sql"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

type HelloRepository struct {
	Database *sql.DB
	Redis    *redis.Client
	Mongo    *mongo.Client
}

func Init(db *sql.DB, rd *redis.Client, mg *mongo.Client) (*HelloRepository, error) {
	r := &HelloRepository{Database: db, Redis: rd, Mongo: mg}
	return r, nil
}
