package cache2

import (
	"fmt"
	"testing"
)

func TestCacheFactory_Create(t *testing.T) {
	factory := &RedisCacheFactory{}
	redis := factory.Create()

	redis.Set("key1", "value1")
	fmt.Println(redis.Get("key1"))
}
