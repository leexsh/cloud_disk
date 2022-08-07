package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var ctx = context.Background()

func TestRedis(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "121.5.114.14:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	res, err := rdb.Ping(ctx).Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)

	// set
	err = rdb.Set(ctx, "key", "val", time.Second*10).Err()
	if err != nil {
		t.Fatal(err)
	}

	// get
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "val", val)

}
