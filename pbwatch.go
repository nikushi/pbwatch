package main

import (
	"sync"
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

type Pbwatch struct {
	wg          sync.WaitGroup
	interval    time.Duration
	notifyCh    chan string
	prevText    string
	currentText string
}

func NewPbwatch(interval time.Duration) *Pbwatch {
	return &Pbwatch{
		interval: interval,
	}
}

func (p *Pbwatch) Start(notifyCh chan string) {
	p.notifyCh = notifyCh
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		defer close(p.notifyCh)
		p.starttick()
	}()
}

func (p *Pbwatch) starttick() {
	p.prevText = getTextFromClipboard()
	p.currentText = p.prevText
	p.notifyCh <- p.currentText
	for {
		if p.prevText != p.currentText {
			p.notifyCh <- p.currentText
			p.prevText = p.currentText
		}
		p.currentText = getTextFromClipboard()
		time.Sleep(p.interval)
	}
}

func (p *Pbwatch) Wait() {
	p.wg.Wait()
}

func getTextFromClipboard() string {
	return C.GoString(C.getTextFromClipboard())
}
