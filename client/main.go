package main

import "fmt"

func main() {
	url := "http://localhost:3001"

	cl := HTTPClient()
	res, err := cl.Get(url)
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
}
