package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	version         = "dev"
	tag             = "none"
	branchName      = "none"
	commit          = "none"
	date            = "unknown"
	arch            = "amd64"
	operatingSystem = "linux"
	fileName        = "NextLaunch"
	metadata        = ""
)

func main() {

	manifest := getBuildManifest()

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

	var value string
	var key string
	greedy := false

	for _, arg := range args {

		if strings.Contains(arg, "=") {
			value = strings.Split(arg, "=")[1]
			key = strings.Split(arg, "=")[0]
		} else {
			if greedy {
				value = arg
				greedy = false
			} else {
				key = arg
				value = "GREEDY"
			}
		}

		if value == "GREEDY" {
			greedy = true
			continue
		}

		switch key {
		case "--os":
			operatingSystem = value
			break
		case "--arch":
			arch = value
			break
		case "--metadata":
			metadata = "+" + value

			break
		default:
			println("Unknown argument: " + arg + " ('" + value + "') - ignoring")
		}
	}

	tag, err = getGitTag()

	if err != nil {
		tag = "0.0.0"
	}

	branchName, err = getGitBranch()

	commit, err = getGitCommit()
	if err != nil {
		commit = "none"
	}
	date, err = getGitDate()
	if err != nil {
		date = "unknown"
	}

	manifest.Tag = strings.ReplaceAll(tag, "\n", "")
	manifest.Version = strings.ReplaceAll(tag, "\n", "")
	manifest.Branch = strings.ReplaceAll(branchName, "\n", "")
	manifest.Commit = strings.ReplaceAll(commit, "\n", "")
	manifest.Date = strings.ReplaceAll(date, "\n", "")
	manifest.Metadata = strings.ReplaceAll(metadata, "\n", "")

	if len(manifest.Files) == 0 {
		date = strings.ReplaceAll(date, "\n", "")

		properDate, err := time.Parse("2006-01-02 15:04:05 -0700", date)

		if err != nil {
			println(err.Error())
		} else {
			manifest.Date = properDate.Format("2006-01-02_15:04")
		}

		manifest.Version = fmt.Sprintf("%s-%s.%s%s", manifest.Version, manifest.Branch, manifest.Commit, manifest.Metadata)
	}

	fileName = fmt.Sprintf("%s_%s_%s", fileName, operatingSystem, arch)

	if operatingSystem == "windows" {
		fileName += ".exe"
	}

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

	printBuildInfo(manifest)
	compile(manifest)

	err = appendBuildManifest(&manifest, "binaries/"+fileName)

	if err != nil {
		println("Error appending manifest.json")
	}

	err = writeBuildManifest(manifest)
	if err != nil {
		println("Error writing manifest.json")
	}

	println("Build Complete")

	fmt.Println("========= Environmental Variables =========")
	//	debug output file line by line

	_, err = handle.Seek(0, 0)

	if err != nil {
		println("Error seeking to beginning of file")
		return
	}

	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("==========================================")

	err = handle.Sync()

	if err != nil {
		return
	}

	err = handle.Close()

	if err != nil {
		return
	}
}

type Manifest struct {
	Version  string         `json:"version"`
	Tag      string         `json:"tag"`
	Branch   string         `json:"branch"`
	Commit   string         `json:"commit"`
	Date     string         `json:"date"`
	Metadata string         `json:"metadata"`
	Files    []ManifestFile `json:"files"`
}

type ManifestFile struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Size   int    `json:"size"`
	Arch   string `json:"arch"`
	OS     string `json:"os"`
	Sha256 string `json:"sha256"`
}

func getBuildManifest() Manifest {
	var manifest Manifest

	file, err := os.Open("manifest.json")
	if err != nil {
		println("Error opening manifest.json")
		return manifest
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	byteValue, _ := io.ReadAll(file)

	err = json.Unmarshal(byteValue, &manifest)
	if err != nil {
		println("Error unmarshalling manifest.json")

		manifest.Files = []ManifestFile{}

		return manifest
	}

	if manifest.Files == nil {
		manifest.Files = []ManifestFile{}
	}

	return manifest
}

func appendBuildManifest(manifest *Manifest, path string) error {
	file, err := os.Open(path) // Open the binary file
	if err != nil {
		println("Error opening binary file")
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	hasher := sha256.New()

	_, err = io.Copy(hasher, file)
	if err != nil {
		return err
	}

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	manifest.Files = append(manifest.Files, ManifestFile{
		Name:   path,
		Path:   path,
		Size:   int(stat.Size()),
		Arch:   arch,
		OS:     operatingSystem,
		Sha256: hex.EncodeToString(hasher.Sum(nil)),
	})

	return nil
}

func writeBuildManifest(manifest Manifest) error {
	file, err := os.Create("manifest.json")
	if err != nil {
		println("Error creating manifest.json")
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(manifest)
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}
	return nil
}

func getGitTag() (string, error) {
	println("Getting git tag")
	cmd := "git describe --tags --abbrev=0"
	out, err := runCommand(cmd)
	if err != nil {
		println("Error getting git tag")
		return "none", err
	}
	println("Git tag: " + out)
	return out, nil
}

func getGitBranch() (string, error) {
	println("Getting git branch")
	cmd := "git rev-parse --abbrev-ref HEAD"
	out, err := runCommand(cmd)
	if err != nil {
		println("Error getting git branch")
		return "none", err
	}
	println("Git branch: " + out)
	return out, nil
}

func getGitCommit() (string, error) {
	println("Getting git commit")
	cmd := "git rev-parse --short HEAD"
	out, err := runCommand(cmd)
	if err != nil {
		println("Error getting git commit")
		return "none", err
	}
	println("Git commit: " + out)
	return out, nil
}

func getGitDate() (string, error) {
	println("Getting git date")
	cmd := "git show -s --format=%ci HEAD"
	out, err := runCommand(cmd)

	if err != nil {
		println("Error getting git date")
		return "unknown", err
	}
	println("Git date: " + out)
	return out, nil
}

func runCommand(cmd string) (string, error) {
	println("Running command: " + cmd)
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

func printBuildInfo(manifest Manifest) {
	println("Build Info:")
	println("  Version:  " + manifest.Version)
	println("  Tag:      " + manifest.Tag)
	println("  Branch:   " + manifest.Branch)
	println("  Metadata: " + manifest.Metadata)
	println("  Commit:   " + manifest.Commit)
	println("  Date:     " + manifest.Date)
	println("  Arch:     " + arch)
	println("  OS:       " + operatingSystem)
	println("  Filename: " + fileName)
}

func compile(manifest Manifest) {
	println("Compiling...")
	command := "go build -ldflags=\""
	command += "-X 'Nextlaunch/src/config.Version=" + manifest.Version + "'"
	command += " -X 'Nextlaunch/src/config.BuildCommit=" + manifest.Commit + "'"
	command += " -X 'Nextlaunch/src/config.BuildDate=" + manifest.Date + "'"
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
