package service

import (
	"github.com/patrickmn/go-cache"
)

type cacheService struct {
	cache *cache.Cache
}

func NewCacheService(c *cache.Cache) *cacheService {
	return &cacheService{cache: c}
}

func (s *cacheService) StoreCache(key string, value interface{}) {
	s.cache.Set(key, value, cache.DefaultExpiration)
}

func (s *cacheService) GetCache(key string) interface{} {
	val, found := s.cache.Get(key)
	if found {
		return val
	}
	return nil
}
