package helpers

import (
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

// ExecuteCommand executes a single command and returns the results
func ExecuteCommand(command string) (result string, err error) {

	args := strings.Fields(command)
	exe := exec.Command(args[0], args[1:]...)
	out, err := exe.CombinedOutput()
	cmdLogger := log.WithFields(log.Fields{
		"command": args[0],
		"args":    strings.Join(args[1:], " "),
		"out":     string(out),
		"err":     err,
	})
	if err != nil {
		cmdLogger.Error("command executed with errors")
	} else {
		cmdLogger.Debug("command executed")
	}
	return string(out), err
}

// ExecuteCommands executes a slice of commands and returns the results
func ExecuteCommands(command []string) (results []string, err error) {

	results = make([]string, len(command))
	for _, c := range command {
		if c == "" {
			continue
		}
		cresult, cerr := ExecuteCommand(c)
		if cerr != nil {
			return results, cerr
		}
		results = append(results, cresult)
	}

	return results, nil
}