# Golang REDIS

This service demonstrates how to connect to a Redis server using Go. It initializes a Redis client, checks the connection with a `PING` command, and handles errors appropriately.

## Prerequisites

- Go (1.18 or later)
- Redis server (local or remote)
- `go-redis` package

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/mbilalsurathia/golang_redis.git
   cd redis-service
   ```

   Make sure you have a Redis server running on your machine or available on the network. The default host is localhost:6379, but this can be changed in the configuration.

Configuration
You can configure the Redis connection in the main.go file by modifying the RedisConfig struct. The default configuration is:

  ```
RedisConfig{
    Host: "localhost:6379", // Replace with your Redis host
    Type: "tcp",            // Usually, Redis uses TCP connections
    Pass: "",               // Set password if needed
}
  ```

Run the Go application:
  ```
go run main.go
  ```
If the connection is successful, you will see:

  ```
Redis connected successfully!
  ```
