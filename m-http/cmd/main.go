package main

import (
	"github.com/rikudesu/go-server/m-http/client"
)

func main() {
	c := client.NewClient("localhost:8080")
	c.Dial()
}
