package cache

type Cache struct {
	keyValue map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		keyValue: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.keyValue[key] = value
}

func (c *Cache) Get(key string) interface{} {
	return c.keyValue[key]
}

func (c *Cache) Delete(key string) {
	delete(c.keyValue, key)
}
