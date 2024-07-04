# xkcd-cli
A simple CLI application to download and display [xkcd comics](https://xkcd.com/).

## Features
- `xkcd-cli get latest` to download the latest comic to the current directory.
- `xkcd-cli get [number]` to download the comic with the given number to current directory.
- `xkcd-cli get random` to download a random comic to current directory.

## Planned Features
- Add `display` command to open the comic in the system's default image viewer without saving the image.

## Installation

### Install with `go install`
```bash
go install github.com/AliGaygisiz/xkcd-cli@latest
```

### Install with `git`
```bash
git clone https://github.com/AliGaygisiz/xkcd-cli.git
cd xkcd-cli
go build .
```
If you want to install the binary to your default Go bin folder:
```bash
go install .
```
