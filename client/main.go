package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	u "github.com/zigamedved/rate-limiter/client/utils"
)

const baseUrl = "http://localhost:8080/?clientId="

func runWorker(clientId int, workerId int, done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Printf("Client %d - Worker %d: Shutting down...\n", clientId, workerId)
			return
		default:
			url := fmt.Sprintf("%s%d", baseUrl, clientId)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Client %d - Error sending request: %v\n", clientId, err)
				return
			}
			fmt.Printf("Client %d - Worker %d: %s\n", clientId, workerId, resp.Status)
			time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		}
	}
}

func createAndRunClients(clients, workersPerClient int, done chan struct{}, wg *sync.WaitGroup) {
	for i := 1; i <= clients; i++ {
		clientId := i
		for j := 1; j <= workersPerClient; j++ {
			workerId := j
			wg.Add(1)
			go runWorker(clientId, workerId, done, wg)
		}
	}
}

func main() {

	clients, workersPerClient, err := u.ParseCliArgs()
	if err != nil {
		fmt.Println("Error while parsing CLI args", err)
		os.Exit(1)
	}
	fmt.Printf("Entered CLI arguments clients: %v workers per client: %v\n", clients, workersPerClient)

	var wg sync.WaitGroup
	var done = make(chan struct{})
	createAndRunClients(clients, workersPerClient, done, &wg)

	fmt.Println("Press any key to stop program...")
	fmt.Scanln()

	close(done)
	wg.Wait()

	fmt.Println("Program successfully stopped.")

	os.Exit(0)
}
