package main

import (
	"fmt"

	"github.com/rikudesu/go-server/b-http/client"
)

func main() {
	c := client.NewClient(":5000")

	resp, err := c.Hello([]byte("hello"))
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
