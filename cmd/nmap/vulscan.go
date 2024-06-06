package nmap

import (
	"slices"
	"sync"
)

// nmap -sV --script=~/vulscan.nse www.example.com

func (n *Client) ScanWithVulscan(c chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	args := slices.Concat(
		[]string{
			"-sV",               // Version detection
			"--script=vulscan/", // Script to run
		},
	)

	if n.Config.GenerateReports {
		c <- n.writeToFile(args, "vulscan", Vulscan)
		return
	}

	c <- n.writeToStdOut(args)
}
