package parser

import (
	"log"
	"net/http"
)

func getRequest(url string) *http.Response {
	var netClient = http.Client{
		// Timeout: time.Second * 1,
	}
	res, err := netClient.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	return res
}

// Проверят кол-во страниц и возвращает кол-во горутин
func checkCountPage(countPage int) int {
	if countPage < 10 {
		return 1
	} else if countPage >= 10 && countPage <= 30 {
		return 3
	} else if countPage >= 30 && countPage <= 60 {
		return 5
	} else if countPage >= 60 && countPage <= 90 {
		return 7
	}

	return 9
}
