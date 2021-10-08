package redis

import (
    "context"
    "strconv"
    "strings"

    "github.com/go-redis/redis/v7"
)

type Config struct {
    Address          string
    Password         string
    Database         string
    MasterName       string
    SentinelPassword string
    // for datacollection, notify if the snapshot redis is in use
    Enable       string
    MaxOpenConns int
}

func NewFromConfig(cfg Config) (Client, error) {
    dbNum, err := strconv.Atoi(cfg.Database)
    if nil != err {
        return nil, err
    }
    if cfg.MaxOpenConns == 0 {
        cfg.MaxOpenConns = 3000
    }

    var client Client
    if cfg.MasterName == "" {
        option := &redis.Options{
            Addr:     cfg.Address,
            Password: cfg.Password,
            DB:       dbNum,
            PoolSize: cfg.MaxOpenConns,
        }
        client = NewClient(option)
    } else {
        hosts := strings.Split(cfg.Address, ",")
        option := &redis.FailoverOptions{
            MasterName:       cfg.MasterName,
            SentinelAddrs:    hosts,
            Password:         cfg.Password,
            DB:               dbNum,
            PoolSize:         cfg.MaxOpenConns,
            SentinelPassword: cfg.SentinelPassword,
        }
        client = NewFailoverClient(option)
    }

    err = client.Ping(context.Background()).Err()
    if err != nil {
        return nil, err
    }

    return client, err
}

func IsNilErr(err error) bool {
    return redis.Nil == err
}
