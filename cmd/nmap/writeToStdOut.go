package nmap

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
)

func (n *Client) writeToStdOut(nmapArgs []string) error {
	target := n.Config.Target

	args := slices.Concat(nmapArgs, []string{
		target, // target
	})

	cmd := exec.Command("nmap", args...)

	log.Printf("cmd: %v", cmd)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := fmt.Errorf("nmap: %w", cmd.Run())

	return err
}
