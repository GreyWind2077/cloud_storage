package redis

import (
	"cloud_storage/config"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var RedisPool *redis.Pool

func init() {
	RedisPool = &redis.Pool{
		MaxIdle:     5,
		MaxActive:   0,
		Wait:        true,
		IdleTimeout: 200 * time.Second,
		Dial: func() (redis.Conn, error) { // 创建链接

			c, err := redis.Dial("tcp", config.Cfg.Redis.Host)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("SELECT", config.Cfg.Redis.Index); err != nil {
				_ = c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { //一个测试链接可用性
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	log.Println("Redis init succeed...")
}

func GetKey(key string) (string, error) {
	rds := RedisPool.Get()
	defer rds.Close()
	return redis.String(rds.Do("GET", key))
}

func SetKey(key, value interface{}, expires int) error {
	rds := RedisPool.Get()
	defer rds.Close()
	if expires == 0 {
		_, err := rds.Do("SET", key, value)
		return err
	} else {
		_, err := rds.Do("SETEX", key, expires, value)
		return err
	}
}

func DelKey(key string) error {
	rds := RedisPool.Get()
	defer rds.Close()
	_, err := rds.Do("DEL", key)
	return err
}
