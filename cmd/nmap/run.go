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

	if len(nmapArgs) > 0 {
		channels = append(channels, make(chan error))

		go n.DirectScan(nmapArgs, channels[0], &wg)
	}

	if n.Config.Vulscan {
		channels = append(channels, make(chan error))

		go n.ScanWithVulscan(channels[1], &wg)
	}

	if n.Config.Vulner {
		channels = append(channels, make(chan error))

		go n.ScanWithVulners(channels[2], &wg)
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
	// 	var wg sync.WaitGroup
	// for i := 0; i < len(channels); i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		select {
	// 		case err := <-channels[i]:
	// 			if err != nil {
	// 				log.Panic(fmt.Errorf("error scanning: %w", err))
	// 			}
	// 			log.Printf("scan %d finished\n", i)
	// 		}
	// 	}(i)
	// }
	// wg.Wait()

	wg.Wait()
	for _, ch := range channels {
		close(ch)
	}

	log.Println("nmap finished")

	// parsing nmap output

	if n.Config.GenerateReports {
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
