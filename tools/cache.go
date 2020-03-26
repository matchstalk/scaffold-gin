package tools

import (
	"github.com/bsm/redislock"
	"github.com/matchstalk/redisqueue"
	"github.com/matchstalk/scaffold-gin/tools/cache"
	"time"
)

type Cache struct {
	prefix string
}

func NewCache(prefix string) Cache {
	return Cache{
		prefix: prefix,
	}
}

func (e Cache) Connect() {
}

// Get val in cache
func (e Cache) Get(key string) (string, error) {
	return cache.Get(e.prefix + "." + key)
}

// Set val in cache
func (e Cache) Set(key string, val interface{}, expire int) error {
	return cache.Set(e.prefix+"."+key, val, expire)
}

// Del delete key in cache
func (e Cache) Del(key string) error {
	return cache.Del(e.prefix + "." + key)
}

// HashGet get val in hashtable cache
func (e Cache) HashGet(hk, key string) (string, error) {
	return cache.HashGet(hk, e.prefix+"."+key)
}

// HashDel delete one key:value pair in hashtable cache
func (e Cache) HashDel(hk, key string) error {
	return cache.HashDel(hk, e.prefix+"."+key)
}

// Increase value
func (e Cache) Increase(key string) error {
	return cache.Increase(e.prefix + "." + key)
}

func (e Cache) Expire(key string, dur time.Duration) error {
	return cache.Expire(e.prefix+"."+key, dur)
}

func (e Cache) NewConsumer() (*redisqueue.Consumer, error) {
	return cache.NewConsumer()
}

func (e Cache) NewProducer() (*redisqueue.Producer, error) {
	return cache.NewProducer()
}

//分布式锁
func (e Cache) Lock(key string, ttl int64, options *redislock.Options) (*redislock.Lock, error) {
	return cache.Lock(e.prefix+"."+key, ttl, options)
}
