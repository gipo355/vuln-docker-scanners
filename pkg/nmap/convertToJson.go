package nmap

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/vdjagilev/nmap-formatter/v3/formatter"
)

// ConvertToJSON converts the nmap xml report to json.
// Specify the name of the report to convert.
// Can be either "vulscan", "direct", or "vulners".
func (n *Client) ConvertToJSON(name ReportName) error {
	mainDir := n.Config.OutputDir
	fileInput := mainDir + "/" + string(name) + "/" + string(name) + "-report.xml"
	fileOutput := mainDir + "/" + string(name) + "/" + string(name) + "-report.json"

	var nmap formatter.NMAPRun

	var config formatter.Config = formatter.Config{}

	// Read XML file that was produced by nmap (with -oX option)
	content, err := os.ReadFile(fileInput)
	if err != nil {
		panic(err)
	}

	// Unmarshal XML and map structure(s) fields accordingly
	if err = xml.Unmarshal(content, &nmap); err != nil {
		return fmt.Errorf("failed to unmarshal xml: %w", err)
	}

	// Output data to console stdout
	// You can use any other io.Writer implementation
	// for example: os.OpenFile("file.json", os.O_CREATE|os.O_EXCL|os.O_WRONLY, os.ModePerm)
	// config.Writer = os.Stdout
	outputFile, err := os.Create(fileOutput)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	config.Writer = outputFile
	// Formatting data to JSON, you can use:
	// CSVOutput, MarkdownOutput, HTMLOutput as well
	config.OutputFormat = formatter.JSONOutput

	// Setting formatter data/options
	templateData := formatter.TemplateData{
		NMAPRun: nmap, // NMAP output data itself
		OutputOptions: formatter.OutputOptions{
			JSONOptions: formatter.JSONOutputOptions{
				PrettyPrint: true, // Additional option to prettify JSON
			},
		},
	}

	// New formatter instance
	formatter := formatter.New(&config)
	if formatter == nil {
		// Not json/markdown/html/csv
		return fmt.Errorf("wrong formatter provided")
	}

	// Attempt to format the data
	if err = formatter.Format(&templateData, "" /* no template content for JSON */); err != nil {
		// html template could not be parsed or some other issue occured
		return fmt.Errorf("failed to format data: %w", err)
	}

	log.Printf("Successfully converted %s to JSON", name)

	return nil
}
