package hoyoverse

const HonkaiStarrailRedemptionURL = "https://public-operation-hkrpg.hoyoverse.com/common/apicdkey/api/webExchangeCdkeyRisk"

type RedemptionRequest struct {
	CDKey      string `json:"cdkey"`
	DeviceUUID string `json:"device_uuid"`
	GameBiz    string `json:"game_biz"`
	Lang       string `json:"lang"`
	Platform   string `json:"platform"`
	Region     string `json:"region"`
	Time       int64  `json:"t"`
	UID        string `json:"uid"`
}

type RedemptionResponse APIResponse[struct{}]

// https://hsr.hoyoverse.com/gift
var RedemptionHeaders = map[string]string{
	"Accept":             "*/*",
	"Accept-Language":    "en",
	"Origin":             "https://hsr.hoyoverse.com",
	"Priority":           "u=1, i",
	"Referer":            "https://hsr.hoyoverse.com/",
	"Sec-Ch-Ua":          `"Chromium";v="134", "Not:A-Brand";v="24", "Microsoft Edge";v="134"`,
	"Sec-Ch-Ua-Mobile":   "?0",
	"Sec-Ch-Ua-Platform": `"Windows"`,
	"Sec-Fetch-Dest":     "empty",
	"Sec-Fetch-Mode":     "cors",
	"Sec-Fetch-Site":     "cross-site",
	"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 Edg/134.0.0.0",
}
