package config

import "time"

const RedisMaxIdle = 20
const RedisIdleTimeout = 120 * time.Second
const RedisMaxActive = 100
