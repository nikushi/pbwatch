package main

import (
	"flag"
	"fmt"
	"github.com/deckarep/gosx-notifier"
	"log"
	"os"
	"os/exec"
	"time"
)

const Version = "0.1.0"

const helpText = `Usage: pbwatch [-n]

  pbwatch - display and update current clipboard text on terminal

Options:
	-n           Send event to desktop notification center
	-version     Print version

`

func printUsage() {
	fmt.Fprintf(os.Stderr, helpText)
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func sendNotification(text string) {
	note := gosxnotifier.NewNotification(text)
	note.Title = "pbwatch"
	note.Subtitle = "Copied!"
	note.Group = "pbwatch"

	err := note.Push()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var optNotify bool
	var optVersion bool

	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.Usage = func() { printUsage() }
	flags.BoolVar(&optNotify, "n", false, "Popup copy event")
	flags.BoolVar(&optNotify, "notification", false, "Popup copy event")
	flags.BoolVar(&optVersion, "version", false, "Version")
	if err := flags.Parse(os.Args[1:]); err != nil {
		flags.Usage()
		os.Exit(1)
	}

	if optVersion {
		fmt.Fprintf(os.Stderr, "pbwatch version %s\n", Version)
		os.Exit(0)
	}

	interval := 500 * time.Millisecond

	pbwatch := NewPbwatch(interval)
	notifyCh := make(chan string)
	go func() {
		// suppress notification on boot
		text := <-notifyCh
		clearScreen()
		fmt.Println(text)

		// then run loop
		for text := range notifyCh {
			clearScreen()
			fmt.Println(text)
			if optNotify {
				sendNotification(text)
			}
		}
	}()

	pbwatch.Start(notifyCh)

	// wait for loop above
	pbwatch.Wait()

}
