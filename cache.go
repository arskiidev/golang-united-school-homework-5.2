package cache

import "time"

type Node struct {
	key, value string
	deadline   time.Time
}

type Cache struct {
	cache []Node
}

func NewCache() Cache {
	return Cache{}

}

func (c *Cache) Put(key, value string) {
	for i, n := range c.cache {
		if key == n.key {
			c.cache[i].value = value
			return
		}
	}
	c.cache = append(c.cache, Node{key: key, value: value})
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	for i, n := range c.cache {
		if key == n.key {
			c.cache[i].value = value
			c.cache[i].deadline = deadline
			return
		}
	}
	c.cache = append(c.cache, Node{key: key, value: value, deadline: deadline})
}

func (c *Cache) Get(key string) (string, bool) {
	for i, n := range c.cache {
		if key == n.key {
			if !n.deadline.IsZero() && !time.Now().Before(n.deadline) {
				c.cache = append(c.cache[:i], c.cache[i+1:]...)
				return "", false
			}
			return n.value, true
		}
	}
	return "", false
}

func (c *Cache) Keys() (s []string) {
	for i, n := range c.cache {
		if !n.deadline.IsZero() && !time.Now().Before(n.deadline) {
			c.cache = append(c.cache[:i], c.cache[i+1:]...)
			continue
		}
		s = append(s, n.key)
	}
	return s
}
