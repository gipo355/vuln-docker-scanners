package nmap

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"slices"
)

// writeToFile writes the nmap output to a file in the reports directory
func (n *Client) writeToFile(userArgs []string, dirName string, report ReportName) error {
	target := n.Config.Target

	mainDir := n.Config.OutputDir

	// TODO: don't hardcode the directory name and the file name
	fileName := string(report) + "-report"

	args := slices.Concat(userArgs, []string{
		"-oA",                                    // output all formats
		mainDir + "/" + dirName + "/" + fileName, // output file name
		target,                                   // target
	})

	cmd := exec.Command("nmap", args...)
	log.Printf("cmd: %v", cmd)

	err := os.MkdirAll(mainDir+"/"+dirName, 0o755)
	if err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	openFile, fileErr := os.Create(mainDir + "/" + dirName + "/" + string(report) + "-output.log")
	if fileErr != nil {
		return fmt.Errorf("error creating file: %w", fileErr)
	}
	defer openFile.Close()

	writer := bufio.NewWriter(openFile)
	defer writer.Flush()

	// create a pipe to capture stdout and stderr
	stdout, outPipeErr := cmd.StdoutPipe()
	if outPipeErr != nil {
		return fmt.Errorf("error creating stdout pipe: %w", outPipeErr)
	}

	stderr, errPipeErr := cmd.StderrPipe()
	if errPipeErr != nil {
		return fmt.Errorf("error creating stderr pipe: %w", errPipeErr)
	}

	// start the command
	if cmdErr := cmd.Start(); cmdErr != nil {
		return fmt.Errorf("error starting command: %w", cmdErr)
	}

	// createa a channel to get the output
	stdErrOutputChan := make(chan []byte)
	stdOutOutputChan := make(chan []byte)

	// create a goroutine to copy the stdout in a stream
	go func() {
		_, copyStdoutErr := io.Copy(writer, stdout)
		if copyStdoutErr != nil {
			stdOutOutputChan <- []byte(fmt.Sprintf("Error copying stdout: %s", copyStdoutErr))
		}
	}()

	// create a goroutine to copy the stderr in a stream
	go func() {
		_, copyStderrErr := io.Copy(os.Stderr, stderr)
		if copyStderrErr != nil {
			stdErrOutputChan <- []byte(fmt.Sprintf("Error copying stderr: %s", copyStderrErr))
		}

		stdErrOutputChan <- []byte{}
	}()

	if stdErrOutput := <-stdErrOutputChan; len(stdErrOutput) > 0 {
		return fmt.Errorf("error copying std: %s", stdErrOutput)
	}

	// wait for the command to finish
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("error waiting for command: %w", err)
	}

	return nil
}
