package nmap

import (
	"log"
	"sync"
)

func RunNmap(n *Client) {
	log.Println("Executing nmap...")

	// Testing
	nmapArgs := n.Config.Flags

	var wg sync.WaitGroup

	channels := []chan error{}

	currentIndex := 0

	if len(nmapArgs) > 0 {
		channels = append(channels, make(chan error))

		go n.DirectScan(nmapArgs, channels[currentIndex], &wg)
		currentIndex++
	}

	if n.Config.Vulscan {
		channels = append(channels, make(chan error))

		go n.ScanWithVulscan(channels[currentIndex], &wg)
		currentIndex++
	}

	if n.Config.Vulner {
		channels = append(channels, make(chan error))

		go n.ScanWithVulners(channels[currentIndex], &wg)
		currentIndex++
	}

	for i := 0; i < len(channels); i++ {
		wg.Add(1)
		go func(i int) {
			err := <-channels[i]
			if err != nil {
				log.Fatal(err)
			}
		}(i)
	}

	wg.Wait()
	for _, ch := range channels {
		close(ch)
	}

	log.Println("nmap finished")

	// parsing nmap output
	if n.Config.GenerateReports {
		if len(nmapArgs) > 0 {
			if cErr := n.ConvertToJSON(Direct); cErr != nil {
				log.Fatal(cErr)
			}
			n.GenerateSarif(Direct)
		}

		if n.Config.Vulner {
			if cErr := n.ConvertToJSON(Vulners); cErr != nil {
				log.Fatal(cErr)
			}
			n.GenerateSarif(Vulners)
		}

		if n.Config.Vulscan {
			if cErr := n.ConvertToJSON(Vulscan); cErr != nil {
				log.Fatal(cErr)
			}
			n.GenerateSarif(Vulscan)
		}
	}
}
