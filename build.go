package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	version         = "dev"
	commit          = "none"
	date            = "unknown"
	arch            = "amd64"
	operatingSystem = "linux"
	fileName        = "NextLaunch"
)

func main() {
	handle, err := os.Create("./output.txt")
	if err != nil {
		panic(err)
	}

	defer func(handle *os.File) {
		err := handle.Close()
		if err != nil {
			panic(err)
		}
	}(handle)

	args := os.Args[1:]

	for _, arg := range args {
		var value string
		var key string

		value = strings.Split(arg, "=")[1]
		key = strings.Split(arg, "=")[0]

		switch key {
		case "--os":
			operatingSystem = value
			break
		case "--arch":
			arch = value
			break
		default:
			println("Unknown argument: " + arg + " ('" + value + "') - ignoring")
		}
	}

	branchName, err := getGitBranch()

	commit, err = getGitCommit()
	if err != nil {
		commit = "none"
	}
	date, err = getGitDate()
	if err != nil {
		date = "unknown"
	}

	version = strings.ReplaceAll(version, "\n", "")
	branchName = strings.ReplaceAll(branchName, "\n", "")
	commit = strings.ReplaceAll(commit, "\n", "")
	date = strings.ReplaceAll(date, "\n", "")

	properDate, err := time.Parse("2006-01-02 15:04:05 -0700", date)

	if err != nil {
		println(err.Error())
	} else {
		date = properDate.Format("2006-01-02_15:04")
	}

	version = fmt.Sprintf("%s-%s-%s+%s", version, branchName, commit, date)

	fileName = fmt.Sprintf("%s_%s_%s", fileName, operatingSystem, arch)

	if operatingSystem == "windows" {
		fileName += ".exe"
	}

	printBuildInfo()
	compile()

	if _, err = handle.WriteString("NEXTLAUNCH_VERSION=" + version + "\r\n"); err != nil {
		fmt.Println("Failed to write output variable: ('NEXTLAUNCH_VERSION') ", err.Error())
	}
	if _, err = handle.WriteString("NEXTLAUNCH_COMMIT=" + commit + "\r\n"); err != nil {
		fmt.Println("Failed to write output variable: ('NEXTLAUNCH_COMMIT') ", err.Error())
	}
	if _, err = handle.WriteString("NEXTLAUNCH_DATE=" + date + "\r\n"); err != nil {
		fmt.Println("Failed to write output variable: ('NEXTLAUNCH_DATE') ", err.Error())
	}
	if _, err = handle.WriteString("NEXTLAUNCH_BRANCH=" + branchName + "\r\n"); err != nil {
		fmt.Println("Failed to write output variable: ('NEXTLAUNCH_BRANCH') ", err.Error())
	}
	if _, err = handle.WriteString("NEXTLAUNCH_FILENAME=" + fileName + "\r\n"); err != nil {
		fmt.Println("Failed to write output variable: ('NEXTLAUNCH_FILENAME') ", err.Error())
	}

	if err := handle.Sync(); err != nil {
		println(err.Error())
	}

	// read all lines from file
	// and print them

	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func getGitBranch() (string, error) {
	cmd := "git rev-parse --abbrev-ref HEAD"
	out, err := runCommand(cmd)
	if err != nil {
		return "none", err
	}
	return out, nil
}

func getGitCommit() (string, error) {
	cmd := "git rev-parse --short HEAD"
	out, err := runCommand(cmd)
	if err != nil {
		return "none", err
	}
	return out, nil
}

func getGitDate() (string, error) {
	cmd := "git show -s --format=%ci HEAD"
	out, err := runCommand(cmd)
	if err != nil {
		return "unknown", err
	}
	return out, nil
}

func runCommand(cmd string) (string, error) {
	out, err := exec.Command("sh", "-c", cmd).Output()

	if err != nil {
		if err.Error() == "exec: \"sh\": executable file not found in %PATH%" {
			out, err = exec.Command("powershell.exe", "-c", cmd).Output()
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}
	return string(out), nil
}

func printBuildInfo() {
	println("Build Info:")
	println("  Version:  " + version)
	println("  Commit:   " + commit)
	println("  Date:     " + date)
	println("  Arch:     " + arch)
	println("  OS:       " + operatingSystem)
	println("  Filename: " + fileName)
}

func compile() {
	println("Compiling...")
	command := "go build -ldflags=\""
	command += "-X 'Nextlaunch/src/config.Version=" + version + "'"
	command += " -X 'Nextlaunch/src/config.BuildCommit=" + commit + "'"
	command += " -X 'Nextlaunch/src/config.BuildDate=" + date + "'"
	command += " -X 'Nextlaunch/src/config.BuildOS=" + operatingSystem + "'"
	command += " -X 'Nextlaunch/src/config.BuildArch=" + arch + "'"

	command += "\""
	command += " -o ./binaries/" + fileName

	command += " main.go"

	if err := exec.Command("sh", "-c", command).Run(); err != nil {
		println("Build Error")

		if err.Error() == "exec: \"sh\": executable file not found in %PATH%" {
			println("Detected Windows, trying to compile under powershell")
			if err := exec.Command("powershell.exe", "-c", command).Run(); err != nil {
				println("Build Error")

				if err.Error() == "exec: \"powershell.exe\": executable file not found in $PATH" {
					println("Could not find powershell.exe, please install it and try again")
				}

				println(err.Error())
			}
		} else {
			println("Build Failed")
			println(err.Error())
		}
	}
}
