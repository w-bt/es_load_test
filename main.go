package main

import (
	"es_load_test/cmd"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	version = "0.0.1"
	commit  = "n/a"
)

func main() {
	registerStackDumpReceiver()
	cli := newCLI()
	cli.Version = fmt.Sprintf("%s (Commit: %s)", version, commit)
	_ = cli.Execute()
}

func registerStackDumpReceiver() {
	sigChan := make(chan os.Signal)
	go func() {
		stacktrace := make([]byte, 32768)
		for range sigChan {
			length := runtime.Stack(stacktrace, true)
			log.Println("Stack Trace Dump")
			log.Println(string(stacktrace[:length]))
		}
	}()
	signal.Notify(sigChan, syscall.SIGQUIT)
}

func newCLI() *cobra.Command {
	cli := &cobra.Command{
		Use:   "es_load_test",
		Short: "es_load_test is the load test service",
	}

	cli.AddCommand(cmd.CreateIndexCmd())
	cli.AddCommand(cmd.StartAPIServerCmd())
	cli.AddCommand(cmd.StartVegetaServerCmd())

	return cli
}
