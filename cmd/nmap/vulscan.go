package nmap

import (
	"context"
	"slices"
	"sync"
)

// nmap -sV --script=~/vulscan.nse www.example.com

func (n *Client) ScanWithVulscan(c chan<- error, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	args := slices.Concat(
		[]string{
			"-sV",               // Version detection
			"--script=vulscan/", // Script to run
		},
	)

	if n.Config.GenerateReports {
		c <- n.writeToFile(args, string(Vulscan), Vulscan)
		return
	}

	c <- n.writeToStdOut(args)
}
