package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	// "github.com/todd-sudo/wb_go/src/structs"
)

func main() {
	fmt.Println("START")
	go saveProduct(1, 3)
	go saveProduct(3, 6)
	fmt.Println("END")
}

func saveProduct(startPage int, endPage int) {
	category := "/zhenshchinam/odezhda/bryuki-i-shorty"
	for startPage < endPage+1 {
		var details []*DetailProduct
		pageUrl := fmt.Sprintf("https://www.wildberries.ru/catalogdata%s?page=%s", category, strconv.Itoa(startPage))
		res := getRequest(pageUrl)

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Fatal(err)
		}
		ids := IDSModel{}

		jsonErr := json.Unmarshal(body, &ids)

		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		productsId := ids.Value.Data.Model.Products
		for _, productId := range productsId {

			detail := getDetailProduct(strconv.Itoa(productId.NmID))
			fmt.Println(detail)

			details = append(details, &detail)
		}

		// save data
		rawDataOut, err := json.MarshalIndent(&details, "", "  ")
		if err != nil {
			log.Fatal("JSON marshaling failed:", err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("data/data_%s.json", strconv.Itoa(startPage)), rawDataOut, 0777)
		if err != nil {
			log.Fatal("Cannot write updated settings file:", err)
		}
		startPage++
	}

}

func getDetailProduct(productID string) DetailProduct {
	urlDetail := fmt.Sprintf(
		"https://wbxcatalog-ru.wildberries.ru/nm-2-card/catalog?spp=3"+
			"&lang=ru&curr=rub&offlineBonus=0&onlineBonus=0&emp=0&locale=ru&nm=%s", productID,
	)
	res := getRequest(urlDetail)
	fmt.Println(res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	detail := DetailProduct{}

	jsonErr := json.Unmarshal(body, &detail)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return detail
}

func getRequest(url string) *http.Response {
	var netClient = http.Client{
		Timeout: time.Second * 1,
	}
	res, err := netClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

type DetailProduct struct {
	Data struct {
		Products []struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Brand      string `json:"brand"`
			BrandID    int    `json:"brandId"`
			SupplierID int    `json:"supplierId"`
			PriceU     int    `json:"priceU"`
			Sale       int    `json:"sale"`
			SalePriceU int    `json:"salePriceU"`
			Extended   struct {
				ClientSale   int `json:"clientSale"`
				ClientPriceU int `json:"clientPriceU"`
			} `json:"extended"`
			Rating    int `json:"rating"`
			Feedbacks int `json:"feedbacks"`
			Colors    []struct {
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"colors"`
			Sizes []struct {
				Name     string `json:"name"`
				OrigName string `json:"origName"`
				Rank     int    `json:"rank"`
				OptionID int    `json:"optionId"`
				Stocks   []struct {
					Wh  int `json:"wh"`
					Qty int `json:"qty"`
				} `json:"stocks"`
			} `json:"sizes"`
			DiffPrice bool `json:"diffPrice"`
		} `json:"products"`
	} `json:"data"`
}

type IDSModel struct {
	Value struct {
		Data struct {
			Model struct {
				Products []struct {
					NmID int `json:"nmId"`
				} `json:"products"`
			} `json:"model"`
		} `json:"data"`
	} `json:"value"`
}

type PagerStruct struct {
	Value struct {
		Data struct {
			Model struct {
				PagerModel struct {
					PagingInfo struct {
						TotalPages int `json:"totalPages"`
					} `json:"pagingInfo"`
				} `json:"pagerModel"`
				Products []struct {
					NmID int `json:"nmId"`
				} `json:"products"`
			} `json:"model"`
		} `json:"data"`
	} `json:"value"`
}
