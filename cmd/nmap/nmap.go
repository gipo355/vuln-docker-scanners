/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package nmap

import (
	"log"

	"github.com/spf13/cobra"
)

// nmapCmd represents the nmap command
var NmapCmd = &cobra.Command{
	Use:   "nmap",
	Short: "Runs nmap against a target",
	Long:  `Runs map against a target.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}

		n, err := NewNmapClient(
			&Config{
				Target:          "localhost",
				Port:            "80",
				GenerateReports: true,
				GenerateSarif:   true,
				OutputDir:       "nmap-reports",
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		RunNmap(n)
	},
}

// must replicate the following command:
// `docker run --network=host --workdir=/app --volume .:/app gipo355/vuln-docker-scanners nmap --vulner --vulscan --target=localhost --port=80 --generate-reports --generate-sarif`

// reference on how to prevent globals
// https://github.com/vmware-tanzu/sonobuoy/blob/main/cmd/sonobuoy/app/delete.go#L38-L58

type NmapFlags struct {
	Target          string
	Port            string
	GenerateReports bool
	GenerateSarif   bool
	Vulner          bool
	Vulscan         bool
	OutputDir       string
}

var (
	target          string
	port            string
	generateReports bool
	generateSarif   bool
	vulner          bool
	vulscan         bool
	outputDir       string
)

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nmapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nmapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// target
	NmapCmd.Flags().StringVarP(&target, "target", "t", "localhost", "Target to scan")
	// if err := NmapCmd.MarkFlagRequired("target"); err != nil {
	// 	log.Println(err)
	// }

	// port
	NmapCmd.Flags().StringVarP(&port, "port", "p", "", "Port to scan")

	// generate-reports
	NmapCmd.Flags().BoolVarP(&generateReports, "generate-reports", "r", true, "Generate reports")

	// generate-sarif
	NmapCmd.Flags().BoolVarP(&generateSarif, "generate-sarif", "s", true, "Generate sarif reports")

	// vulner
	NmapCmd.Flags().
		BoolVarP(&vulner, "vulner", "v", false, "Scan for vulnerabilities using vulner scripts")

	// vulscan
	NmapCmd.Flags().
		BoolVarP(&vulscan, "vulscan", "V", false, "Scan for vulnerabilities using vulscan scripts")

	// output-dir
	NmapCmd.Flags().
		StringVarP(&outputDir, "output-dir", "o", "nmap-reports", "Output directory for reports")
}
