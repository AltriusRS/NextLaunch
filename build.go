package main

import (
	"fmt"
	"os/exec"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {

	branchName := getGitBranch()

	commit = getGitCommit()
	date = getGitDate()
	version = fmt.Sprintf("%s-%s-%s", version, branchName, commit)

	if len(version) == 0 {
		version = "dev"
	}
	if len(commit) == 0 {
		commit = "none"
	}
	if len(date) == 0 {
		date = "unknown"
	}

	build := version
	if commit != "none" {
		build += "-" + commit
	}
	if date != "unknown" {
		build += "-" + date
	}

	printBuildInfo(build)
}

func getGitBranch() string {
	cmd := "git rev-parse --abbrev-ref HEAD"
	out, err := runCommand(cmd)
	if err != nil {
		return "none"
	}
	return out
}

func getGitCommit() string {
	cmd := "git rev-parse --short HEAD"
	out, err := runCommand(cmd)
	if err != nil {
		return "none"
	}
	return out
}

func getGitDate() string {
	cmd := "git show -s --format=%ci HEAD"
	out, err := runCommand(cmd)
	if err != nil {
		return "unknown"
	}
	return out
}

func runCommand(cmd string) (string, error) {
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func printBuildInfo(build string) {
	println("Build Info:")
	println("  Version: " + build)
	println("  Commit:  " + commit)
	println("  Date:    " + date)
	//panic("Build Complete")
}
