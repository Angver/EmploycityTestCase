package client

//go:generate mockgen -source=client.go -destination=./client_mock.go -package=client

import "github.com/bradfitz/gomemcache/memcache"

type MemcachedClient interface {
	Get(key string) (item *memcache.Item, err error)
	Set(item *memcache.Item) error
	Delete(key string) error
}
