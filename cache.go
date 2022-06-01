package cache

import "time"

type Data struct {
	value    string
	deadline time.Time
	timesup
	bool
}

type Cache struct {
	data map[string]Data
}

func NewCache() Cache {
	return Cache{make(map[string]Data)}
}

func (c Cache) Get(key string) (string, bool) {
	res, ok := c.data[key]
	return res.value, ok
}

func (c Cache) Put(key, value string) {
	c.data[key] = Data{
		value:   value,
		timesup: false,
	}
}

func (c Cache) Keys() []string {
	res := make([]string, 0, len(c.data))
	for k, v := range c.data {
		if !v.timesup || time.Since(v.deadline).Seconds() < 0 {
			res = append(res, k)
		}
	}

	return res
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.data[key] = Data{
		value:    value,
		deadline: deadline,
		timesup:  true,
	}
}
