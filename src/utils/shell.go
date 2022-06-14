package utils

import (
	"os/exec"
)

func ExecuteShell(cmdString string, arg ...string) (string, error) {

	LogDebug("Running this command: %s", cmdString)

	cmd := exec.Command(cmdString, arg...)

	// Start long-running execution
	err := cmd.Start()
	if err != nil {
		return "", err
	}

	// Wait until execution end
	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	stdout, err := cmd.Output()
	if err != nil {
		LogError(err.Error())
		return "", err
	}

	stdoutStr := string(stdout)
	LogDebug("Command result: %s", stdoutStr)

	return stdoutStr, nil
}
