```text
██╗   ██╗██████╗ ██╗ ██████╗████████╗
╚██╗ ██╔╝██╔══██╗██║██╔════╝╚══██╔══╝
 ╚████╔╝ ██║  ██║██║██║        ██║   
  ╚██╔╝  ██║  ██║██║██║        ██║   
   ██║   ██████╔╝██║╚██████╗   ██║   
   ╚═╝   ╚═════╝ ╚═╝ ╚═════╝   ╚═╝   
 ```

[![Release][3]][4] [![MIT licensed][5]][6] [![Build Status][1]][2] [![Go Report Card][7]][8]

[1]: https://travis-ci.org/TimothyYe/ydict.svg?branch=master
[2]: https://travis-ci.org/TimothyYe/ydict
[3]: https://img.shields.io/badge/release-v0.7-brightgreen.svg
[4]: https://github.com/TimothyYe/ydict/releases
[5]: https://img.shields.io/dub/l/vibe-d.svg
[6]: LICENSE
[7]: https://goreportcard.com/badge/github.com/timothyye/ydict
[8]: https://goreportcard.com/report/github.com/timothyye/ydict

Ydict, another command line dictionary for geeks!

![](https://raw.githubusercontent.com/TimothyYe/ydict/master/snapshots/ydict.gif)

## Features

* Chinese -> English
* English -> Chinese

## Installation

#### Homebrew

```bash
brew tap timothyye/tap
brew install timothyye/tap/ydict                                                                                                                                      [08297
```

#### Using Go

```bash
go get github.com/TimothyYe/ydict
```

#### Manual Installation

Download it from [releases](https://github.com/TimothyYe/ydict/releases), and extact it to /usr/bin.

## Usage

```text
ydict <word to query>
```

## SOCKS5 proxy

Starting from V0.5, you can use SOCKS5 proxy. At the same directory of ydict, just create a ```.env``` file:

```text
SOCKS5=127.0.0.1:7070
```

Now all the queries will go through the specified SOCKS5 proxy.

## Help

Just type "ydict" to get help.
  
## Licence

[MIT License](https://github.com/TimothyYe/ydict/blob/master/LICENSE)
