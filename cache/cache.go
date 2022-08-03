package cache

import "errors"

// 定义一个Cache接口作为父类
type Cache interface {
	Set(key, value string)
	Get(key string) string
}

// 实现具体的cache： RedisCache
type RedisCache struct {
	data map[string]string
}

func (rc *RedisCache) Set(key, value string) {
	rc.data[key] = value
}
func (rc *RedisCache) Get(key string) string {
	return "RedisCache: " + rc.data[key]
}

// 实现具体的cache： MemoryCache
type MemoryCache struct {
	data map[string]string
}

func (mc *MemoryCache) Set(key, value string) {
	mc.data[key] = value
}
func (mc *MemoryCache) Get(key string) string {
	return "MemoryCache: " + mc.data[key]
}

// 实现Cache的简单工厂
type CacheType int

const (
	CACHE_REDIS CacheType = iota
	CACHE_MEMORY
)

type CacheFactory struct {
}

func (cf *CacheFactory) Create(ct CacheType) (Cache, error) {
	if ct == CACHE_REDIS {
		return &RedisCache{
			data: map[string]string{},
		}, nil
	}

	if ct == CACHE_MEMORY {
		return &MemoryCache{
			data: map[string]string{},
		}, nil
	}

	return nil, errors.New("CacheType error")
}
