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
[3]: http://github-release-version.herokuapp.com/github/timothyye/ydict/release.svg?style=flat
[4]: https://github.com/TimothyYe/ydict/releases
[5]: https://img.shields.io/dub/l/vibe-d.svg
[6]: LICENSE
[7]: https://goreportcard.com/badge/github.com/timothyye/ydict
[8]: https://goreportcard.com/report/github.com/timothyye/ydict

Ydict, another command-line youdao dictionary for geeks!

![](https://raw.githubusercontent.com/TimothyYe/ydict/master/snapshots/ydict.gif)

([中文介绍文档](https://github.com/TimothyYe/ydict/blob/master/README_CN.md))

## Features

* Chinese -> English
* English -> Chinese
* Show hints if word is not found
* Speech
* Show example sentences
* Vim support

## Installation

#### Homebrew

```bash
brew tap timothyye/tap
brew install timothyye/tap/ydict
```

#### Using Go

```bash
go get github.com/TimothyYe/ydict
```

#### Manual Installation

Download it from [releases](https://github.com/TimothyYe/ydict/releases), and extact it to /usr/bin.

#### Integrate with Vim

To query words from Vim, you need another Vim plugin: [vim-ydict](https://github.com/TimothyYe/vim-ydict)

## Speech

Starting from V0.9, speech feature is available. You need to install mpg123 to enable this feature.

___NOTICE:___ Currently, speech feature is only available for MacOS/Linux.

#### Mac OS

```bash
brew install mpg123
```
#### Ubuntu

```bash
sudo apt-get install mpg123
```

#### CentOS

```bash
yum install -y mpg123
```

## Usage

1. Query

```text
ydict <word(s) to query>
```

2. Query with speetch (__Available for MacOS & Linux__)

```text
ydict <word(s) to query> -v
```

3. Query and show more example sentences

```text
ydict <word(s) to query> -m
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
