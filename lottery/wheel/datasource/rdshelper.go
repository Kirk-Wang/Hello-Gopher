package datasource

import (
	"sync"

	"github.com/gomodule/redigo/redis"
)

var rdsLock sync.Mutex
var cacheInstance *redis.Conn

type RedisConn struct {
	pool      *redis.pool
	showDebug bool
}

func (rds *RedisConn) Do(commandName string,
	args ...interface{}) (reply interface{}, err error) {
	conn := rds.pool.Get()
	defer conn.Close()
}
