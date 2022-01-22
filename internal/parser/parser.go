package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/wb_go/internal/dto"
)

// Проходит по страница категории и собирает данные
func saveProduct(startPage int, endPage int) {
	category := "/zhenshchinam/odezhda/bryuki-i-shorty"

	for startPage <= endPage+1 {
		var details []dto.DetailProduct

		fmt.Printf("page = %s\n", strconv.Itoa(startPage))

		pageUrl := fmt.Sprintf(
			"https://www.wildberries.ru/catalogdata%s?page=%s",
			category,
			strconv.Itoa(startPage),
		)
		res := getRequest(pageUrl)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Fatal(err)
		}
		ids := dto.IDSModel{}

		jsonErr := json.Unmarshal(body, &ids)

		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		productsId := ids.Value.Data.Model.Products
		fmt.Println(len(productsId))
		for _, productId := range productsId {
			go func() {
				detail := getDetailProduct(strconv.Itoa(productId.NmID))

				details = append(details, detail.Data.Products...)

				// save data
				rawDataOut, err := json.MarshalIndent(&details, "", "  ")
				if err != nil {
					log.Fatal("JSON marshaling failed:", err)
				}

				err = ioutil.WriteFile(fmt.Sprintf("data/data_%s.json", strconv.Itoa(startPage)), rawDataOut, 0777)
				if err != nil {
					log.Fatal("Cannot write updated settings file:", err)
				}
			}()
		}
		startPage++
		// time.Sleep(10 * time.Millisecond)
	}
	// return nil
}

// Получает детальную инфу о товаре
func getDetailProduct(productID string) dto.DetailProductData {
	urlDetail := fmt.Sprintf(
		"https://wbxcatalog-ru.wildberries.ru/nm-2-card/catalog?spp=3"+
			"&lang=ru&curr=rub&offlineBonus=0&onlineBonus=0&emp=0&locale=ru&nm=%s", productID,
	)
	res := getRequest(urlDetail)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	detail := dto.DetailProductData{}
	jsonErr := json.Unmarshal(body, &detail)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return detail
}
