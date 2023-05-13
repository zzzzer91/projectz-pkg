package alarmy

const (
	urlHoroscope = "https://api.alar.my/horoscope?zodiac=%s&timezone=%s&language=zh-Hans"
)

var (
	Constellations = map[string]string{
		"aries":       "白羊座", // 白羊座（3月21日～4月20日）
		"taurus":      "金牛座", // 金牛座（4月21～5月21日）
		"gemini":      "双子座", // 双子座（5月22日～6月21日）
		"cancer":      "巨蟹座", // 巨蟹座（6月22日～7月22日）
		"leo":         "狮子座", // 狮子座（7月23日～8月23日）
		"virgo":       "处女座", // 处女座（8月24日～9月23日）
		"libra":       "天秤座", // 天秤座（9月24日～10月23日）
		"scorpio":     "天蝎座", // 天蝎座（10月24日～11月22日）
		"sagittarius": "射手座", // 射手座（11月23日～12月21日）
		"capricorn":   "摩羯座", // 摩羯座（12月22日～1月20日）
		"aquarius":    "水瓶座", // 水瓶座（1月21日～2月19日）
		"pisces":      "双鱼座", // 双鱼座（2月20日～3月20日）
	}

	ColorsZh = map[string]string{
		"white":  "白色",
		"black":  "黑色",
		"purple": "紫色",
		"orange": "橙色",
		"red":    "红色",
		"blue":   "蓝色",
		"green":  "绿色",
		"yellow": "黄色",
		"gray":   "灰色",
		"pink":   "粉色",
		"brown":  "棕色",
	}
)
