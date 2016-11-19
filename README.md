# debug-this-motherfucker
A collection of mind fucking trolling hacks. The idea is to troll your teammates for training purposes - you'll never forget Linux's dark corners after debugging a mind boggling issue for 30 minutes.

Example (the invisble-net trolling):

```
$ ssh test-server
Last login: Tue Feb 23 05:06:01 2016 from 10.0.2.2
user@test-server:~$ ip a
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00

user@test-server:~$ # wtf? how did I ssh into a server that has no network?
```

## Building/Installing

`debug-this-motherfucker` is written in golang and compiled as a static binary. Thus no installation is required, simply download a binary from the [releases](https://github.com/avishai-ish-shalom/debug-this-motherfucker/releases) page.

To build, run `./build.sh` or `go get -d -v && go build -o dbtm`

## Usage

- `--help` or `help` subcommand to display help
- `--self-destruct` will remove the binary after trolling the server (to avoid easy detection)
- `--explain` can be used with any troll plugin to display an explanation of what the troll does
