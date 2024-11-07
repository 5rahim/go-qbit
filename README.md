# go-qbit

Wrapper for [qBittorrent Web API](https://github.com/qbittorrent/qBittorrent/wiki/#login) (> v3.1.x)

Forked from [KnutZuidema/go-qbittorrent](https://github.com/KnutZuidema/go-qbittorrent)


### Example

```go
package main

import (
	"fmt"
	qbittorrent "github.com/5rahim/go-qbit"
)

func main() {
	client := qbittorrent.NewClient(&qbittorrent.NewClientOptions{
		Username: "",
		Password: "",
		Port:     8080,
		Host:     "127.0.0.1",
		BinaryPath: "C:/Program Files/qBittorrent/qbittorrent.exe",
	})
	
	err := client.Login()
	fmt.Printf("failed to login: %v\n", err)
	
	err = client.Start()
	fmt.Printf("failed to start qBittorrent: %v\n", err)
}
```


### qBittorrent settings

1. Go to `Options > Web UI`
2. Check the box for `Web User Interface (Remote Control)`
