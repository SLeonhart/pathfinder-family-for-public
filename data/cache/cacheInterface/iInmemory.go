package cacheInterface

//go:generate mockgen -source=iInMemory.go -destination=../cacheMock/inmemory.go -package=cacheMock

type IInMemory interface {
	Set(key string, item interface{})
	Get(key string) interface{}
	Delete(key string)
}
