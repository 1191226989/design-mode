package cache

import (
	"fmt"
	"testing"
)

func TestCache_Create(t *testing.T) {
	redisCache := &RedisCache{
		data: map[string]string{},
	}
	redisCache.Set("key1", "value1")
	fmt.Println(redisCache.Get("key1"))

	memoryCache := &MemoryCache{
		data: map[string]string{},
	}
	memoryCache.Set("key2", "value2")
	fmt.Println(memoryCache.Get("key2"))
}

func TestCacheFactory_Create(t *testing.T) {
	cacheFactory := &CacheFactory{}

	redis, err := cacheFactory.Create(CACHE_REDIS)
	if err != nil {
		fmt.Println(err)
	}
	redis.Set("key1", "value1")
	fmt.Println(redis.Get("key1"))

	memory, err := cacheFactory.Create(CACHE_MEMORY)
	if err != nil {
		fmt.Println(err)
	}
	memory.Set("key2", "value2")
	fmt.Println(memory.Get("key2"))
}
