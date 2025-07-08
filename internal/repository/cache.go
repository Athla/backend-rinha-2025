package repository

type Cache interface {
	Add(key string, value any) error
	Get(key string) (any, error)
}

type localCache map[string]any

func NewLocalCache() Cache {
	cache := new(localCache)
	return cache
}

func (c *localCache) Add(key string, value any) error {
	return nil
}
func (c *localCache) Get(key string) (any, error) {

	return "", nil
}
