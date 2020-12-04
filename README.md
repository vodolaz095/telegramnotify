Telegram Notify
==============================
User friendly cross platform console application to send notifications (text, images, files) via telegram bots in proper way.


Why?
==============================

I think getting [group id](https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id) for 
telegram bot is not so straightforward. Also sending images, files, documents via bash script can 
be [complicated too](https://stackoverflow.com/questions/50414388/telegram-bot-cant-send-photo-from-disc-bash-script).
That's why i have made this application, to make things easier.


Install
==============================

Code was tested on Fedora 31 or Centos 8 linux. You can get [precompiled binaries here](https://github.com/vodolaz095/telegramnotify/releases/).
Put them into `/usr/bin/telegramnotify` or in `~/bin/telegramnotify` or anywhere in $PATH.
If you want to build code from sources, see below.

Setup
==============================

Start application in setup mode, follow instructions:

```

    $ telegramnotify setup
    Please, provide bot token, received from https://t.me/BotFather official bot.
    It will be something like 238222314:BAjcF4IKGAIiL.
    Press ENTER when you are ready:
    > 238222314:BAjcF4IKGAIiLBAjcF4IKGAIiLBAjcF4IKGAIiLBAjcF4IKGAIiL
    We have authorized as bot steellocalbot #238222314!
    Now please invite bot to groups, where it should post notifications...
    Bot steellocalbot (#238222334) was invited to chat OldCity (#-1906) of type group. It as us! Lets send message!
    Confirmation is send to group OldCity (#-1906)
    Sink OldCity is added to config at /home/vodolaz095/projects/telegramnotify/telegramnotify.json
    If you want to send messages to other groups, you can add them now!
    Press CTRL+C when you have finished adding groups to this bot...

```

It will create configuration file `telegramnotify.json` with format like this:

```json

{
  "work": {
    "token": "238222314:BAjcF4IKGAIiLBAjcF4IKGAIiLBAjcF4IKGAIiLBAjcF4IKGAIiL",
    "chatID": -19069
  },
  "Church": {
    "token": "238222314:BAjcF4F4IKGAIiLBAjcF4IKGAIiLBAjcF4IKGAIiL",
    "chatID": -19071
  },
  "misc": {
    "token": "238222312:IKGAIiLBAjcIKGAIiLBAjcIKGAIiLBAjc",
    "chatID": -19072
  }
}

``` 
We assume you have two bots one with token `238222314:BAjcF4IKGAIiLBAjcF4IKGAIiLBAjcF4IKGAIiLBAjcF4IKGAIiL`
invited to chats `work` with ID `-19069` and `Church` with id `-19071`. And second bot has token 
`238222312:IKGAIiLBAjcIKGAIiLBAjcIKGAIiLBAjc` and can only post messages to chat `misc` with id `-19072`.
Our program uses `work`, `Church` and `misc` as names of sink, where it can send notifications.
So, to send notification into `work` chat, we can invoke application like this one:

```

  $ telegramnotify send "Hello, this is test plain text message to be send to sink work!" work

```

Its worth notice, that we can send notifications via more than one bot to few different groups.

Config discovery
=============================

On linux machines, configuration file `telegramnotify.json` is discovered in this locations:

- /etc/telegramnotify.json
- ~/.config/telegramnotify.json
- file `telegramnotify.json` in current working directory

On other OSes config file is found using [os.UserConfigDir](https://godoc.org/os#UserConfigDir)  


Usage
=============================

In general, calling

```

    $ telegramnotify


```

will reveal online help, that is more reliable, than this documentation.

Something like this

```

    Console application to send notifications into telegram channels via bot-api
    
    Usage:
      telegramnotify [flags]
      telegramnotify [command]
    
    Available Commands:
      file        Upload file to sink provided
      help        Help about any command
      image       Upload image to sink provided
      list        list sinks
      setup       Perform initial setup
      test        Send test message to channel provided
      text        Send plain text to sink provided
    
    Flags:
      -c, --config string   path to config file (default "/home/vodolaz095/projects/telegramnotify/telegramnotify.json")
      -h, --help            help for telegramnotify
      -v, --verbose         verbose output
          --version         version for telegramnotify
    
    Use "telegramnotify [command] --help" for more information about a command.

```


Sending plain text messages:

```

    $ telegramnotify send \"Hello, this is test plain text message to be send to sink >>>work<<<!\" work
    $ telegramnotify send -m Markdown \"*Hello from telegramnotify*\nThis is markdown formatted notification.\" work
    $ telegramnotify send -m HTML '<a href=\"https://www.rt.com/\">Stay up to date with latest news!</a>' work

```

Different [text encoding options](https://tlgrm.ru/docs/bots/api#markdown-style) can be defined by `-m` parameter. 

We can pipe output of command into chat too:

```

   $ uptime | telegramnotify send work

```


Upload document as file:

```

   $  telegramnotify upload /etc/passwd work
   $  telegramnotify share /etc/passwd work
   $  telegramnotify file /etc/passwd work
   $  telegramnotify document /etc/passwd work

```

Upload image as photo:

```

   $  telegramnotify image /home/vodolaz095/Pictures/image.jpg work
   $  telegramnotify photo /home/vodolaz095/Pictures/image.png work
   $  telegramnotify pic /home/vodolaz095/Pictures/something.gif work

```


Requirements to build code
=============================

1. Ensure you have 64 bit linux running. Anatolij compiled code using either [Fedora LXDE](https://spins.fedoraproject.org/en/lxde) or [Centos 8](https://www.centos.org/download/)
2. Ensure you have [GNU make](https://www.gnu.org/software/make/) installed, usually `dnf install make`
3. Ensure you have [Golang](https://golang.org/dl/) at least of 1.12.x version installed (Anatolij used `go1.12.17 linux/amd64`), usually `dnf install golang golang-godoc` 
4. Ensure you have [UPX](http://upx.sourceforge.net/) installed, usually `dnf install upx`.

Golang configuration parameters
=========================

On Anatolij's machine this golang environment parameters are set (mainly, stack ones)

```

[vodolaz095@steel admin]$ go env
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/vodolaz095/.cache/go-build"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/home/vodolaz095/go"
GOPROXY=""
GORACE=""
GOROOT="/usr/lib/golang"
GOTMPDIR=""
GOTOOLDIR="/usr/lib/golang/pkg/tool/linux_amd64"
GCCGO="gccgo"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD="/home/vodolaz095/projects/secur3ltd/admin/go.mod"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build560136693=/tmp/go-build -gno-record-gcc-switches"

```

Anatolij uses Golang from official distro repository

```

[vodolaz095@steel admin]$ dnf info golang
Installed Packages
Name         : golang
Version      : 1.12.17
Release      : 1.fc30
Architecture : x86_64
Size         : 6.9 M
Source       : golang-1.12.17-1.fc30.src.rpm
Repository   : @System
From repo    : updates
Summary      : The Go Programming Language
URL          : http://golang.org/
License      : BSD and Public Domain
Description  : The Go Programming Language.

```

Also Anatolij can build code using stack golang package on Centos 8 linux:

```

[vodolaz095@holod admin]$ dnf info golang
Installed Packages
Name         : golang
Version      : 1.13.4
Release      : 2.module_el8.2.0+306+4f5ea1ce
Architecture : x86_64
Size         : 7.8 M
Source       : golang-1.13.4-2.module_el8.2.0+306+4f5ea1ce.src.rpm
Repository   : @System
From repo    : AppStream
Summary      : The Go Programming Language
URL          : http://golang.org/
License      : BSD and Public Domain
Description  : The Go Programming Language.


```

Building code
============================

```

    $ make build_prod

```

Binary will be compiled in `build/telegramnotify`

Releases
=============================

They will be published here [https://github.com/vodolaz095/telegramnotify/releases](https://github.com/vodolaz095/telegramnotify/releases)

License
=============================

The MIT License (MIT)

Copyright (c) 2020 Ostroumov Anatolij <ostroumov095 at gmail dot com>

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
