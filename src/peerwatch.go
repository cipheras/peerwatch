package src

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/webview/webview"
)

const (
	VERSION = "v2.0"
	RESET   = "\033[0m"
	RED     = "\033[31m"
	GREEN   = "\033[32m"
	CYAN    = "\033[36m"
	PURPLE  = "\033[35m"
	BOLD    = "\033[1m"
	BLINK   = "\033[5m"
)

var (
	port         = flag.Int("p", 8080, "Port on which the tool will stream")
	tport        = flag.Int("tp", 50010, "Port on which the engine will work")
	tcp          = flag.Bool("tcp", true, "Connection over TCP or UDP")
	player       = flag.String("player", "vlc", "Video player to play video [vlc | mpv | mplayer]")
	link_regex   = "torrent/[0-9]{7}/[a-zA-Z0-9?%-]*/"
	magnet_regex = "magnet:\\?xt=urn:btih:[a-zA-Z0-9]*"
	query        string
	client       Client
	w            webview.WebView
	err          error
	ch           chan string
)

func Peerwatch(ln net.Listener) {
	windw()
	interrupt()
	banner()
	checkUpdates()
	flag.Usage = func() {
		fmt.Print(GREEN + "\n----------------------------------" + PURPLE + " PeerWatch " + GREEN + "----------------------------------\n" + RESET)
		fmt.Print(BOLD + CYAN + "[+] ./peerwatch " + RED + "[name]" + CYAN + " or ./peerwatch \"" + RED + "[name with multiple words]" + CYAN + "\"\n" + RESET)
		fmt.Print(BOLD + CYAN + "[+] Press " + RED + "ctrl + c" + CYAN + " to close\n\n" + RESET)
		flag.PrintDefaults()
		fmt.Print(GREEN + "----------------------------------" + PURPLE + " PeerWatch " + GREEN + "----------------------------------\n\n" + RESET)
	}
	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		// Run GUI as a thread
		ch = make(chan string)
		go gui(ln)
		query = <-ch
	} else {
		query = flag.Arg(0)
	}
	name, magnet := find(query)
	// Found or Not
	if magnet == "" {
		fmt.Println(BOLD + RED + "[-] " + BLINK + "FIle Not Found" + RESET + BOLD + RED + "...Try with another name" + RESET)
		w.Eval(`document.getElementById("status").innerText="File Not Found!!"`)
		os.Exit(0)
	}
	w.Eval(fmt.Sprintf("document.getElementById('status').innerText='%s'", strings.Split(name, "/")[2]))
	// w.Destroy()

	// Engine config
	cfg := ClientConfig{
		TorrentPath:    magnet,
		Port:           *port,
		TorrentPort:    *tport,
		Seed:           false,
		TCP:            *tcp,
		MaxConnections: 200,
	}
	engine(cfg)
}

func callback(uiname string) {
	ch <- uiname
}

func gui(ln net.Listener) {
	debug := false
	w = webview.New(debug)
	defer w.Destroy()
	w.SetTitle("PeerWatch")
	w.SetSize(480, 320, webview.HintNone)
	w.Bind("callback", callback)
	w.Navigate(fmt.Sprintf("http://%s/www/", ln.Addr()))
	w.Run()
}

func find(query string) (string, string) {
	var magnet, name string
	// Get searches from query
	url := strings.TrimSpace(strings.ReplaceAll(fmt.Sprintf("https://1337x.wtf/search/%v/1/", query), " ", "+"))
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("unable to connect")
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("unable to read defined search")
		}
		// regex find link
		re := regexp.MustCompile(link_regex)
		name = re.FindString(string(body))
	}

	// Get magnet links from top search result
	get_magnet := fmt.Sprintf("https://1337x.wtf/%v", name)
	get_magnet_resp, err := http.Get(get_magnet)
	if err != nil {
		log.Fatalf("unable to connect")
	}
	defer get_magnet_resp.Body.Close()
	if get_magnet_resp.StatusCode == 200 {
		magnet_body, err := ioutil.ReadAll(get_magnet_resp.Body)
		if err != nil {
			log.Fatalf("unable to read defined result")
		}
		// regex find magnet
		re := regexp.MustCompile(magnet_regex)
		magnet = re.FindString(string(magnet_body))
	}
	return name, magnet
}

func engine(cfg ClientConfig) {
	// Start torrent client
	client, err = NewClient(cfg)
	if err != nil {
		log.Fatalf(err.Error())
		os.Exit(0)
	}
	// Http handler
	go func() {
		http.HandleFunc("/", client.GetFile)
		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(cfg.Port), nil))
	}()
	// Open VLC player
	go func() {
		for !client.ReadyForPlayback() {
			time.Sleep(time.Second)
		}
		openPlayer(*player, cfg.Port)
	}()
	// Cli render
	for {
		client.Render()
		time.Sleep(time.Second)
	}
}

func banner() {
	bb := BOLD + GREEN + `
8888888b.                            888       888          888             888      
888   Y88b                           888   o   888          888             888      
888    888                           888  d8b  888          888             888      
888   d88P  .d88b.   .d88b.  888d888 888 d888b 888  8888b.  888888  .d8888b 88888b.    ` + PURPLE + `
8888888P"  d8P  Y8b d8P  Y8b 888P"   888d88888b888     "88b 888    d88P"    888 "88b 
888        88888888 88888888 888     88888P Y88888 .d888888 888    888      888  888   ` + GREEN + ` 
888        Y8b.     Y8b.     888     8888P   Y8888 888  888 Y88b.  Y88b.    888  888 
888         "Y8888   "Y8888  888     888P     Y888 "Y888888  "Y888  "Y8888P 888  888   ` + RED + `
====================================================================================   ` + CYAN + `
Created by:				                       Version: -~{ ` + RESET + RED + VERSION + BOLD + CYAN + ` }~-` + GREEN + `			    
+-+ +-+ +-+ +-+ +-+ +-+ +-+ +-+	
|C| |i| |p| |h| |e| |r| |a| |s|
+-+ +-+ +-+ +-+ +-+ +-+ +-+ +-+	                                                                                
` + RESET
	fmt.Print(bb + "\n")
}

func checkUpdates() {
	resp, err := http.Get("https://api.github.com/repos/cipheras/peerwatch/releases")
	if err != nil {
		log.Fatalln("unable to check updates:", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("unable to check updates")
	}
	var jsondata []map[string]interface{}
	err = json.Unmarshal(data, &jsondata)
	if err != nil {
		log.Fatalln("error parsing update data")
	}
	version := jsondata[0]["tag_name"].(string)
	releasName := jsondata[0]["name"].(string)
	if version != VERSION {
		fmt.Println(BOLD + CYAN + "[+] Update available..." + BLINK + GREEN + version + " [" + releasName + "]" + RESET)
	}
}

func windw() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Println("windows color not working\n", err)
		}
		fmt.Printf("\033[2J\033[H")
	}
}

func interrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		os.Interrupt,
	)
	go func(c chan os.Signal) {
		<-c
		fmt.Print("\r")
		log.Println("Exiting...!!")
		time.Sleep(time.Second)
		os.Exit(0)

	}(c)
}
