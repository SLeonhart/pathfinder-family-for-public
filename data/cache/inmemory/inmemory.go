package inmemory

import (
	"pathfinder-family/config"

	"time"

	gocache "github.com/patrickmn/go-cache"
)

type InMemory struct {
	client         *gocache.Cache
	expirationMsec int64
}

func NewInMemory(cfg *config.Config) *InMemory {
	return &InMemory{
		expirationMsec: cfg.Cache.InMemory.ExpirationMsec,
		client:         gocache.New(5*time.Minute, 10*time.Minute),
	}
}

func (c *InMemory) Set(key string, item interface{}) {
	if c.expirationMsec == 0 {
		return
	}
	c.client.Set(key, item, time.Duration(c.expirationMsec)*time.Second)
}

func (c *InMemory) Delete(key string) {
	c.client.Delete(key)
}

func (c *InMemory) Get(key string) interface{} {
	if c.expirationMsec == 0 {
		return nil
	}
	if cached, found := c.client.Get(key); found {
		return cached
	}
	return nil
}
