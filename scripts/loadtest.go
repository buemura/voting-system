package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	totalRequests   = 10000
	concurrentUsers = 20
	apiEndpoint     = "http://localhost:9999/vote"
	contentType     = "application/json"
)

var candidateIDs = []string{
	"cjd7n3v6g0001gq9g7j2m3pbk",
	"cjd7n3v6g0002gq9g7j2m3pbk",
	"cjd7n3v6g0003gq9g7j2m3pbk",
	"cjd7n3v6g0004gq9g7j2m3pbk",
	"cjd7n3v6g0005gq9g7j2m3pbk",
	"clzc1pqd0000008mnfmkq9r50",
}

func main() {
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrentUsers)

	startTime := time.Now() // Start timer

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func(i int) {
			defer wg.Done()
			sendVoteRequest(candidateIDs[i%len(candidateIDs)])
			<-sem
		}(i)
	}

	wg.Wait()
	endTime := time.Now() // End timer

	totalTime := endTime.Sub(startTime).Seconds()
	fmt.Printf("Load test completed in %.2f seconds\n", totalTime)
	fmt.Printf("Requests per second: %.2f\n", float64(totalRequests)/totalTime)
}

func sendVoteRequest(candidateID string) {
	payload := map[string]string{"candidate_id": candidateID}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Failed to marshal JSON: %v\n", err)
		return
	}

	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Failed to create request: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status: %s\n", resp.Status)
	}
}
