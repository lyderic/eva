package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
)

type Agent struct {
	Port    string
	Address string
	Server  http.Server
	Pid     int
}

func (agent Agent) isRunning() (running bool) {
	conn, err := net.Dial("tcp", agent.Address)
	if err != nil {
		running = false
	}
	if conn != nil {
		running = true
		conn.Close()
	}
	return
}

func (agent *Agent) start() (err error) {
	agent.Pid = os.Getpid()
	fmt.Println("Starting eva agent. PID:", agent.Pid)
	mux := http.NewServeMux()
	agent.Server = http.Server{Addr: agent.Address, Handler: mux}
	mux.HandleFunc("/", process)
	return agent.Server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	command := r.URL.String()
	switch command {
	case "/info":
		fmt.Fprintf(w, "Version: %s\n", VERSION)
		fmt.Fprintf(w, "PID:     %d\n", agent.Pid)
		fmt.Fprintf(w, "Port:    %s\n", agent.Port)
	case "/kill":
		fmt.Fprintf(w, "Killed eva agent (PID: %d)\n", agent.Pid)
		go func() {
			agent.Server.Shutdown(context.Background())
		}()
	default:
		fmt.Fprintf(w, "%s: command not found!\n", command)
		fmt.Fprintf(w, "Valid commands are: info, kill\n")
	}
}
