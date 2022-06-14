package utils

import (
	"os"
	"strings"
)

var rsyncCmdString string
var touchCmdString string

func init() {
	// Find rsync command in the system
	stdout, err := ExecuteShell("/usr/bin/whereis", "rsync")
	if err != nil {
		LogPanic("Command `rsync` not found. Reason: %s", err.Error())
	}

	words := strings.Fields(stdout)
	rsyncCmdString = words[1]

	LogDebug("rsync location: %s", rsyncCmdString)

	// Find touch command in the system
	stdout, err = ExecuteShell("/usr/bin/whereis", "rsync")
	if err != nil {
		LogPanic("Command `rsync` not found. Reason: %s", err.Error())
	}

	words = strings.Fields(stdout)
	rsyncCmdString = words[1]

	LogDebug("rsync location: %s", rsyncCmdString)
}

func CheckConnectivity(host string) {
	fileName := "/tmp/emptyFile.txt"
	file, err := os.Create(fileName)
	defer file.Close()

	if err != nil {
		LogError("Failed to create empty file (%s). Reason: %s", fileName, err.Error())
	}

	ExecuteShell(rsyncCmdString, "-avz", fileName, host+":"+fileName+".copy")
}
