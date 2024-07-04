# xkcd-cli
A simple CLI application to download and display [xkcd comics](https://xkcd.com/).

## Usage
Display the latest comic with the system's default image viewer:
```bash
xkcd-cli display latest
```
Display a specific comic with the system's default image viewer:
```bash
xkcd-cli display [number]
```
Display a random comic with the system's default image viewer:
```bash
xkcd-cli display random
```
Download the latest comic to the current directory:
```bash
xkcd-cli get latest
```
Download a specific comic to the current directory:
```bash
xkcd-cli get [number]
```
Download a random comic to the current directory:
```bash
xkcd-cli get random
```

## Installation

### Install with prebuilt binaries
Get the binary for your system in [releases](https://github.com/AliGaygisiz/xkcd-cli/releases/latest).

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
