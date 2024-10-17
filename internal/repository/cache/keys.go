package cache

import "fmt"

const (
	FindUsersCacheKey = "users:find"
)

func FindUserByIDCacheKey(id string) string {
	return fmt.Sprintf("users:find:%s", id)
}
