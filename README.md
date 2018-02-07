# riff

[![Build Status](https://travis-ci.org/gimke/riff.svg?branch=master)](https://travis-ci.org/gimke/riff) [![GoDoc](https://godoc.org/github.com/gimke/riff?status.svg)](https://godoc.org/github.com/gimke/riff)

Riff is a tool for service discovery and configuration.

* Service Discovery
* Health Checking
* Web console

## Quick Start
If you wish to work on Riff, you'll first need Go installed (version 1.9+ is required).

If you need Web Console, you'll need npm installed dependencies            

```bash
npm install
```
Next clone this repository and build Riff or download from [release](https://github.com/gimke/riff/releases)
```bash
make
bin/riff -v
```
## Commands(CLI)
```bash
$ ./riff

Usage: riff [--version] <command> [<args>]

Available commands are:

  query       Query
  restart     Restart service
  run         Run Riff
  start       Start service
  stop        Stop service
  version     Prints the Riff version

```

### run
run Riff
```bash
$ ./riff run -h

Usage: run [options]

  Run riff service

Options:

  -name       Node name
  -dc         DataCenter name
  -http       Http address of riff (-http 127.0.0.1:8610)
  -rpc        RPC address of riff (-rpc [::]:8630)
  -join       Join RPC address (-join 192.168.1.1:8630,192.168.1.2:8630,192.168.1.3:8630)

```

## Web Console

![Riff console](https://raw.githubusercontent.com/gimke/riff/gh-pages/images/screen.png)
