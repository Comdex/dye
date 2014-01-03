package dye

import (
	"bytes"
	"os/exec"
	"strings"
)

var bin = "pygmentize"

func Binary(path string) error {
	if _, err := exec.LookPath(path); err != nil {
		return err
	}

	bin = path
	return nil
}

func Which() string {
	return bin
}

func Highlight(code string, lexer string, format string, enc string) (string, error) {
	cmd := exec.Command(bin, "-l"+lexer, "-f"+format, "-O encoding="+enc)
	cmd.Stdin = strings.NewReader(code)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return stdout.String(), nil
}
