package main 

import (
    "net/http"
    "sync"
    "testing"
    "time"
)

func TestConcurrentSeatBooking(t *testing.T) {
    url := "http://127.0.0.1:8080/api/showtimes/15/seats/1" 
    token := ""             // Paste your fresh token here

    totalRequests := 5
    var wg sync.WaitGroup
    wg.Add(totalRequests)

    statusCodes := make(chan int, totalRequests)
    startSignal := make(chan struct{}) 

    for i := 0; i < totalRequests; i++ {
        go func() {
            defer wg.Done()
            <-startSignal 

            client := &http.Client{Timeout: 5 * time.Second}
            req, err := http.NewRequest(http.MethodPut, url, nil)
            if err != nil {
                return
            }
            req.Header.Set("Authorization", "Bearer "+token)

            resp, err := client.Do(req)
            if err != nil {
                return
            }
            defer resp.Body.Close()

            statusCodes <- resp.StatusCode
        }()
    }

   
    time.Sleep(1 * time.Second)
    
   
    close(startSignal)

    wg.Wait()
    close(statusCodes)

    successCount := 0
    conflictCount := 0

    for code := range statusCodes {
        if code == http.StatusOK {
            successCount++
        } else if code == http.StatusConflict {
            conflictCount++
        }else{
            t.Logf(">>> UNEXPECTED STATUS CODE:%d",code)
        }
    }

    t.Logf(">>> Concurrency Test Results: Success (200) = %d, Conflicts (409) = %d", successCount, conflictCount)

    if successCount != 1 {
        t.Errorf("Race condition detected! Expected exactly 1 success, got %d", successCount)
    }
}


func TestConcurrentSeatBookingDifferentUsers(t *testing.T) {
	url := "http://localhost:8080/api/showtimes/15/seats/35" 

	
	tokens := []string{
		"", // Paste the first user fresh token here
		"", // Paste the second user fresh token here
	}

	totalRequests := len(tokens)
	var wg sync.WaitGroup
	wg.Add(totalRequests)

	statusCodes := make(chan int, totalRequests)
	startSignal := make(chan struct{})

	for i := 0; i < totalRequests; i++ {
		token := tokens[i] 
		go func() {
			defer wg.Done()
			<-startSignal 

			client := &http.Client{Timeout: 5 * time.Second}
			req, err := http.NewRequest(http.MethodPut, url, nil)
			if err != nil {
				return
			}
			req.Header.Set("Authorization", "Bearer "+token)

			resp, err := client.Do(req)
			if err != nil {
				return
			}
			defer resp.Body.Close()

			statusCodes <- resp.StatusCode
		}()
	}

	
	time.Sleep(1 * time.Second)
	
	
	close(startSignal)

	wg.Wait()
	close(statusCodes)

	successCount := 0
	conflictCount := 0

	for code := range statusCodes {
		if code == http.StatusOK {
			successCount++
		} else if code == http.StatusConflict {
			conflictCount++
		} else {
			t.Logf(">>> UNEXPECTED STATUS CODE: %d", code)
		}
	}

	t.Logf(">>> Multi-User Concurrency Results: Success (200) = %d, Conflicts (409) = %d", successCount, conflictCount)

	if successCount != 1 {
		t.Errorf("Race condition detected! Expected exactly 1 success, got %d", successCount)
	}
}