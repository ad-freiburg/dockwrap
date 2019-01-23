package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/ad-freiburg/wharfer/wrap"
)

func execDocker(args ...string) {
	const dockerbin = "/usr/bin/docker"
	cmd := exec.Command(dockerbin, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

var build wrap.Build
var run wrap.Run
var ps wrap.Ps
var kill wrap.Kill
var rm wrap.Rm
var logs wrap.Logs
var pull wrap.Pull
var images wrap.Images
var networkCreate wrap.NetworkCreate
var networkList wrap.NetworkList
var networkRemove wrap.NetworkRemove

func init() {
	build.InitFlags()
	run.InitFlags()
	ps.InitFlags()
	kill.InitFlags()
	logs.InitFlags()
	rm.InitFlags()
	pull.InitFlags()
	images.InitFlags()
	networkCreate.InitFlags()
	networkList.InitFlags()
	networkRemove.InitFlags()
}

var version = "no-release"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "<COMMAND>|--version")
		fmt.Fprintln(os.Stderr, "Commands:")
		fmt.Fprintln(os.Stderr, "\tbuild")
		fmt.Fprintln(os.Stderr, "\trun")
		fmt.Fprintln(os.Stderr, "\tps")
		fmt.Fprintln(os.Stderr, "\tkill")
		fmt.Fprintln(os.Stderr, "\trm")
		fmt.Fprintln(os.Stderr, "\tlogs")
		fmt.Fprintln(os.Stderr, "\tpull")
		fmt.Fprintln(os.Stderr, "\timages")
		fmt.Fprintln(os.Stderr, "\tnetwork")
		os.Exit(1)
	}

	var args []string
	switch os.Args[1] {
	case "build":
		args = build.ParseToArgs(os.Args[2:])
	case "run":
		args = run.ParseToArgs(os.Args[2:])
	case "kill":
		args = kill.ParseToArgs(os.Args[2:])
	case "rm":
		args = rm.ParseToArgs(os.Args[2:])
	case "logs":
		args = logs.ParseToArgs(os.Args[2:])
	case "ps":
		args = ps.ParseToArgs(os.Args[2:])
	case "pull":
		args = pull.ParseToArgs(os.Args[2:])
	case "images":
		args = images.ParseToArgs(os.Args[2:])
	case "network":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Missing subcommand")
			os.Exit(1)
		} else {
			switch os.Args[2] {
			case "create":
				args = networkCreate.ParseToArgs(os.Args[3:])
			case "ls":
				args = networkList.ParseToArgs(os.Args[3:])
			case "rm":
				args = networkRemove.ParseToArgs(os.Args[3:])
			default:
				fmt.Fprintln(os.Stderr, "Unknown subcommand", os.Args[2])
				os.Exit(1)
			}
		}
	case "--version":
		fmt.Fprintln(os.Stderr, os.Args[0], "version", version)
		os.Exit(0)
	default:
		fmt.Fprintln(os.Stderr, "Unknown subcommand", os.Args[1])
		os.Exit(2)
	}

	execDocker(args...)
}
