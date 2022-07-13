package cache

import (
	"github.com/karlseguin/ccache/v2"
	"time"
)

/*
	Cache implements an in-memory cache. Some considerations: I've implemented this type of cache to fulfill the
	exercise requirements and taking into account that I've deployed the application in a single VM environment.
	If we want to scale the application horizontally, we have to implement a distributed cache or a mechanism to
	synchronize the VMs (like a cron job)
*/
type Cache struct {
	cache *ccache.Cache
}

func NewCache() *Cache {
	return &Cache{
		cache: ccache.New(
			ccache.Configure().MaxSize(100000)),
	}
}

func (c *Cache) Get(key string) interface{} {

	item := c.cache.Get(key)

	if item == nil {
		return nil
	}

	return item.Value()
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.cache.Set(key, value, ttl)
}

func (c *Cache) Update(key string, value interface{}, ttl time.Duration) {
	c.cache.Delete(key)
	c.Set(key, value, ttl)
}
