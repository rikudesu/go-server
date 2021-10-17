package client

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client {
	return &Client{addr: addr}
}

func (c *Client) Dial() {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s", c.addr), nil)
	if err != nil {
		panic(err)
	}

	req.Write(conn)
	res, err := http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}
