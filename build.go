package main

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
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

func printBuildInfo(build string) {
	println("Build Info:")
	println("  Version: " + build)
	println("  Commit:  " + commit)
	println("  Date:    " + date)
}
