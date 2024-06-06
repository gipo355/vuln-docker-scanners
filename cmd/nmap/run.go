package nmap

import (
	"fmt"
	"log"
	"sync"
)

func RunNmap(n *Client) {
	log.Println("Executing nmap...")

	// Testing
	nmapArgs := n.Config.Args

	var wg sync.WaitGroup

	channels := []chan error{}

	if len(nmapArgs) > 0 {
		channels = append(channels, make(chan error))
		wg.Add(1)

		go n.DirectScan(nmapArgs, channels[0], &wg)
	}

	if n.Config.Vulscan {
		channels = append(channels, make(chan error))

		wg.Add(1)

		go n.ScanWithVulscan(channels[1], &wg)
	}

	if n.Config.Vulner {
		channels = append(channels, make(chan error))

		wg.Add(1)

		go n.ScanWithVulners(channels[2], &wg)
	}

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

	if n.Config.GenerateReports && n.Config.GenerateSarif {
		log.Println("Generating reports...")

		if len(nmapArgs) > 0 {
			if cErr := n.ConvertToJSON(Direct); cErr != nil {
				log.Fatal(cErr)
			}
			n.GenerateSarif(Direct)
		}

		if n.Config.Vulscan {
			if cErr := n.ConvertToJSON(Vulners); cErr != nil {
				log.Fatal(cErr)
			}
			n.GenerateSarif(Vulners)
		}

		if n.Config.Vulner {
			if cErr := n.ConvertToJSON(Vulscan); cErr != nil {
				log.Fatal(cErr)
			}
			n.GenerateSarif(Vulscan)
		}
	}
}
