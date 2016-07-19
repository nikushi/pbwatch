# pbwatch

`pbwatch` always display current clibpard text on loop.

## Features

* Displaying current cliboard buffer to your terminal in foreground.
* Send notification to notification center when text on cliboard changes.

## Requirements

* OS X v10.6 and later

## Install

    go get github.com/niku4i/pbwatch

## Usage

Just run `pbwatch` on your terminal.

```
$ pbwatch
Here text you pasted before

```

Then, paste something by `⌘ + c`. The screen of your terminal will change, and notification will be sent.
