package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/wb_go/internal/dto"
)

// Получает кол-во страниц по категории
func getCountPage(category string) int {
	url := fmt.Sprintf(
		"https://www.wildberries.ru/catalogdata%s?page=2",
		category,
	)
	res := getRequest(url)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	countPage := dto.PagerStruct{}
	jsonErr := json.Unmarshal(body, &countPage)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return countPage.Value.Data.Model.PagerModel.PagingInfo.TotalPages

}
