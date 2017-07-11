package cache

import (
	"archive/zip"
	"github.com/dubrovin/coding-challenge/storage"
)

type CachedFile struct {
	file zip.File
	ttl  int64
}

type Cache struct {
	Storage *storage.Storage
	Files   map[string]map[string]*CachedFile
	TTL     int64
}

func NewCache(stor *storage.Storage, ttl int64) *Cache {
	return &Cache{
		Storage: stor,
		TTL:     ttl,
		Files:   make(map[string]map[string]*CachedFile),
	}
}
