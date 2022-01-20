package dto

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
