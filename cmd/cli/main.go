package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gipo355/vuln-docker-scanners/pkg/nmap"
	"github.com/gipo355/vuln-docker-scanners/pkg/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Printf("No args provided")
	} else {
		for _, arg := range args {
			log.Printf("arg: %v", arg)
		}
	}

	log.Println("print pwd")
	utils.PrintPwd()

	// TODO: nmap must move to its own package
	// nmap section

	log.Println("Executing nmap...")

	n, err := nmap.NewNmapClient(
		&nmap.Config{
			Target:          "localhost",
			Port:            "80",
			GenerateReports: true,
			GenerateSarif:   true,
			OutputDir:       "nmap-reports",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Testing
	// TODO: remove hardcoded args
	nmapArgs := []string{"-sP"}

	var wg sync.WaitGroup

	channels := []chan error{
		make(chan error),
		make(chan error),
		make(chan error),
	}
	for range channels {
		wg.Add(1)
	}

	// directChan := make(chan error)
	go n.DirectScan(nmapArgs, channels[0], &wg)

	// vulscanChan := make(chan error)
	go n.ScanWithVulscan(channels[1], &wg)

	// vulnerChan := make(chan error)
	go n.ScanWithVulners(channels[2], &wg)

	// for i := 0; i < 3; i++ {
	for i := 0; i < len(channels); i++ {
		select {
		// case directErr := <-directChan:
		case directErr := <-channels[0]:
			if directErr != nil {
				log.Panic(fmt.Errorf("error direct scanning: %w", directErr))
			}
			log.Println("direct scan finished")

		case vulnerErr := <-channels[1]:
			if vulnerErr != nil {
				log.Panic(fmt.Errorf("error scanning with vulners: %w", vulnerErr))
			}
			log.Println("vulners scan finished")

		case vulscanErr := <-channels[2]:
			if vulscanErr != nil {
				log.Panic(fmt.Errorf("error scanning with vulscan: %w", vulscanErr))
			}
			log.Println("vulscan scan finished")
		}
	}
	wg.Wait()
	for _, ch := range channels {
		close(ch)
	}

	log.Println("nmap finished")

	// parsing nmap output

	if cErr := n.ConvertToJSON(nmap.Direct); cErr != nil {
		log.Fatal(cErr)
	}

	if cErr := n.ConvertToJSON(nmap.Vulners); cErr != nil {
		log.Fatal(cErr)
	}

	if cErr := n.ConvertToJSON(nmap.Vulscan); cErr != nil {
		log.Fatal(cErr)
	}

	n.GenerateSarifReport(nmap.Vulners)
	n.GenerateSarifReport(nmap.Direct)
	n.GenerateSarifReport(nmap.Vulscan)
}
