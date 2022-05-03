package storage

import "appslab-ke/kipchoge-go/internal/platform"

type Cache interface {
	Set(key, data string) error
	Ping()
}

func NewCache() (Cache, error) {
	cache, err := platform.NewCache()
	if err != nil {
		return nil, err
	}

	return cache, nil
}
