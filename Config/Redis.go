package Config

import (
	"github.com/go-redis/redis"
	"mygra.tech/project1/Utils/Helpers/Log"
)

var (
	Client *redis.Client
)

func InitRedis() (*redis.Client, error) {
	const FILE_NAME string = "Redis"

	Client := redis.NewClient(&redis.Options{
		Addr:       "127.0.0.1:6379",
		MaxRetries: 3,
		DB:         0,
	})

	pong, err := Client.Ping().Result()

	if err != nil {
		Log.INFO("%s: Error occurred %v\n", FILE_NAME, err.Error())
		return Client, err
	}

	Log.INFO("%s: Success connected %v\n", FILE_NAME, pong)
	return Client, nil
}

func GetRedis() *redis.Client {
	return Client
}
