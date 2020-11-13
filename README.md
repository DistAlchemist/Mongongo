# Mongongo

TL;DR: Cassandra and Zookeeper in Go.

# Getting Started

* Install [Go 1.15.2](https://golang.org/dl/)
* Enable Go module and Setup Go package proxy

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
```

* Download

```shell
git clone https://github.com/DistAlchemist/Mongongo.git
cd Mongongo/src
```

* Setup [tmux](https://github.com/tmux/tmux/wiki) for multi-terminal (recommended)

```shell
sudo apt install tmux
```

* Create a new session of tmux:

```shell
tmux new -s mg
```

* Inside one terminal, run Mongongo server:

```shell
go run cmd/mgserver/main.go
```

* Inside another terminal, run command line interface:

```shell
go run cmd/cli/main.go
```

# Contributing


# Example

![mongongo](pics/mongongo.gif)

# License

MIT License