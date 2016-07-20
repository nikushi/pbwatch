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

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func sendNotification(text string) {
	note := gosxnotifier.NewNotification(text)
	note.Title = "pbwatch"
	note.Subtitle = "Pasted!"
	note.Group = "pbwatch"

	err := note.Push()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var optNotify bool
	flag.BoolVar(&optNotify, "n", false, "Popup copy event")
	flag.BoolVar(&optNotify, "notification", false, "Popup copy event")
	flag.Parse()

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
