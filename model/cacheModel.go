package model

import "fmt"

type CacheType int

const (
	AuthorizationType CacheType = iota
)

func (d CacheType) String() string {
	return [...]string{"AuthorizationType"}[d]
}

type CacheKey interface {
	IsCache()
	String() string
}

type AuthCacheKey struct {
	Token string
}

func (AuthCacheKey) IsCache() {}

func (key AuthCacheKey) String() string {
	return fmt.Sprintf("Auth:%v", key.Token)
}
