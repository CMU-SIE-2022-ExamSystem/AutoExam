package initialize

import (
	"fmt"
	"strconv"
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
)

func InitRedis() {
	info := global.Settings.Redis

	// create redis connection pool
	redisPool := &redis.Pool{
		MaxIdle:     3,                 // maximum number of idle connections in the pool
		MaxActive:   0,                 // maximum number of connections allocated by the pool at a given time
		IdleTimeout: 240 * time.Second, // close connections after remaining idle for this duration
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL("redis://" + info.Host + ":" + strconv.Itoa(info.Port))
			if err != nil {
				fmt.Println("redis error")
				fmt.Println(err)
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			fmt.Println("redis error")
			fmt.Println(err)
			return err
		},
	}
	var err error
	// initialize celery client
	global.Redis, err = gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		&gocelery.RedisCeleryBackend{Pool: redisPool},
		5, // number of workers
	)
	if err != nil {
		fmt.Println("redis error")
		fmt.Println(err)
	}

}
