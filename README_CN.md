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
[3]: https://img.shields.io/badge/release-v0.9-brightgreen.svg
[4]: https://github.com/TimothyYe/ydict/releases
[5]: https://img.shields.io/dub/l/vibe-d.svg
[6]: LICENSE
[7]: https://goreportcard.com/badge/github.com/timothyye/ydict
[8]: https://goreportcard.com/report/github.com/timothyye/ydict

Ydict, 专为命令行极客打造的有道词典!

![](https://raw.githubusercontent.com/TimothyYe/ydict/master/snapshots/ydict.gif)

## 功能一览

* 中文翻译为英文
* 英文翻译为中文
* 查询不到单词时，自动显示推荐搜索提示
* 语音朗读功能，朗读你所查询的单词

## 安装

#### Homebrew

```bash
brew tap timothyye/tap
brew install timothyye/tap/ydict
```

#### 使用go get安装

```bash
go get github.com/TimothyYe/ydict
```

#### 手动安装

从 [releases](https://github.com/TimothyYe/ydict/releases) 下载最新发布版本, 解压可执行文件到 /usr/bin

#### Vim集成插件

使用配套Vim插件，可以直接在Vim中通过ydict查询单词。安装与配置，请移步: [vim-ydict](https://github.com/TimothyYe/vim-ydict)

## 语音朗读功能

从V0.9版本开始，提供语音朗读功能。为开启此功能，你需要先安装mpg123组件。

___注意:___ 语音朗读功能当前仅支持操作系统 MacOS/Linux。

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

## 使用的正确姿势

1. 仅查询单词

```text
ydict <要查询的单词或词组>
```

2. 查询并朗读单词 (__目前仅支持 MacOS 和 Linux__)

```text
ydict <要查询的单词或词组> -v
```

## SOCKS5 代理支持

从版本 V0.5 开始, 支持SOCKS5代理功能. 在ydict的相同目录下，创建 ```.env``` 文件，并填入如下示例内容:

```text
SOCKS5=127.0.0.1:7070
```

配置成功后，所有的查询将使用配置指定的SOCKS5代理。

## 帮助与更多信息

命令行中，输入 "ydict" 获取更多帮助。
  
## 开源协议

[MIT License](https://github.com/TimothyYe/ydict/blob/master/LICENSE)
