package hoyoverse

import (
	"crypto/rand"
	"math/big"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

var DefaultCookies = []*http.Cookie{
	{
		Name:  "HYV_LOGIN_PLATFORM_LOAD_TIMEOUT",
		Value: "{}",
	},
	{
		Name:  "HYV_LOGIN_PLATFORM_OPTIONAL_AGREEMENT",
		Value: "{%22content%22:[]}",
	},
	{
		Name:  "HYV_LOGIN_PLATFORM_TRACKING_MAP",
		Value: "{}",
	},
	{
		Name:  "mi18nLang",
		Value: "en-us",
	},
}

const (
	CookieDeviceFp      = "DEVICEFP"
	CookieLifecycleID   = "HYV_LOGIN_PLATFORM_LIFECYCLE_ID"
	CookieHoYoverseUUID = "_HYVUUID"
	CookieMiHoYoUUID    = "_MHYUUID"
)

func GenerateDeviceFp() string {
	return GenerateRandomNumberString(11)
}

func GenerateRandomNumberString(length int) string {
	result := ""
	for range length {
		i, _ := rand.Int(rand.Reader, big.NewInt(10)) // [0, 10)
		result += strconv.FormatInt(i.Int64(), 10)
	}

	return result
}

func GenerateUUID() string {
	return uuid.New().String()
}
