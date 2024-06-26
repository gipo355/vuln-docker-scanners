package nmap

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gipo355/vuln-docker-scanners/pkg/sarif"
)

// TODO: must implement

// TODO: must create a generate sarif for direct, vulner and vulscan reports
// separately since they differ in outputs

type Report struct {
	Version string `json:"Version"`
	Host    []struct {
		Port []struct {
			Service struct {
				Name    string `json:"Name"`
				Product string `json:"Product"`
				Version string `json:"Version"`
			} `json:"Service"`
			Protocol string `json:"Protocol"`
			PortID   int    `json:"PortID"`
		} `json:"Port"`
		HostAddress []struct {
			Address string `json:"Address"`
		} `json:"HostAddress"`
	} `json:"Host"`
}

// GenerateSarif generates a SARIF report from the nmap output xml.
// func (n *Client) GenerateSarif() error {
// 	return nil
// }

func (n *Client) GenerateSarif(name ReportName) {
	mainDir := n.Config.OutputDir
	fileInput := mainDir + "/" + string(name) + "/" + string(name) + "-report.json"
	fileOutput := mainDir + "/" + string(name) + "/" + string(name) + "-report.sarif"

	// Load the Nmap JSON report
	nmapReportBytes, _ := os.ReadFile(fileInput)
	var nmapReport Report
	json.Unmarshal(nmapReportBytes, &nmapReport)

	// Initialize the SARIF report
	sarifReport := sarif.SarifReport2{
		Schema:  "https://schemastore.azurewebsites.net/schemas/json/sarif-2.1.0-rtm.5.json",
		Version: "2.1.0",
	}

	// Convert each host in the Nmap report to a SARIF run
	for _, host := range nmapReport.Host {
		run := struct {
			Tool struct {
				Driver struct {
					Name    string `json:"name"`
					Version string `json:"version"`
					Rules   []struct {
						ID              string `json:"id"`
						Name            string `json:"name"`
						FullDescription struct {
							Text string `json:"text"`
						} `json:"fullDescription"`
						HelpURI string `json:"helpUri"`
					} `json:"rules"`
				} `json:"driver"`
			} `json:"tool"`
			Results []struct {
				RuleID  string `json:"ruleId"`
				Level   string `json:"level"`
				Message struct {
					Text string `json:"text"`
				} `json:"message"`
				Locations []struct {
					PhysicalLocation struct {
						Address struct {
							AbsoluteAddress string `json:"absoluteAddress"`
						} `json:"address"`
					} `json:"physicalLocation"`
				} `json:"locations"`
			} `json:"results"`
		}{
			Tool: struct {
				Driver struct {
					Name    string `json:"name"`
					Version string `json:"version"`
					Rules   []struct {
						ID              string `json:"id"`
						Name            string `json:"name"`
						FullDescription struct {
							Text string `json:"text"`
						} `json:"fullDescription"`
						HelpURI string `json:"helpUri"`
					} `json:"rules"`
				} `json:"driver"`
			}{
				Driver: struct {
					Name    string `json:"name"`
					Version string `json:"version"`
					Rules   []struct {
						ID              string `json:"id"`
						Name            string `json:"name"`
						FullDescription struct {
							Text string `json:"text"`
						} `json:"fullDescription"`
						HelpURI string `json:"helpUri"`
					} `json:"rules"`
				}{
					Name:    "Nmap",
					Version: nmapReport.Version,
				},
			},
		}

		// Convert each port in the host to a SARIF rule and result
		for _, port := range host.Port {
			rule := struct {
				ID              string `json:"id"`
				Name            string `json:"name"`
				FullDescription struct {
					Text string `json:"text"`
				} `json:"fullDescription"`
				HelpURI string `json:"helpUri"`
			}{
				ID:   port.Protocol + "/" + string(port.PortID),
				Name: port.Service.Name,
				FullDescription: struct {
					Text string `json:"text"`
				}{
					Text: port.Service.Product + " version " + port.Service.Version,
				},
				HelpURI: "https://nmap.org/book/man.html",
			}
			run.Tool.Driver.Rules = append(run.Tool.Driver.Rules, rule)

			result := struct {
				RuleID  string `json:"ruleId"`
				Level   string `json:"level"`
				Message struct {
					Text string `json:"text"`
				} `json:"message"`
				Locations []struct {
					PhysicalLocation struct {
						Address struct {
							AbsoluteAddress string `json:"absoluteAddress"`
						} `json:"address"`
					} `json:"physicalLocation"`
				} `json:"locations"`
			}{
				RuleID: rule.ID,
				Level:  "note",
				Message: struct {
					Text string `json:"text"`
				}{
					Text: "Port " + string(port.PortID) + " is open.",
				},
				Locations: []struct {
					PhysicalLocation struct {
						Address struct {
							AbsoluteAddress string `json:"absoluteAddress"`
						} `json:"address"`
					} `json:"physicalLocation"`
				}{
					{
						PhysicalLocation: struct {
							Address struct {
								AbsoluteAddress string `json:"absoluteAddress"`
							} `json:"address"`
						}{
							Address: struct {
								AbsoluteAddress string `json:"absoluteAddress"`
							}{
								AbsoluteAddress: host.HostAddress[0].Address,
							},
						},
					},
				},
			}
			run.Results = append(run.Results, result)
		}

		sarifReport.Runs = append(sarifReport.Runs, run)
	}

	// Save the SARIF report
	sarifReportBytes, _ := json.MarshalIndent(sarifReport, "", "  ")
	os.WriteFile(fileOutput, sarifReportBytes, 0o644)

	log.Printf("SARIF report saved to %s\n", fileOutput)
}
