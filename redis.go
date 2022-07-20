package pubghealth

import (
	"time"
	"github.com/heptiolabs/healthcheck"
	"github.com/go-redis/redis"
)

// RedisPingCheck issues a Ping command to Redis on a timeout to verify connectivity.
func RedisPingCheck(client *redis.Client, timeout time.Duration, frequency time.Duration) healthcheck.Check {
	checkFunc := func() error {
		_, err := client.Ping().Result()
		return err
	}
	return healthcheck.Timeout(healthcheck.Async(checkFunc, frequency), timeout)
}
