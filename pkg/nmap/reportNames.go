package nmap

type ReportName string

const (
	Vulscan ReportName = "vulscan"
	Direct  ReportName = "direct"
	Vulners ReportName = "vulners"
)
