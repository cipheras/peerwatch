<div>
<img align="left" src="../assets/peerwatch.ico?raw=true" alt="PeerWatch" width="50px" height="50px" />

# PeerWatch &nbsp; ![GitHub release (latest by date)](https://img.shields.io/github/v/release/cipheras/peerwatch?style=flat-square&logo=superuser) 
  
</div>
  
#### A tool to stream videos directly into the media player without any hazzle by just name of the video. It also have UI version to make it easy for common users. To run UI version, just execute binary directly.

![GO version](https://img.shields.io/github/go-mod/go-version/cipheras/peerwatch?style=flat-square&color=blue)
&nbsp;&nbsp;&nbsp;&nbsp;
![GitHub All Releases](https://img.shields.io/github/downloads/cipheras/peerwatch/total?style=flat-square&color=darkgreen)
&nbsp;&nbsp;&nbsp;&nbsp;
![platform](https://img.shields.io/badge/dynamic/json?url=https://jsonkeeper.com/b/KNO7&label=platform&query=platform&style=flat-square&labelColor=grey&color=purple)

![gif](../assets/screen.gif?raw=true)


## Download
Download the tool from here:
Windows | Linux
--------|-------
[win-x64](https://github.com/cipheras/peerwatch/releases/download/v2.2/peerwatch-x64.zip) | [linux-x64](https://github.com/cipheras/peerwatch/releases/download/v2.2/peerwatch-x64)


## Building
You can use tool's pre-compiled binaries directly or you can compile from source.<br>
To build from source, GO must be installed.<br><br>
For linux installation:
```
sudo make
```
For linux build:
```
sudo make build
```
For linux uninstall:
```
sudo make uninstall
```
For windows:
***NOTE: For windows, DLLs are necessary to be kept in the same directory.***
```
go build
```


## Usage
**It may take some time to buffer as it depends on the file size and network speed.**
```
[+] ./peerwatch [name] or ./peerwatch "[name with multiple words]"
[+] Press ctrl + c to close

  -p int
        Port on which the tool will stream (default 8080)
  -player string
        Video player to play video [vlc | mpv | mplayer] (default "vlc")
  -tcp
        Connection over TCP or UDP (default true)
  -tp int
        Port on which the engine will work (default 50010)

Example:
peerwatch.exe xyz
or
./peerwatch "xyz blueray"
```

## Disclaimer
*This tool or author are not responsible for any type of copywrite claim. This tool work similarly as common user visiting the source and streaming videos on their browsers.*

## License
**peerwatch** is made by **@cipheras** and is released under the terms of the &nbsp;![GitHub License](https://img.shields.io/github/license/cipheras/peerwatch?color=darkgreen)


## Contact &nbsp; [![Twitter](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fcipheras%2Fpeerwatch&label=Tweet)](https://twitter.com/intent/tweet?text=Hi:&url=https%3A%2F%2Fgithub.com%2Fcipheras%2Fpeerwatch)
> Feel free to submit a bug, add features or issue a pull request.

