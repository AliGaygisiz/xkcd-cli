A simple CLI application to download and display [xkcd comics](https://xkcd.com/).

Used [urfave/cli](https://github.com/urfave/cli) for the cli *magic*.

{% alert(fact=true) %}
I originally planned to use [cobra](https://github.com/spf13/cobra) for this project. However, on the day I started development, a dependency of `cobra` was down (the developer was hosting it on a private server), preventing me from downloading it. So, I started with `urfave/cli` instead.
{% end %}

## Usage

### Display Command

Display the latest comic using your system's default image viewer:

```bash
xkcd-cli display latest
```

Display a specific comic by number:

```bash
xkcd-cli display 42
```

Display a random comic:

```bash
xkcd-cli display random
```

### Get Command

Download the latest comic to the current directory:

```bash
xkcd-cli get latest
```

Download a specific comic:

```bash
xkcd-cli get 42
```

Download a random comic:

```bash
xkcd-cli get random
```

## Installation

### Prebuilt Binaries

Download the binary for your system in [Releases Page](https://github.com/AliGaygisiz/xkcd-cli/releases/latest)

### Install via Go

```bash
go install github.com/AliGaygisiz/xkcd-cli@latest
```

### Build from Source

```bash
git clone https://github.com/AliGaygisiz/xkcd-cli.git
cd xkcd-cli
go build .
```

To install the binary to your `$GOPATH/bin`:

```bash
go install .
```
