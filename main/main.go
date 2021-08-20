package main

import (
	"fmt"
	"math/rand"

	"github.com/njayp/proto/client"
	"github.com/njayp/proto/server"
)

func full() {
	go server.Start()
	fmt.Scanln()

	println("sending the blob")
	client.SendBytes(blob())
	fmt.Scanln()
}

func blob() []byte {
	blob := make([]byte, 1024*1024*128)
	rand.Read(blob)
	return blob
}

func main() {
	full()
}
