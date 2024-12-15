package telemetry

import (
	"fmt"
	"os"
	"path"
)

type FFDB struct {
	filename string
	store    map[string][]byte
}

func NewFFDB() *FFDB {
	filename := GetCacheFileName()

	return &FFDB{
		filename: filename,
		store:    make(map[string][]byte),
	}
}

func (f *FFDB) RegisterFlag(key string, available bool, required bool, metadata map[string]interface{}, group string) (*FeatureFlag, error) {
	flag := &FeatureFlag{
		Key:       key,
		Available: available,
		OptedIn:   false,
		Metadata:  metadata,
		Group:     group,
	}

	return flag, nil
}

func (f *FFDB) Get(key string) ([]byte, error) {
	if val, ok := f.store[key]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("key not found")
}

func (f *FFDB) Set(key string, val []byte) error {
	f.store[key] = val
	return nil
}

func (f *FFDB) Delete(key string) {
	delete(f.store, key)
}

func (f *FFDB) Close() {
	f.store = nil
}

func (f *FFDB) Save() error {
	return nil
}

func GetCacheFileName() string {
	cacheDir, err := os.UserCacheDir()

	if err != nil {
		fmt.Println(err)
		panic("Failed to get cache directory")
	}

	appCacheDir := path.Join(cacheDir, "NextLaunch")

	appCacheFilePath := path.Join(appCacheDir, "cache.db")
	return appCacheFilePath
}
