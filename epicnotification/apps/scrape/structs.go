package scrape

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
