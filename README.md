# In-memory cache go

### Simple implementation of in-memory cache in Go

Implements a package with three methods:

- NewCache() - creates a new cache object

### Example #1
```go
package main

import (
	"fmt"
	"github.com/aleksiumish/in-memory-cache"
)
func main() {
	cache := cache.NewCache()
	fmt.Println("new object cache ", cache)
}
```


