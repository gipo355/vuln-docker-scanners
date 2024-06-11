package nmap

// names enum.
type ReportName string

const (
	Vulscan ReportName = "vulscan"
	Direct  ReportName = "direct"
	Vulner  ReportName = "vulner"
)
