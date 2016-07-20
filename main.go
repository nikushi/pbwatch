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

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

const char* getTextFromClipboard() {
	NSPasteboard* pasteboard = [NSPasteboard generalPasteboard];
	NSString* pbBuf = [pasteboard stringForType:NSStringPboardType];
	const char *strc = [pbBuf UTF8String];
	return strc;
}
*/
import "C"

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var optNotify bool
	flag.BoolVar(&optNotify, "n", false, "Popup copy event")
	flag.BoolVar(&optNotify, "notification", false, "Popup copy event")
	flag.Parse()

	interval := 500 * time.Millisecond
	var prevText string

	for {

		// get buffer
		text := C.GoString(C.getTextFromClipboard())

		// To desktop notification
		if optNotify && prevText != "" && prevText != text && text != "" {
			note := gosxnotifier.NewNotification(text)
			note.Title = "pbwatch"
			note.Subtitle = "Pasted!"
			note.Group = "pbwatch"

			err := note.Push()
			if err != nil {
				log.Fatal(err)
			}
		}

		// To terminal
		if prevText == "" || (prevText != text && text != "") {
			clearScreen()
			fmt.Println(text)
		}

		// memo
		prevText = text

		time.Sleep(interval)
	}
}
