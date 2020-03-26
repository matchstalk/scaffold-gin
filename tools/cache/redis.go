package cache

import (
	"fmt"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v7"
	"github.com/matchstalk/redisqueue"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

// Redis cache implement
type Redis struct {
	client *redis.Client
	mutex  *redislock.Client
}

// Setup connection
func (r *Redis) Connect() {
	r.client = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.host"),
		Password:     viper.GetString("redis.auth"),
		DB:           viper.GetInt("redis.dao"),
		PoolSize:     viper.GetInt("redis.pool.max"),
		MinIdleConns: viper.GetInt("redis.pool.min"),
	})
	_, err := r.client.Ping().Result()
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not connected to redis : %s", err.Error()))
	}
	r.mutex = redislock.New(r.client)
	log.Info("Successfully connected to redis")
}

// Get from key
func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

// Set value with key and expire time
func (r *Redis) Set(key string, val interface{}, expire int) error {
	return r.client.Set(key, val, time.Duration(expire)*time.Second).Err()
}

// Del delete key in redis
func (r *Redis) Del(key string) error {
	return r.client.Del(key).Err()
}

// HashGet from key
func (r *Redis) HashGet(hk, key string) (string, error) {
	return r.client.HGet(hk, key).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *Redis) HashDel(hk, key string) error {
	return r.client.HDel(hk, key).Err()
}

// Increase
func (r *Redis) Increase(key string) error {
	return r.client.Incr(key).Err()
}

// Set ttl
func (r *Redis) Expire(key string, dur time.Duration) error {
	return r.client.Expire(key, dur).Err()
}

// New Consumer
func (r *Redis) NewConsumer() (*redisqueue.Consumer, error) {
	return redisqueue.NewConsumerWithOptions(&redisqueue.ConsumerOptions{
		VisibilityTimeout: time.Duration(viper.GetInt64("redis.queue.visibility_timeout")) * time.Second,
		BlockingTimeout:   time.Duration(viper.GetInt64("redis.queue.blocking_timeout")) * time.Second,
		ReclaimInterval:   time.Duration(viper.GetInt64("redis.queue.reclaim_interval")) * time.Second,
		BufferSize:        viper.GetInt("redis.queue.buffer_size"),
		Concurrency:       viper.GetInt("redis.queue.concurrency"),
		RedisClient:       r.client,
	})
}

// New Producer
func (r *Redis) NewProducer() (*redisqueue.Producer, error) {
	return redisqueue.NewProducerWithOptions(&redisqueue.ProducerOptions{
		StreamMaxLength:      viper.GetInt64("redis.queue.stream_max_length"),
		ApproximateMaxLength: viper.GetBool("redis.queue.approximate_max_length"),
		RedisClient:          r.client,
	})
}

func (r *Redis) Lock(key string, ttl int64, options *redislock.Options) (*redislock.Lock, error) {
	if r.mutex == nil {
		r.mutex = redislock.New(r.client)
	}
	return r.mutex.Obtain(key, time.Duration(ttl)*time.Second, options)
}
