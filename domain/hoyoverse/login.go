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

const LoginTokenType = 6

var LoginHeaders = map[string]string{
	"Accept":              "application/json, text/plain, */*",
	"Accept-Language":     "en",
	"Cache-Control":       "no-cache",
	"Content-Type":        "application/json",
	"Origin":              "https://account.hoyoverse.com",
	"Pragma":              "no-cache",
	"Referer":             "https://account.hoyoverse.com/",
	"Sec-Ch-Ua":           `"Chromium";v="136", "Microsoft Edge";v="136", "Not.A/Brand";v="99"`,
	"Sec-Ch-Ua-Mobile":    "?0",
	"Sec-Ch-Ua-Platform":  `"Windows"`,
	"Sec-Fetch-Dest":      "empty",
	"Sec-Fetch-Mode":      "cors",
	"Sec-Fetch-Site":      "same-site",
	"User-Agent":          "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36 Edg/136.0.0.0",
	"X-Rpc-Age_gate":      "true",
	"X-Rpc-Aigis_v4":      "true",
	"X-Rpc-App_id":        "dnci5hp1qxa8",
	"X-Rpc-App_version":   "",
	"X-Rpc-Client_type":   "4",
	"X-Rpc-Device_fp":     "00000000000",
	HeaderXRpcDeviceID:    "",
	"X-Rpc-Device_model":  "Microsoft Edge 136.0.0.0",
	"X-Rpc-Device_name":   "Microsoft Edge",
	"X-Rpc-Device_OS":     "Windows 10 64-bit",
	"X-Rpc-Game_Biz":      "plat_os",
	"X-Rpc-Language":      "en-us",
	HeaderXRpcLifecycleID: "",
	"x-rpc-referrer":      "https://account.hoyoverse.com/passport/index.html#/",
	"x-rpc-sdk_version":   "2.39.0",
}

const (
	HeaderXRpcDeviceID    = "X-Rpc-Device_id"
	HeaderXRpcLifecycleID = "X-rpc-Lifecycle_id"
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
