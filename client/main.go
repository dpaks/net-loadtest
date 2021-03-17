package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var cl = HTTPClient()

func run(wg *sync.WaitGroup) {
	defer wg.Done()
	//url := "http://localhost:3001"
	url := "http://172.17.0.2:3001"

	res, err := cl.Post(url, []byte("some data"))
	if err != nil {
		fmt.Println(err)
		return
	}

	hCode, _, err := cl.ParseResponse(res)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Received", hCode)
	if hCode != http.StatusOK {
		panic("hCode")
	}
}

func main() {
	n := 2
	for i := 0; i < n; i++ {
		spawn()
		time.Sleep(20 * time.Second)
	}
	select {}
}

func spawn() {
	n := 10000

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go run(&wg)
	}

	wg.Wait()
}
