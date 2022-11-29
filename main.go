package main

import (
	"fmt"
	"go-redis-cache/redisCache/providers"
	"time"
)

func main() {
	cacheObject := providers.NewCacheClient()

	// setting value in the cache
	err := cacheObject.Set("emp_details", `{"name":"Younus", "age":32, "designation":"GoLang Developer"}`, time.Duration(time.Now().Add(time.Minute*60).Unix()))
	if err != nil {
		return
	}

	// getting value form the cache
	value := cacheObject.Get("emp_details")
	fmt.Printf("value : %v\n", value)

	// deleting value from cache
	_ = cacheObject.Delete("emp_details")
	value = cacheObject.Get("emp_details")
	fmt.Printf("value : %v", value)
}
