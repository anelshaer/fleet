package redis

import (
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/fleetdm/fleet/v4/server/fleet"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/require"
)

func SetupRedis(tb testing.TB, cluster, redir bool) fleet.RedisPool {
	if _, ok := os.LookupEnv("REDIS_TEST"); !ok {
		tb.Skip("set REDIS_TEST environment variable to run redis-based tests")
	}
	if cluster && (runtime.GOOS == "darwin" || runtime.GOOS == "windows") {
		tb.Skipf("docker networking limitations prevent running redis cluster tests on %s", runtime.GOOS)
	}

	var (
		addr     = "127.0.0.1:"
		password = ""
		database = 0
		useTLS   = false
		port     = "6379"
	)
	if cluster {
		port = "7001"
	}
	addr += port

	pool, err := NewRedisPool(PoolConfig{
		Server:                    addr,
		Password:                  password,
		Database:                  database,
		UseTLS:                    useTLS,
		ConnTimeout:               5 * time.Second,
		KeepAlive:                 10 * time.Second,
		ClusterFollowRedirections: redir,
	})
	require.NoError(tb, err)

	conn := pool.Get()
	defer conn.Close()
	_, err = conn.Do("PING")
	require.Nil(tb, err)

	tb.Cleanup(func() {
		err := EachRedisNode(pool, func(conn redis.Conn) error {
			_, err := conn.Do("FLUSHDB")
			return err
		})
		require.NoError(tb, err)
		pool.Close()
	})

	return pool
}
