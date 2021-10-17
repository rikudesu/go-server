package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func handler(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	request, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
	res := http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 0,
		Body: ioutil.NopCloser(
			strings.NewReader("Hello World\n")),
	}
	res.Write(conn)
}

const PORT = "localhost:8080"

func main() {
	l, e := net.Listen("tcp", PORT)
	if e != nil {
		log.Fatal(e)
		return
	}
	fmt.Printf("Server is running at %s", PORT)

	for {
		conn, e := l.Accept()
		if e != nil {
			log.Fatal(e)
			return
		}
		go handler(conn)
	}
}
