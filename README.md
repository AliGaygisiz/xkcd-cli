# xkcd-cli
A simple CLI application to download and display [xkcd comics](https://xkcd.com/).

## Features
- `get` command to download latest xkcd comic to current directory.

## Planned Features
- Add the ability to download a specific comic by its ID.
- Add `display` command to open the comic in systems default image viewer without saving the image.
- Add `random` option to get a random comic.

## Installation

### Install with `git`
```bash
git clone https://github.com/AliGaygisiz/xkcd-cli.git
cd xkcd-cli
```
If you just want to build and test:
```bash
go build .
```
If you want to install the binary to your default Go bin folder:
```bash
go install .
```
