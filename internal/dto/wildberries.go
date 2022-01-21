package dto

type DetailProduct struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Brand      string `json:"brand"`
	BrandID    int    `json:"brandId"`
	SupplierID int    `json:"supplierId"`
	PriceU     int    `json:"priceU"`
	Sale       int    `json:"sale"`
	SalePriceU int    `json:"salePriceU"`
	Extended   struct {
		BasicSale    int `json:"basicSale"`
		BasicPriceU  int `json:"basicPriceU"`
		PromoSale    int `json:"promoSale"`
		PromoPriceU  int `json:"promoPriceU"`
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
}

type DetailProductData struct {
	Data struct {
		Products []DetailProduct `json:"products"`
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
