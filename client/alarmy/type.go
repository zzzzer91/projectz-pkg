package alarmy

type GetHoroscopeResp struct {
	Date        string `json:"date"`
	Zodiac      int    `json:"zodiac"`
	ShortDesc   string `json:"shortDesc"`
	LongDesc    string `json:"longDesc"`
	LuckyColor  string `json:"luckyColor"`
	LuckyNumber int    `json:"luckyNumber"`
}
