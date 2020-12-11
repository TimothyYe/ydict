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

#### Windows x64

>   `Windows` 下语音朗读适配方案由 [ycrao](https://github.com/ycrao/learning_golang/tree/main/cmd-bass-player) 提供。

- 拷贝 `bass.dll` 和 `mpg123.exe` （也可同 `ydict.exe`） 文件一起 到 Windows 系统目录 (如 `C:\Windows\` 或 `C:\Windows\System32`)。
- 或者拷贝 `bass.dll` 和 `mpg123.exe` （也可同 `ydict.exe`） 文件一起文件到某一特定目录下，然后添加该目录路径到 `PATH` 系统环境变量中。


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

1. 仅查询单词

```text
ydict <要查询的单词或词组>
```

2. 查询并朗读单词

```text
ydict -v <要查询的单词或词组>
```
3. 查询并显示更多例句

```text
ydict -m <要查询的单词或词组>
```

4. 开启本地缓存查询单词，单词将被缓存在本地缓存中，再一次使用相同参数查询相同的单词，将获得更快的显示速度

```text
ydict -c <要查询的单词或词组>
```

5. 查询整个句子

```text
ydict -s "你觉得咋样？"
```

## SOCKS5 代理支持

从版本 V0.5 开始, 支持SOCKS5代理功能. 在ydict的相同目录下，创建 ```.env``` 文件，并填入如下示例内容:

```text
SOCKS5=127.0.0.1:7070
```

配置成功后，所有的查询将使用配置指定的SOCKS5代理。

## 单词本功能

从新版 ydict V2.0 开始，支持单词本功能，方便增删和管理生词，并且可以通过定时消息推送进行随机回放，方便背诵和记忆。

* 增加新词到单词本
```bash
ydict -c hello
```

* 从单词本中删除单词
```bash
ydict -d hello
```

* 从单词本中列出所有单词
```bash
ydict -l
```

* 每隔10秒随机推送并展示单词
```bash
ydict -p 10
```
![](https://raw.githubusercontent.com/TimothyYe/ydict/master/snapshots/play.png)

## 帮助与更多信息

命令行中，输入 "ydict" 获取更多帮助。
  
## 开源协议

[MIT License](https://github.com/TimothyYe/ydict/blob/master/LICENSE)
