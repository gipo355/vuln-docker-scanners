package nmap

import (
	"context"
	"slices"
	"sync"
)

// nmap -sV --script=~/vulscan.nse www.example.com

func (n *Client) ScanWithVulners(c chan<- error, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	args := slices.Concat(

		[]string{
			"-sV",                    // Version detection
			"--script=nmap-vulners/", // Script to run
		},
	)

	if n.Config.GenerateReports {
		c <- n.writeToFile(args, "vulners", Vulners)
		return
	}

	c <- n.writeToStdOut(args)
}
