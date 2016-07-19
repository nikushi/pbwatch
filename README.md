# pbwatch

`pbwatch` - display and update current clipboard text on terminal.

## Features

* Display current cliboard text to your terminal in foreground.
* Send event to notification center when you copy new text.

## Requirements

* OS X v10.6 and later

## Install

    go get github.com/niku4i/pbwatch

## Usage

Just run `pbwatch`,

```
$ pbwatch
Here text you pasted before

```

Then, paste something by `⌘ + c`. The text that you've just copied will appear to the terminal screen. Also the event will be sent to the notification center.

