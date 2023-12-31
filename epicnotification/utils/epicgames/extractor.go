package epicgames

import (
	"log"
	"time"

	"epicnotification/epicnotification/helpers"
	urls "epicnotification/epicnotification/utils"
)

type Sanitized struct {
	Name          string `json:"name"`
	EffectiveDate string `json:"effectiveData"`
	OfferType     string `json:"offerType"`
	Status        string `json:"status"`
	StartDate     string `json:"startDate"`
	ProductLink   string `json:"productLink"`
	OriginalPrice string `json:"originalPrice"`
	DiscountPrice string `json:"discountPrice"`
	Free          bool   `json:"Free"`
}

func ExtractData(response helpers.ApiResponse) []Sanitized {
	// var sanitizedResponse Sanitized
	elements := response.Data.Catalog.SearchStore.Elements
	log.Println("count of game: ", len(elements))

	var resp []Sanitized
	var baseUrl = urls.BaseURLs["epic-games"]

	for index, val := range elements {
		resp = append(resp, Sanitized{
			Name:          val.Title,
			EffectiveDate: isoToString(val.EffectiveDate),
			OfferType:     val.OfferType,
			Status:        val.Status,
			StartDate:     startDate(index, response),
			ProductLink:   createURL(index, baseUrl, response),
			OriginalPrice: val.Price.TotalPrice.FmtPrice.OriginalPrice,
			DiscountPrice: val.Price.TotalPrice.FmtPrice.DiscountPrice,
			Free:          isFree(val.Price.TotalPrice.OriginalPrice, val.Price.TotalPrice.DiscountPrice),
		})
	}

	log.Print("extracted data: ", resp)
	return resp
}

func isoToString(date time.Time) string {
	log.Println("converting date to string: ", date)
	return date.Format("2006-01-02 15:04:05")
}

func startDate(index int, d helpers.ApiResponse) string {
	element := d.Data.Catalog.SearchStore.Elements[index]
	log.Println("converting start date to string", element)
	return isoToString(element.ViewableDate)
}

func isFree(op int, dp int) bool {
	return op == dp
}

func createURL(index int, baseURL string, d helpers.ApiResponse) string {
	log.Println("Creating Url...")
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
