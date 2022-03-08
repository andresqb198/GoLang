package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func request(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[%d] %s\n", res.StatusCode, url)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: go run main.go <url1> <url2> ... <urln>")
	}
	for _, url := range os.Args[1:] {
		request("https://" + url)
	}
}
