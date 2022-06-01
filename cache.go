package cache

import "time"

type Cache struct {
	KeyValue map[string]string
}

func NewCache() Cache {
	return Cache{make(map[string]string)}
}

func (c *Cache) Get(key string) (string, bool) {
	value, exists := c.KeyValue[key]
	return value, exists
}

func (c *Cache) Put(key, value string) {
	c.KeyValue[key] = value
}

func (c *Cache) Keys() []string {
	var sliceKey []string
	for key := range c.KeyValue {
		sliceKey = append(sliceKey, key)
	}
	return sliceKey
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.KeyValue[key] = value
	go func() {
		time.Sleep(deadline.Sub(time.Now()))
		delete(c.KeyValue, key)
	}()

}
