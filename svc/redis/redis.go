package redis

import (
	"context"
	"os"
	"strings"

	"github.com/go-logr/logr"
	redis "github.com/redis/go-redis/v9"
)

var (

	// should the operator use a connection to redis or not
	USEREDIS bool = false

	RDS *redis.ClusterClient
)

// handle the redis connection tests
func RedisConnected() error {

	if !USEREDIS {
		return nil
	}

	if err := RDS.ForEachShard(context.Background(), func(ctx context.Context, client *redis.Client) error {
		return client.Ping(ctx).Err()
	}); err != nil {
		return err
	}

	return nil
}

// create the redis connection info
func Connect(_log logr.Logger) error {

	// check if redis is configured
	clusteraddresses, set := os.LookupEnv("REDIS_CLUSTER_ADDRESSES")
	if !set {
		_log.Info("redis not used")
		return nil
	}

	RDS = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: strings.Split(clusteraddresses, ","),
		NewClient: func(opt *redis.Options) *redis.Client {

			opt.Username = os.Getenv("REDIS_USER")

			opt.Password = os.Getenv("REDIS_PASSWORD")

			opt.DB = 0

			return redis.NewClient(opt)
		},
	})

	_log.Info("redis connection configured", "endpoints", strings.Split(clusteraddresses, ","))

	// testing connection
	if err := RDS.ForEachShard(context.Background(), func(ctx context.Context, client *redis.Client) error {

		return client.Ping(ctx).Err()

	}); err != nil {

		return err
	}

	// since the connection tests are successful, now we can enable redis
	USEREDIS = true
	return nil
}

// disconnect from redis
func Disconnect() error {
	return RDS.Close()
}
