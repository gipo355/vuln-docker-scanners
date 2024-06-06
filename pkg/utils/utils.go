package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func AppendToFile(path, content string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	return nil
}

func PrintPwd() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PWD: %v", pwd)
}

func PrintEnvVars() {
	for _, env := range os.Environ() {
		log.Printf("ENV: %v", env)
	}
}

func ListFolderContent(path string) {
	cmd := exec.Command("ls", "-la", path)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func PrintFileContent(path string) {
	buf, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	stringBuf := string(buf)

	os.Stdout.WriteString(stringBuf)
}

func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)

	return err == nil
}
