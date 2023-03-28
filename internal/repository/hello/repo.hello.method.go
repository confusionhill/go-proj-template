package hello

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"time"
)

func (r HelloRepository) GetHelloMessage() string {
	res, err := r.Redis.Get("mykey").Result()
	if err == nil {
		return res + "\n"
	}
	rows, err := r.Database.Query("SELECT msg from Board LIMIT 1")
	if err != nil {
		log.Error(err)
		return "DB Not Working!"
	}
	var msg string
	for rows.Next() {
		rows.Scan(&msg)
	}
	defer rows.Close()
	rdMsg := fmt.Sprintf("%s from Redis", msg)
	if err = r.Redis.Set("mykey", rdMsg, time.Minute*5).Err(); err != nil {
		log.Error(err)
	}
	if err != nil {
		panic(err)
	}
	return msg + "\n"
}

func (r *HelloRepository) GetHelloFromMongo(ctx context.Context) {
}
