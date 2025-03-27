package hoyoverse

import (
	"fmt"

	"github.com/SlashNephy/auto-claimer/domain/entity"
)

const HonkaiStarrailListAccountsURL = "https://api-account-os.hoyoverse.com/account/binding/api/getUserGameRolesOfRegionByCookieToken?game_biz=hkrpg_global"

type GameAccount struct {
	Game     entity.Game
	UID      string
	Nickname string

	GameBiz  string
	Language string
	Region   string
}

func (a *GameAccount) GetGame() entity.Game {
	return a.Game
}

func (a *GameAccount) GetID() string {
	return a.UID
}

func (a *GameAccount) String() string {
	return fmt.Sprintf("%s (UID:%s)", a.Nickname, a.UID)
}

var (
	_ entity.Account = new(GameAccount)
	_ fmt.Stringer   = new(GameAccount)
)

type ListAccountsResponse APIResponse[ListAccountsResponseData]

type ListAccountsResponseData struct {
	List []*struct {
		GameUID  string `json:"game_uid"`
		Nickname string `json:"nickname"`
		Region   string `json:"region"`
	} `json:"list"`
}

var ListAccountsHeaders = map[string]string{
	"Accept":             "*/*",
	"Accept-Language":    "en",
	"Origin":             "https://hsr.hoyoverse.com",
	"Referer":            "https://hsr.hoyoverse.com/",
	"Sec-Ch-Ua":          `"Chromium";v="134", "Not:A-Brand";v="24", "Microsoft Edge";v="134"`,
	"Sec-Ch-Ua-Mobile":   "?0",
	"Sec-Ch-Ua-Platform": `"Windows"`,
	"Sec-Fetch-Dest":     "empty",
	"Sec-Fetch-Mode":     "cors",
	"Sec-Fetch-Site":     "cross-site",
	"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 Edg/134.0.0.0",
	"x-rpc-language":     "en-US",
}
