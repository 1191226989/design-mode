package cache2

// 定义一个Cache接口作为父类
type Cache interface {
	Set(key, value string)
	Get(key string) string
}

// 实现具体的cache： RedisCache
type RedisCache struct {
	data map[string]string
}

func NewRedisCache() Cache {
	return &RedisCache{
		data: map[string]string{},
	}
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

func NewMemoryCache() Cache {
	return &MemoryCache{
		data: map[string]string{},
	}
}

func (mc *MemoryCache) Set(key, value string) {
	mc.data[key] = value
}
func (mc *MemoryCache) Get(key string) string {
	return "MemoryCache: " + mc.data[key]
}

// 实现Cache的抽象工厂
type CacheFactory interface {
	Create() Cache
}

type RedisCacheFactory struct {
}

func (rcf *RedisCacheFactory) Create() Cache {
	return NewRedisCache()
}

type MemoryCacheFactory struct {
}

func (mcf *MemoryCacheFactory) Create() Cache {
	return NewMemoryCache()
}
