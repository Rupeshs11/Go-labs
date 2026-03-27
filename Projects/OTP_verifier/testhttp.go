package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	start := time.Now()
	fmt.Println("Connecting to Twilio...")
	resp, err := client.Get("https://verify.twilio.com/")
	if err != nil {
		fmt.Printf("Error: %v (took %v)\n", err, time.Since(start))
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Success! Status: %s (took %v)\n", resp.Status, time.Since(start))
}
