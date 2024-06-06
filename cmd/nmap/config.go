package nmap

type Config struct {
	Target          string
	Port            string
	OutputDir       string
	Flags           []string
	Args            []string
	GenerateReports bool
	Vulner          bool
	Vulscan         bool
}

func NewConfig(writeToFile, sarif bool, target, outputDir, port string) *Config {
	return &Config{
		Target:          target,
		Port:            port,
		GenerateReports: writeToFile,
		OutputDir:       outputDir,
	}
}
