package scrape

import "time"

type Headers struct {
	Accept          string `json:"accept"`
	AcceptLanguage  string `json:"accept-language"`
	SecChUa         string `json:"sec-ch-ua"`
	SecChUaMobile   string `json:"sec-ch-ua-mobile"`
	SecChUaPlatform string `json:"sec-ch-ua-platform"`
	SecFetchDest    string `json:"sec-fetch-dest"`
	SecFetchMode    string `json:"sec-fetch-mode"`
	SecFetchSite    string `json:"sec-fetch-site"`
	XRequestedWith  string `json:"x-requested-with"`
	Referer         string `json:"Referer"`
	ReferrerPolicy  string `json:"Referrer-Policy"`
}

var RequestHeaders = `{
		"accept": "application/json, text/plain, */*",
		"accept-language": "en-US,en;q=0.9",
		"sec-ch-ua": "\"Chromium\";v=\"112\", \"Google Chrome\";v=\"112\", \"Not:A-Brand\";v=\"99\"",
		"sec-ch-ua-mobile": "?0",
		"sec-ch-ua-platform": "\"macOS\"",
		"sec-fetch-dest": "empty",
		"sec-fetch-mode": "cors",
		"sec-fetch-site": "same-site",
		"x-requested-with": "XMLHttpRequest",
		"Referer": "https://store.epicgames.com/en-US/free-games",
		"Referrer-Policy": "no-referrer-when-downgrade"
	}`

type ResponseStruct struct {
	Data struct {
		Catalog struct {
			SearchStore struct {
				Elements []struct {
					Title                string    `json:"title"`
					ID                   string    `json:"id"`
					Namespace            string    `json:"namespace"`
					Description          string    `json:"description"`
					EffectiveDate        time.Time `json:"effectiveDate"`
					OfferType            string    `json:"offerType"`
					ExpiryDate           time.Time `json:"expiryDate"`
					ViewableDate         time.Time `json:"viewableDate"`
					Status               string    `json:"status"`
					IsCodeRedemptionOnly bool      `json:"isCodeRedemptionOnly"`
					KeyImages            []struct {
						Type string `json:"type"`
						URL  string `json:"url"`
					} `json:"keyImages"`
					Seller struct {
						ID   string `json:"id"`
						Name string `json:"name"`
					} `json:"seller"`
					ProductSlug string `json:"productSlug"`
					URLSlug     string `json:"urlSlug"`
					URL         string `json:"url"`
					Items       []struct {
						ID        string `json:"id"`
						Namespace string `json:"namespace"`
					} `json:"items"`
					CustomAttributes []struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"customAttributes"`
					Categories []struct {
						Path string `json:"path"`
					} `json:"categories"`
					Tags []struct {
						ID string `json:"id"`
					} `json:"tags"`
					CatalogNs struct {
						Mappings []struct {
							PageSlug string `json:"pageSlug"`
							PageType string `json:"pageType"`
						} `json:"mappings"`
					} `json:"catalogNs"`
					OfferMappings []struct{} `json:"offerMappings"`
					Price         struct {
						TotalPrice struct {
							DiscountPrice   int    `json:"discountPrice"`
							OriginalPrice   int    `json:"originalPrice"`
							VoucherDiscount int    `json:"voucherDiscount"`
							Discount        int    `json:"discount"`
							CurrencyCode    string `json:"currencyCode"`
							CurrencyInfo    struct {
								Decimals int `json:"decimals"`
							} `json:"currencyInfo"`
							FmtPrice struct {
								OriginalPrice     string `json:"originalPrice"`
								DiscountPrice     string `json:"discountPrice"`
								IntermediatePrice string `json:"intermediatePrice"`
							} `json:"fmtPrice"`
						} `json:"totalPrice"`
						LineOffers []struct {
							AppliedRules []struct {
								ID              string    `json:"id"`
								EndDate         time.Time `json:"endDate"`
								DiscountSetting struct {
									DiscountType string `json:"discountType"`
								} `json:"discountSetting"`
							} `json:"appliedRules"`
						} `json:"lineOffers"`
					} `json:"price"`
					Promotions struct {
						PromotionalOffers []struct {
							PromotionalOffers []struct {
								StartDate       time.Time `json:"startDate"`
								EndDate         time.Time `json:"endDate"`
								DiscountSetting struct {
									DiscountType       string `json:"discountType"`
									DiscountPercentage int    `json:"discountPercentage"`
								} `json:"discountSetting"`
							} `json:"promotionalOffers"`
						} `json:"promotionalOffers"`
						UpcomingPromotionalOffers []struct{} `json:"upcomingPromotionalOffers"`
					} `json:"promotions"`
				} `json:"elements"`
			} `json:"searchStore"`
		} `json:"Catalog"`
	} `json:"data"`
	Extensions struct{} `json:"extensions"`
}
