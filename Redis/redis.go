package redis

import (
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Redis interface {
	Set(key Key, value string) error
	SetWithExpiry(key Key, value string, seconds int) error
	Get(key Key) (string, error)
	Mget(keys ...string) ([]string, error)
	Del(key string) error
	RPush(key string, values ...interface{}) error
	LPush(key string, values ...interface{}) error
	RPop(key string) (string, error)
	LPop(key string) (string, error)
	LRange(key string, start, stop int) ([]string, error)
	Keys(key string) ([]string, error)
}

var (
	ErrKeyNotFound = fmt.Errorf("key not found in redis")
)

type UserLoginSession struct {
	Device   string `json:"device,omitempty"`
	Location string `json:"location,omitempty"`
	IP       string `json:"ip,omitempty"`
	JwtToken string `json:"jwt_token,omitempty"`
}

type RedisConfig struct {
	Host string
	Type string
	Pass string
}

type redisClient struct {
	client *redis.Redis
}

func Connect(conf RedisConfig) (Redis, error) {
	redisConf := redis.RedisConf{
		Host: conf.Host,
		Type: conf.Type,
		Pass: conf.Pass,
	}

	redisInstance, err := redis.NewRedis(redisConf)
	if err != nil {
		return nil, fmt.Errorf("unable to bootstrap redis: %w", err)
	}

	return &redisClient{
		client: redisInstance,
	}, nil
}

func (r *redisClient) Set(key Key, value string) error {
	err := r.client.Set(key.String(), value)
	if err != nil {
		return fmt.Errorf("error setting key in redis: %w", err)
	}

	return nil
}

func (r *redisClient) Get(key Key) (string, error) {
	result, err := r.client.Get(key.String())
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrKeyNotFound
		}

		return "", fmt.Errorf("error getting key from redis: %w", err)
	}

	return result, nil
}

func (r *redisClient) Mget(keys ...string) ([]string, error) {
	result, err := r.client.Mget(keys...)
	if err != nil {
		return nil, fmt.Errorf("failed to redis.Mget: %w", err)
	}

	return result, nil
}

func (r *redisClient) Del(key string) error {
	if _, err := r.client.Del(key); err != nil {
		return fmt.Errorf("error removing key in redis: %w", err)
	}

	return nil
}

func (r *redisClient) RPush(key string, values ...interface{}) error {
	_, err := r.client.Rpush(key, values...)

	return err
}

func (r *redisClient) LPush(key string, values ...interface{}) error {
	_, err := r.client.Lpush(key, values...)

	return err
}

func (r *redisClient) RPop(key string) (string, error) {
	val, err := r.client.Rpop(key)
	if err != nil && !errors.Is(err, redis.Nil) {
		return "", err
	}

	return val, nil
}

func (r *redisClient) LPop(key string) (string, error) {
	val, err := r.client.Lpop(key)
	if err != nil && !errors.Is(err, redis.Nil) {
		return "", err
	}

	return val, nil
}

func (r *redisClient) LRange(key string, start, stop int) ([]string, error) {
	vals, err := r.client.Lrange(key, start, stop)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	return vals, nil
}

func (r *redisClient) Keys(key string) ([]string, error) {
	keys, err := r.client.Keys(fmt.Sprintf("%s*", key))
	if err != nil {
		return nil, err
	}

	return keys, nil
}

// set key in REDIS with expiry in seconds
func (r *redisClient) SetWithExpiry(key Key, value string, seconds int) error {
	err := r.client.Setex(key.String(), value, seconds)
	if err != nil {
		return fmt.Errorf("error setting key in redis: %w", err)
	}

	return nil
}
