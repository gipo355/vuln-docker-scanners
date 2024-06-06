package nmap

import (
	"fmt"
	"log"
	"os/exec"
)

type Client struct {
	Config *Config
}

func NewNmapClient(c *Config) (*Client, error) {
	path, err := exec.LookPath("nmap")
	if err != nil {
		return nil, fmt.Errorf("nmap not found: %w", err)
	}

	log.Println(path) // bin/ls

	return &Client{
		Config: c,
	}, nil
}
