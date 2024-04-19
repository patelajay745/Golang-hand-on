package main

import (
	"fmt"
	"time"
)

func fetchDataFromServer(serverUrl string) {
	fmt.Printf("Fetching data from server: %s \n", serverUrl)
}

func updateDashboard(data map[string]interface{}) {
	fmt.Println("updating dashboard with fetched data:", data)

}
func main() {
	serverUrl := []string{"http://server1.com", "http://server2.com", "http://server3.com"}

	ticker := time.NewTicker(5 * time.Second) //fetch data every 5 seconds

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				fetchedData := make(map[string]interface{})
				for _, url := range serverUrl {
					fetchDataFromServer(url) //use models to populate data into it.
				}
				updateDashboard(fetchedData)

			}
		}
	}()

	time.Sleep(30 * time.Second)
	ticker.Stop()

	done <- true
	fmt.Println("Ticker stopped")

}
