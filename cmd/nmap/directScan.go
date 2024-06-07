package nmap

import (
	"context"
	"sync"
)

func (n *Client) DirectScan(
	nmapArgs []string,
	c chan<- error,
	wg *sync.WaitGroup,
	ctx context.Context,
) {
	defer wg.Done()

	if n.Config.GenerateReports {
		c <- n.writeToFile(nmapArgs, "direct", Direct)
		return
	}

	c <- n.writeToStdOut(nmapArgs)
}
