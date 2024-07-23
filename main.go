package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Current directory:", dir)

	if len(os.Args) < 2 {
		fmt.Println("Usage OST <command> [args]")
		fmt.Println("Use OST help for more information")
		return
	}

	switch os.Args[1] {
	case "help":
		fmt.Println("OST is a tool for quickly setting up an OSH project.")
		fmt.Println("Commands List:")
		fmt.Println("init: Initialize a new OSH project")
		fmt.Println("--node: Initialize a new OSH node project")
		fmt.Println("--client: Initialize a new OSH client project")
		fmt.Println("help: Display this help message")
		fmt.Println("version: Display the version of OST")

	case "init":
		fmt.Println("Initializing OSH project...")
		if len(os.Args) < 3 {
			fmt.Println("Usage OST init <project>")
			return
		}
		if os.Args[2] == "--node" {
			fmt.Println("Initializing OSH node project...")
			cloneRepo("git@github.com:opensensorhub/osh-node-dev-template.git", "master", "node")

		} else if os.Args[2] == "--client" {
			fmt.Println("Initializing OSH client project...")
			cloneRepo("git@github.com:opensensorhub/osh-client-dev-template.git", "master", "client")

		} else if os.Args[2] == "--viewer" {
			fmt.Println("Initializing OSH viewer project...")
			cloneRepo("git@github.com:opensensorhub/osh-viewer.git", "main", "viewer")

		}

	case "version":
		fmt.Println("OpenSensorHub Setup Tool v0.1.0")
	}

}

func init() {
	fmt.Println("OpenSensorHub Setup Tool v0.1.0")
}

func cloneRepo(repoUrl, branch, dir string) {
	opts := &git.CloneOptions{
		URL:               repoUrl,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		RemoteName:        branch,
	}

	_, err := git.PlainClone(dir, false, opts)

	if err != nil {
		fmt.Println(err)
		return
	}
}
