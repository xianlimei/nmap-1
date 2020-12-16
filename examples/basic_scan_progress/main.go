package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ullaakut/nmap/v2"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	scanner, err := nmap.NewScanner(
		nmap.WithTargets("localhost"),
		nmap.WithPorts("1-4000"),
		nmap.WithServiceInfo(),
		nmap.WithVerbosity(3),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	progress := make(chan float32, 1)

	// Function to listen and print the progress
	go func() {
		for p := range progress {
			fmt.Printf("Progress: %v %%\n", p)
		}
		fmt.Println("Exit")
	}()

	result, _, err := scanner.RunWithProgress(progress)
	if err != nil {
		log.Printf("unable to run nmap scan: %v", err)
	}

	if err == nil {
		fmt.Printf("Nmap done: %d hosts up scanned in %.2f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)
	}

	//time.Sleep(time.Second*2)

	fmt.Println("-----Exit")
}
