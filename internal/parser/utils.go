package parser

import (
	"log"
	"net"
	"net/http"
	"time"
)

func getRequest(url string) *http.Response {
	// var proxy = "http://" + user + ":" + password + "@" + ip + ":" + port

	// //creating the proxyURL
	// proxyURL, err := proxy.Parse(proxy)

	// if err != nil {
	// 	log.Println(err)
	// }

	// transport := &http.Transport{
	// 	Proxy: http.ProxyURL(proxyURL),
	// }
	http.DefaultClient.Timeout = time.Minute * 1
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 0 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 60 * time.Second,
	}

	var HttpClient = &http.Client{
		Transport: transport,
		// Timeout: time.Second * 1,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Linux; Android 10; SAMSUNG SM-A205FN) "+
			"AppleWebKit/537.36 (KHTML, like Gecko) "+
			"SamsungBrowser/15.0 "+
			"Chrome/90.0.4430.210 Mobile Safari/537.36",
	)
	req.Close = true

	resp, err := HttpClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

// Проверят кол-во страниц и возвращает кол-во горутин
func checkCountPage(countPage int) int {
	if countPage <= 100 {
		return 1
	} else if countPage > 100 && countPage <= 300 {
		return 3
	} else if countPage > 300 && countPage <= 500 {
		return 5
	} else if countPage > 500 && countPage <= 700 {
		return 7
	} else if countPage > 700 && countPage <= 900 {
		return 9
	}

	return 11
}
