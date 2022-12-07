# In-memory cache go

## Simple implementation of in-memory cache in Go



Implements a package with 4 methods:

- NewCache() - creates a new cache object
- Set() - value entry to cache by key
- Get() - getting value by key
- Delete() - deleting value by key

### Instalation 
```go
go get github.com/aleksiumish/in-memory-cache
```

### Example #1
#### Creating new cache object
```go
package main

import (
	"fmt"
	"github.com/aleksiumish/in-memory-cache"
)
func main() {
	cache := cache.NewCache()
	fmt.Println("new cache object ", cache)
}
```
### Example #2
#### 1. Value entry to cache by key
#### 2. Getting value 
#### 3. Delliting value
```go
import (
	"fmt"
	"github.com/aleksiumish/in-memory-cache"
)
func main() {
	
	// value entry to cache by key
	cache := cache.NewCache()
	cache.Set("userId", 234234)
	
	// Getting value
	userId := cache.Get("userId") 
	
	// Deleting value by key
	cache.Delete("userId")
}

```



