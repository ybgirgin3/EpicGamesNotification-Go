package utils

import (
	"time"

	"epicnotification/epicnotification/helpers"
)

type Sanitized struct {
	Name          string `json:"name"`
	EffectiveDate string `json:"effectiveData"`
	OfferType     string `json:"offerType"`
	StartDate     string `json:"startDate"`
	ProductLink   string `json:"productLink"`
	OriginalPrice string `json:"originalPrice"`
	DiscountPrice string `json:"discountPrice"`
	Free          bool   `json:"Free"`
}

func ExtractData(response helpers.ApiResponse) []Sanitized {
	// var sanitizedResponse Sanitized
	elements := response.Data.Catalog.SearchStore.Elements

	var resp []Sanitized
	var baseUrl string = "https://store.epicgames.com/en-US/p"

	for index, val := range elements {

		resp = append(resp, Sanitized{
			Name:          val.Title,
			EffectiveDate: isoToString(val.EffectiveDate),
			OfferType:     val.OfferType,
			StartDate:     startDate(index, response),
			ProductLink:   createLink(index, baseUrl, response),
			OriginalPrice: val.Price.TotalPrice.FmtPrice.OriginalPrice,
			DiscountPrice: val.Price.TotalPrice.FmtPrice.DiscountPrice,
			Free:          isFree(val.Price.TotalPrice.OriginalPrice, val.Price.TotalPrice.DiscountPrice),
		})
	}

	return resp
}

func isoToString(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

func startDate(index int, d helpers.ApiResponse) string {
	data := d.Data.Catalog.SearchStore.Elements[index]
	return isoToString(data.ViewableDate)
}

func isFree(op int, dp int) bool {
	return op == dp
}

func createLink(index int, baseURL string, d helpers.ApiResponse) string {
	data := d.Data.Catalog.SearchStore.Elements[index]

	var slug interface{}

	if data.ProductSlug != nil {
		slug = data.ProductSlug
	} else {
		slug = data.OfferMappings[0].PageSlug
	}

	if str, ok := slug.(string); ok {
		return baseURL + "/" + str
	} else {
		return baseURL
	}
}
