/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package nmap

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/gipo355/vuln-docker-scanners/pkg/utils"
)

// must replicate the following command:
// `docker run --network=host --workdir=/app --volume .:/app gipo355/vuln-docker-scanners nmap --vulner --vulscan --target=localhost --port=80 --generate-reports --generate-sarif`

// reference on how to prevent globals
// https://github.com/vmware-tanzu/sonobuoy/blob/main/cmd/sonobuoy/app/delete.go#L38-L58

type nmapFlags struct {
	Target          string
	Port            string
	OutputDir       string
	GenerateReports bool
	GenerateSarif   bool
	Vulner          bool
	Vulscan         bool
}

// nmapCmd represents the nmap command
func NewCmdNmap() *cobra.Command {
	f := nmapFlags{}

	cmd := &cobra.Command{
		Use:   "nmap",
		Short: "Runs nmap against a target",
		Long:  `Runs map against a target.`,
		Run:   nmapRun(&f),
	}

	utils.AddFlag(
		&f.Target,
		cmd.Flags(),
		"target",
		"t",
		"localhost",
		"Target to scan",
	)
	if err := utils.MarkFlagRequired(cmd, "target"); err != nil {
		log.Println(err)
	}

	utils.AddPersistentFlag(
		&f.OutputDir,
		cmd.PersistentFlags(),
		"output-dir",
		"o",
		"nmap-reports",
		"Output directory for reports",
	)
	utils.AddFlag(
		&f.Port,
		cmd.Flags(),
		"port",
		"p",
		"",
		"Port to scan",
	)

	return cmd
}

// func deleteSonobuoyRun(f *deleteFlags) func(cmd *cobra.Command, args []string) {

func nmapRun(f *nmapFlags) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		n, err := NewNmapClient(
			&Config{
				Target:          f.Target,
				Port:            f.Port,
				GenerateReports: f.GenerateReports,
				GenerateSarif:   f.GenerateSarif,
				OutputDir:       f.OutputDir,
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		RunNmap(n)
	}
}
