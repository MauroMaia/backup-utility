package main

import (
	"backup-utility/cmd"
)

const VERSION =  ""
const COMMIT_ID =  ""
const BUILD_DATE =  "2022-05-12T13:42:46+01:00"

func main() {
	cmd.Execute(VERSION, COMMIT_ID, BUILD_DATE)
}
