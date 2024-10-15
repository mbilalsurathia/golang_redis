package main

import (
	"fmt"

	redis "example.com/m/Redis"
)

// RedisConfig holds the configuration for Redis
type RedisConfig struct {
	Host string
	Type string
	Pass string
}

func main() {
	// Set up the Redis configuration
	c := RedisConfig{
		Host: "localhost:6379", // Replace with your Redis host
		Type: "tcp",            // Usually, Redis uses TCP connections
		Pass: "",               // Set password if needed
	}

	_, err := redis.Connect(redis.RedisConfig{
		Host: c.Host,
		Type: c.Type,
		Pass: c.Pass,
	})
	if err != nil {
		panic(fmt.Errorf("failed to init redisClient: %w", err))
	}

	fmt.Println("Redis connected successfully!")
}
