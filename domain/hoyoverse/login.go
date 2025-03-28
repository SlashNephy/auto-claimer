package hoyoverse

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

const LoginURL = "https://sg-public-api.hoyoverse.com/account/ma-passport/api/webLoginByPassword"

type LoginRequest struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	TokenType int    `json:"token_type"`
}

type LoginResponse APIResponse[struct{}]

const LoginTokenType = 4

var LoginHeaders = map[string]string{
	"Accept":             "application/json, text/plain, */*",
	"accept-language":    "en",
	"Content-Type":       "application/json",
	"Origin":             "https://account.hoyoverse.com",
	"Referer":            "https://account.hoyoverse.com/",
	"Sec-Ch-Ua":          `"Chromium";v="134", "Not:A-Brand";v="24", "Microsoft Edge";v="134"`,
	"Sec-Ch-Ua-Mobile":   "?0",
	"Sec-Ch-Ua-Platform": `"Windows"`,
	"Sec-Fetch-Dest":     "empty",
	"Sec-Fetch-Mode":     "cors",
	"Sec-Fetch-Site":     "cross-site",
	"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 Edg/134.0.0.0",
	"x-rpc-age_gate":     "true",
	"x-rpc-aigis_v4":     "true",
	"x-rpc-app_id":       "de8ohyzxreo0",
	"x-rpc-app_version":  "",
	"x-rpc-client_type":  "4",
	"x-rpc-device_fp":    "",
	"x-rpc-device_id":    "",
	"x-rpc-device_model": "Microsoft Edge 134.0.0.0",
	"x-rpc-device_name":  "Microsoft Edge",
	"x-rpc-device_os":    "Windows 10 64-bit",
	"x-rpc-game_biz":     "plat_os",
	"x-rpc-language":     "en-us",
	"x-rpc-lifecycle_id": "",
	"x-rpc-referrer":     "https://account.hoyoverse.com/passport/index.html#/",
	"x-rpc-sdk_version":  "2.38.0",
}

const (
	HeaderXRpcDeviceFp    = "x-rpc-device_fp"
	HeaderXRpcDeviceID    = "x-rpc-device_id"
	HeaderXRpcLifecycleID = "x-rpc-lifecycle_id"
)

// https://webstatic.hoyoverse.com/dora/biz/hoyoverse-account-sdk/main.js
const encryptKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4PMS2JVMwBsOIrYWRluY
wEiFZL7Aphtm9z5Eu/anzJ09nB00uhW+ScrDWFECPwpQto/GlOJYCUwVM/raQpAj
/xvcjK5tNVzzK94mhk+j9RiQ+aWHaTXmOgurhxSp3YbwlRDvOgcq5yPiTz0+kSeK
ZJcGeJ95bvJ+hJ/UMP0Zx2qB5PElZmiKvfiNqVUk8A8oxLJdBB5eCpqWV6CUqDKQ
KSQP4sM0mZvQ1Sr4UcACVcYgYnCbTZMWhJTWkrNXqI8TMomekgny3y+d6NX/cFa6
6jozFIF4HCX5aW8bp8C8vq2tFvFbleQ/Q3CU56EWWKMrOcpmFtRmC18s9biZBVR/
8QIDAQAB
-----END PUBLIC KEY-----`

func Encrypt(value string) (string, error) {
	block, _ := pem.Decode([]byte(encryptKey))
	if block == nil {
		return "", errors.New("failed to decode PEM block")
	}

	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse public key: %w", err)
	}

	publicKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return "", errors.New("not an rsa.PublicKey")
	}

	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(value))
	if err != nil {
		return "", fmt.Errorf("failed to encrypt: %w", err)
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil
}
