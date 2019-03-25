package main

import (
	"fmt"

	"./mock"
	"./real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://live.gushidaoshi.com/166/")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake imooc.com"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
