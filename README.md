# eztables

[![Build Status](https://travis-ci.org/posener/eztables.svg?branch=master)](https://travis-ci.org/posener/eztables)
[![codecov](https://codecov.io/gh/posener/eztables/branch/master/graph/badge.svg)](https://codecov.io/gh/posener/eztables)
[![GoDoc](https://godoc.org/github.com/posener/eztables?status.svg)](http://godoc.org/github.com/posener/eztables)
[![Go Report Card](https://goreportcard.com/badge/github.com/posener/eztables)](https://goreportcard.com/report/github.com/posener/eztables)

Easy to understand web view of iptables rules

![screenshot](./screenshot.png "Screenshot")

## Run with Docker

```bash
docker run --rm --net host --privileged posener/eztables
```

> You should have docker installed, configured and running

## Install

```bash
go get -u github.com/posener/eztables
bash -c "sudo cp $(which eztables) /usr/bin/"
```

> `eztables` must run with root privileges since it runs `iptables` as a sub process.
> Therefore, I recommend copy the executable to `/usr/bin`.

## Usage

```bash
sudo eztables
```

> `eztables` should run with root privileges, since it runs `iptables` as a sub process.

