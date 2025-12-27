package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	
	"golang.org/x/net/proxy"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Adresler []string `yaml:"adresler"`
}

func main() {
	logFile, err := os.Create("scan_report.log")
	if err != nil {
		fmt.Println("HATA: Rapor dosyası oluşturulamadı:", err)
		return
	}
	defer logFile.Close()

	dialer, _ := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
	client := &http.Client{
		Transport: &http.Transport{Dial: dialer.Dial},
		Timeout:   30 * time.Second,
	}

	ipResp, err := client.Get("http://check.torproject.org/api/ip")
	if err == nil {
		body, _ := io.ReadAll(ipResp.Body)
		fmt.Println("Tor IP:", strings.TrimSpace(string(body)))
		ipResp.Body.Close()
	}

	data, _ := os.ReadFile("targets.yaml")
	var cfg Config
	yaml.Unmarshal(data, &cfg)

	fmt.Fprintln(logFile, "--- TARAMA RAPORU ---")

	for _, url := range cfg.Adresler {
		fmt.Printf("Taraniyor: %s ", url)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("-> başaramadık abi %v\n", err)
			fmt.Fprintf(logFile, "Çalışmıyor. %s | Hata: %v\n", url, err)
			continue
		}

		fmt.Println("-> BASARILI")
		fmt.Fprintf(logFile, "Çalışıyor. %s | Kod: %d\n", url, resp.StatusCode)

		name := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(url, "http://", ""), "https://", ""), "/", "_") + ".html"
		f, _ := os.Create(name)
		io.Copy(f, resp.Body)
		f.Close()
		resp.Body.Close()
	}
}
