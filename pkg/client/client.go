package client

type Client interface {
	DirectScan() error
	CreteSarif() error
	VulnScan() error
}
