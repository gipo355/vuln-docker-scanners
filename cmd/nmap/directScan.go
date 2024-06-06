package nmap

import (
	"sync"
)

func (n *Client) DirectScan(nmapArgs []string, c chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	if n.Config.GenerateReports {
		c <- n.writeToFile(nmapArgs, "direct", Direct)
		return
	}

	c <- n.writeToStdOut(nmapArgs)
}
