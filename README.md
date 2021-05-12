# go insect
# devops gateway go etcd client

```go
package main

import (
go_insect"github.com/yametech/go-insect"
)

func main() {
	go_insect.GlobalEtcdAddress = "http://127.0.0.1:2379" // etcd host
	// go-insect.GlobalEtcdTTL = 60                          // ttl
	// go-insect.INSECT_SERVER_URL = "0.0.0.0"               // register server host
	go_insect.INSECT_SERVER_PORT = 8080
	go_insect.INSECT_SERVER_NAME = "goserver" // register server name
	go go_insect.EtcdProxy()

	for {
	}
}
```

### etcdctl --endpoints=0.0.0.0:2379 --prefix=true get goserver
```
goserver_31ac1d93-1fa4-4b7a-b9da-55d7da6f4ecd
0.0.0.0:8080
```
