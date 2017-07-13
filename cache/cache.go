package cache

import (
	"errors"
	"fmt"
	"github.com/dubrovin/gotest/storage"
	"sync"
	"time"
)

type CachedFile struct {
	file map[string][]byte
	ttl  int64
}

type Cache struct {
	Storage *storage.Storage
	Files   map[string]*CachedFile
	TTL     time.Duration
	mu      sync.RWMutex
}

func NewCache(stor *storage.Storage, ttl time.Duration) *Cache {
	return &Cache{
		Storage: stor,
		TTL:     ttl,
		Files:   make(map[string]*CachedFile),
	}
}

func (c *Cache) GetZipFile(ZipFile string) ([]byte, error) {
	// целый зип файл всегда достаем из хранилища
	c.mu.RLock()
	if zf, ok := c.Storage.Files[ZipFile]; ok {
		return zf.Bytes()
	}
	c.mu.RUnlock()
	return nil, errors.New(fmt.Sprintf("Zip file with name = %s is not found", ZipFile))
}

func (c *Cache) loadFile(ZipFile string) (bool, error) {
	// если файлы уже загружены, то ничего не делаем
	if _, ok := c.Files[ZipFile]; ok {
		return ok, nil
	}

	// иначе вычитываем все файлы, которые лежат в хранилище
	files, err := c.Storage.Files[ZipFile].Read()
	if err != nil {
		return false, err
	}

	// заполняем кэш и время жизни
	c.Files[ZipFile] = &CachedFile{
		file: files,
		ttl:  time.Now().Add(c.TTL).UnixNano(),
	}
	return true, nil
}

func (c *Cache) GetFiles(ZipFile string, Files []string) (map[string][]byte, error) {
	c.mu.RLock()
	// если файлов нет, то загружаем
	if _, ok := c.Files[ZipFile]; !ok {
		if ok, err := c.loadFile(ZipFile); !ok {
			return nil, err
		}
	}

	// обновляем время жизни
	c.Files[ZipFile].ttl = time.Now().Add(c.TTL).UnixNano()
	c.mu.RUnlock()
	return c.Files[ZipFile].file, nil
}

func (c *Cache) Checker() {
	for {
		time.Sleep(c.TTL)
		c.mu.Lock()
		for k := range c.Files {
			if c.Files[k].ttl < time.Now().UnixNano() {
				delete(c.Files, k)
			}
		}
		c.mu.Unlock()

	}

}

func (c *Cache) AddZipFile(path string) error{
	c.mu.Lock()
	zp, err := storage.NewZipFile(path)
	if err != nil {
		return err
	}
	c.Storage.Add(zp)
	c.mu.Unlock()
	return nil
}

func (c *Cache) GetAllZipFiles() map[string]*storage.ZipFile{
	c.mu.RLock()
	files := c.Storage.Files
	c.mu.RUnlock()
	return files
}