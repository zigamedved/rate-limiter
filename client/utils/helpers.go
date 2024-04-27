package utils

import (
	"fmt"
	"os"
	"strconv"
)

func ParseCliArgs() (int, int, error) {
	if len(os.Args) != 3 {
		fmt.Println("Wrong usage!")
		return 0, 0, fmt.Errorf("\nusage: go run . <clients> <workersPerClient>")
	}

	clients, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid number of clients:", err)
		return 0, 0, err
	}

	workersPerClient, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid number of workers per client:", err)
		return 0, 0, err
	}

	return clients, workersPerClient, nil
}
