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

Download it from [releases](https://github.com/TimothyYe/ydict/releases), and extract  it to /usr/bin.

#### Integrate with Vim

To query words from Vim, you need another Vim plugin: [vim-ydict](https://github.com/TimothyYe/vim-ydict)

## Speech

Starting from V0.9, speech feature is available. You need to install mpg123 to enable this feature.

#### Windows x64

>   Speech Adaptation by [ycrao](https://github.com/ycrao/learning_golang/tree/main/cmd-bass-player) under `Windows` OS.

- Just copy `bass.dll` and `mpg123.exe` (can also with `ydict.exe`) file to `Windows` system path (such as `C:\Windows\` or `C:\Windows\System32` ) .
- Or copy `bass.dll` and `mpg123.exe` (can also with `ydict.exe`) file to somewhere in the same directory, and setting that directory in `PATH` System Environment Variables .

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

```text
ydict [flags]

Flags:
  -c, --cache       Query with local cache, and save the query word(s) into the cache.
  -d, --delete      Remove word(s) from the cache.
  -h, --help        help for ydict
  -l, --list        List all the words from the local cache.
  -m, --more        Query with more example sentences.
  -p, --play int    Scan and display all the words in local cache.
  -q, --quiet       Query with quiet mode, don't show spinner.
  -r, --reset       Clear all the words from the local cache.
  -s, --sentence    Translation of sentences.
  -v, --voice int   Query with voice speech, the default voice play count is 0.
```

1. Query

```text
ydict <word(s) to query>
```

2. Query with speech

```text
ydict -v 1 <word(s) to query>
```

3. Query and show more example sentences

```text
ydict -m <word(s) to query>
```

4. Query and add this word into local cache, next time when you query the same word, it will be feched from the local cache and be much more faster.

```text
ydict -c <word(s) to query>
```

5. Query sentence

```text
ydict -s "你觉得咋样？"
```

## SOCKS5 proxy

Starting from V0.5, you can use SOCKS5 proxy. At the same directory of ydict, just create a `.env` file:

```text
SOCKS5=127.0.0.1:7070
```

Now all the queries will go through the specified SOCKS5 proxy.

## New words notebook

Starting from ydict V2.0, new words notebook is supported. You can use is to add/delete your new words and play it.

* Add a new word to the notebook
```bash
ydict -c hello
```

* Remove a word from the notebook
```bash
ydict -d hello
```

* List all the words from the notebook
```bash
ydict -l
```

* Display a random word from the notebook for every 10 seconds
```bash
ydict -p 10
```
![](https://raw.githubusercontent.com/TimothyYe/ydict/master/snapshots/play.png)

## Help

Just type "ydict" to get help.
  
## Licence

[MIT License](https://github.com/TimothyYe/ydict/blob/master/LICENSE)
