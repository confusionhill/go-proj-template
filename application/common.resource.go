package application

import (
	"context"
	"database/sql"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"project-template/internal/config"
)

type CommonResource struct {
	Database *sql.DB
	Redis    *redis.Client
	Mongo    *mongo.Client
}

func InitResource(cfg *config.MainConfig) (*CommonResource, error) {
	r := new(CommonResource)
	db, err := sql.Open("postgres", cfg.Database.Host)
	if err != nil {
		return nil, err
	}
	r.Database = db
	if err = r.Database.Ping(); err != nil {
		return nil, err
	}
	log.Println("DATABASE INITIATED!")

	rdc := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	r.Redis = rdc
	log.Println("REDIS INITIATED!")

	clientOptions := options.Client().ApplyURI(cfg.Mongo.Host)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	r.Mongo = client
	log.Println("MONGODB INITIATED!")

	log.Println("RESOURCE INITIATED")
	return r, nil
}
