# PeerWatch &nbsp; ![GitHub release (latest by date)](https://img.shields.io/github/v/release/cipheras/peerwatch?style=flat-square&logo=superuser)
#### A tool to stream videos directly into the VLC media player. 

![Lines of code](https://img.shields.io/tokei/lines/github.com/cipheras/peerwatch?style=flat-square)
&nbsp;&nbsp;&nbsp;&nbsp;![GO version](https://img.shields.io/github/go-mod/go-version/cipheras/peerwatch?style=flat-square&color=green)
&nbsp;&nbsp;&nbsp;&nbsp;![GitHub All Releases](https://img.shields.io/github/downloads/cipheras/peerwatch/total?style=flat-square)
&nbsp;&nbsp;&nbsp;&nbsp;![platform](https://img.shields.io/badge/dynamic/json?url=https://jsonkeeper.com/b/KNO7&label=platform&query=platform&style=flat-square&labelColor=grey&color=purple)

![example](../asset/screen.gif?raw=true)


## Download
Download the tool from here:
Windows | Linux
--------|-------
[win-x64](https://github.com/cipheras/peerwatch/releases/download/v1.0/peerwatch-x64.exe) | [linux-x64](https://github.com/cipheras/peerwatch/releases/download/v1.0/peerwatch-x64)


## Building
You can use tool's pre-compiled binaries directly or you can compile from source.<br>
To build from source, GO must be installed.<br>
For linux installation:
```sudo make```
For linux build:
```sudo make build```
For linux uninstall:
```sudo make uninstall```
For windows:
```go build```


## Usage
**It may take some time to buffer as it depends on the file size and network speed.**
```
./peerwatch [-h] [-p port] [name]

-p     port on which tool will run
name   name of the movie(use double quotes for multi words name)

Example:
peerwatch.exe xyz
or
./peerwatch "xyz blueray"
```

## Disclaimer
*This tool or author are not responsible for any type of copywrite claim. This tool work similarly as common user visiting the source and stream videos on their browsers.*

## License
**peerwatch** is made by **@cipheras** and is released under the terms of the &nbsp;![GitHub License](https://img.shields.io/github/license/cipheras/peerwatch?color=darkgreen)


## Contact &nbsp; [![Twitter](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fcipheras%2Fpeerwatch&label=Tweet)](https://twitter.com/intent/tweet?text=Hi:&url=https%3A%2F%2Fgithub.com%2Fcipheras%2Fpeerwatch)
> Feel free to submit a bug, add features or issue a pull request.

