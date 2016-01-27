# tools for binaries reverse engineering

Cross-platform cli tools and packages for binaries reverse engineering.

## Installation

To install bre, please use `go get`.

### ldd like tool

```
$ go get github.com/ArtemKulyabin/bre/cmd/bre-ldd
```

### objdump like tool

```
$ go get github.com/ArtemKulyabin/bre/cmd/bre-dump
```

## Usage

### Command line tool

```
$ bre-ldd a.out
...
$ bre-dump a.out
...
```

## Operating system support
* Linux, FreeBSD, OpenBSD, NetBSD, Darwin(OS X), Windows

### Tested on
* Ubuntu 14.04, FreeBSD 10.1, OpenBSD 5.6, NetBSD 6.1.5, OS X Yosemite 10.10.3, Windows 7
