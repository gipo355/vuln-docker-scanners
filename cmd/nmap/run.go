package nmap

import (
	"context"
	"log"
	"sync"
)

func Run(n *Client) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("Executing nmap...")

	// Testing
	nmapArgs := n.Config.Flags

	var wg sync.WaitGroup

	channels := []chan error{}

	currentIndex := 0

	if len(nmapArgs) > 0 {
		channels = append(channels, make(chan error))

		go n.DirectScan(nmapArgs, channels[currentIndex], &wg, ctx)
		currentIndex++
	}

	if n.Config.Vulscan {
		channels = append(channels, make(chan error))

		go n.ScanWithVulscan(channels[currentIndex], &wg, ctx)
		currentIndex++
	}

	if n.Config.Vulner {
		channels = append(channels, make(chan error))

		go n.ScanWithVulners(channels[currentIndex], &wg, ctx)
		currentIndex++
	}

	for i := 0; i < len(channels); i++ {

		wg.Add(1)

		go func(i int, ctx context.Context) {
			select {
			case <-ctx.Done():
				return
			default:
				err := <-channels[i]
				if err != nil {
					log.Fatal(err)
				}
			}
		}(i, ctx)
	}

	wg.Wait()
	for _, ch := range channels {
		close(ch)
	}

	log.Println("nmap finished")

	// parsing nmap output
	if n.Config.GenerateReports {
		if len(nmapArgs) > 0 {
			if cErr := n.CreateJSONReport(Direct); cErr != nil {
				log.Fatal(cErr)
			}
			// n.GenerateSarif(Direct)
		}

		if n.Config.Vulner {
			if cErr := n.CreateJSONReport(Vulners); cErr != nil {
				log.Fatal(cErr)
			}
			// n.GenerateSarif(Vulners)
		}

		if n.Config.Vulscan {
			if cErr := n.CreateJSONReport(Vulscan); cErr != nil {
				log.Fatal(cErr)
			}
			// n.GenerateSarif(Vulscan)
		}
	}
}
