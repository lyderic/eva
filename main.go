package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/lyderic/tools"
)

const (
	VERSION = "0.0.0"
)

var (
	debug     bool
	agent     Agent
	variables []Variable
)

func init() {
}

func main() {
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.StringVar(&agent.Port, "port", "19970", "agent TCP port")
	flag.Usage = usage
	flag.Parse()
	agent.Address = net.JoinHostPort("localhost", agent.Port)
	if debug {
		tools.PrintRed("*** DEBUG MODE ***")
	}
	if !agent.isRunning() {
		agent.start()
	}
	if len(flag.Args()) == 0 {
		usage()
		return
	}
	command := flag.Args()[0]
	switch command {
	case "add":
		add("name", "value")
	case "get":
		get("name")
	case "list", "ls":
		list()
	case "del", "rm":
		del("name")
	case "version":
		version()
	default:
		tools.PrintRed("command not found!")
		usage()
	}
}

func version() {
	fmt.Println("eva", VERSION, "(c) Lyderic Landry, London 2019")
}

func usage() {
	version()
	fmt.Println("Usage: eva <option> command [<name>] [<value>]")
	fmt.Println("Commands:")
	fmt.Println("  add <name> <value>  add variable")
	fmt.Println("  get <name>          get variable")
	fmt.Println("  list                list variables in agent")
	fmt.Println("  del <name>          delete variable")
	fmt.Println("  version             show version and exit")
	fmt.Println("Options:")
	flag.PrintDefaults()
}
