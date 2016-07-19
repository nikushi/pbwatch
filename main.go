package main

import (
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

func main() {
	interval := 500 * time.Millisecond
	var prevText string

	for {
		// clear current screen
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		// get buffer
		text := C.GoString(C.getTextFromClipboard())

		// To desktop notification
		if prevText != text {
			note := gosxnotifier.NewNotification(text)
			note.Title = "pbwatch"
			note.Subtitle = "Pasted!"

			err := note.Push()
			if err != nil {
				log.Fatal(err)
			}
		}

		// To terminal
		fmt.Println(text)

		// memo
		prevText = text

		time.Sleep(interval)
	}
}
