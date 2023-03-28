package hello

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"time"
)

func (r HelloRepository) GetHelloMessage() (string, error) {
	res, err := r.Redis.Get("mykey").Result()
	if err == nil {

		return fmt.Sprintf("%s\n", res), nil
	}
	rows, err := r.Database.Query("SELECT msg from Board LIMIT 1")
	if err != nil {
		log.Error(err)
		return "", err
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
	return fmt.Sprintf("%s\n", msg), nil
}

func (r *HelloRepository) GetHelloFromMongo() {
}
