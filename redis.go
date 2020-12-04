package tools

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisMethod struct {
	Pool             *redis.Pool
	RedisHost        string
	RedisPort        string
	RedisPassword    string
	RedisDb          string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout int
}

var RedisClient RedisMethod

func initRedis() error{
	RedisClient.Pool = &redis.Pool{
		MaxIdle:     RedisClient.RedisMaxIdle,
		MaxActive:   RedisClient.RedisMaxActive,
		IdleTimeout: time.Duration(RedisClient.RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp", RedisClient.RedisHost+":"+RedisClient.RedisPort)
			//redis.DialConnectTimeout(10*time.Second),
			//redis.DialReadTimeout(10*time.Second),
			//redis.DialWriteTimeout(10*time.Second))
			if err!=nil {
				return nil, err
			}
			if len(RedisClient.RedisPassword) != 0 {
				_, err = c.Do("AUTH", RedisClient.RedisPassword)
				if err!=nil {
					return nil, err
				}
			}
			if len(RedisClient.RedisDb) != 0 {
				_, err = c.Do("SELECT", RedisClient.RedisDb)
				if err!=nil {
					return nil, err
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		Wait: true,
	}
	return nil
}

func (r *RedisMethod) Announce(key, s string) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	c.Do("PUBLISH", key, s)
}

func (r *RedisMethod) Set(key string, value interface{}, time int64) (bool, error) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	_, err := c.Do("SET", key, value, "EX", time)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func (r *RedisMethod) Get(key string) (interface{}, error) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	return c.Do("GET", key)
}

func (r *RedisMethod) Del(key string) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	c.Do("DEL", key)
}

func (r *RedisMethod) Incr(key string) (interface{}, error) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	return c.Do("INCR", key)
}

func (r *RedisMethod) Expire(key string, num int64) (interface{}, error) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	return c.Do("EXPIRE", key, num)
}

func (r *RedisMethod) EXISTS(key string) (int, error) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	return redis.Int(c.Do("EXISTS", key))
}

func (r *RedisMethod) Info() (string, error) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	return redis.String(c.Do("INFO"))
}

func (r *RedisMethod) Lpush(key string, value interface{}) (bool, error) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	_, err := c.Do("LPUSH", key, value)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func (r *RedisMethod) Rpop(key string) (interface{}, error) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	return c.Do("RPOP", key)
}

func (r *RedisMethod) Do(method string, arg ...interface{}) (reply interface{}, err error) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	return c.Do(method, arg)
}

func (r *RedisMethod) HIncrByFloat(key, field string, value float64) (bool, error) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	_, err := c.Do("HINCRBYFLOAT", key, field, value)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func (r *RedisMethod) HIncrBy(key, field string, value int64) (bool, error) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	_, err := c.Do("HINCRBY", key, field, value)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func (r *RedisMethod) HGet(key, field string) (interface{}, error) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	return c.Do("HGET", key, field)
}

func (r *RedisMethod) HKeys(key string) ([]string, error) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	return redis.Strings(c.Do("HKEYS", key))
}

func (r *RedisMethod) HDel(key, field string) (bool, error) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	_, err := c.Do("HDEL", key, field)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func (r *RedisMethod) SAdd(key string, value string) (bool, error) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	_, err := c.Do("SADD", key, value)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func (r *RedisMethod) SMembers(key string) ([]string, error) {
	c := RedisClient.Pool.Get()
	defer c.Close()
	return redis.Strings(c.Do("SMEMBERS", key))
}
