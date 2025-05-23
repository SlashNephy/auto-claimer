package repository

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/SlashNephy/auto-claimer/domain/entity"
	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
	"github.com/goccy/go-json"
	"github.com/samber/do/v2"
)

type HoYoverseRepository struct {
	httpClient *http.Client

	hoYoverseUUID string
	miHoYoUUID    string
	lifecycleID   string
}

func NewHoYoverseRepository(i do.Injector) (*HoYoverseRepository, error) {
	jar := NewSharedCookieJar(hoyoverse.DefaultCookies)

	return &HoYoverseRepository{
		httpClient: &http.Client{
			Jar: jar,
		},
	}, nil
}

type hoYoverseAvailableCodes struct {
	Active []*hoYoverseCode `json:"active"`
}

type hoYoverseCode struct {
	Code    string   `json:"code"`
	Rewards []string `json:"rewards"`
}

func (r *HoYoverseRepository) Login(ctx context.Context, email, password string, hoYoverseUUID, miHoYoUUID *string) error {
	encryptedEmail, err := hoyoverse.Encrypt(email)
	if err != nil {
		return fmt.Errorf("failed to encrypt email: %w", err)
	}

	encryptedPassword, err := hoyoverse.Encrypt(password)
	if err != nil {
		return fmt.Errorf("failed to encrypt password: %w", err)
	}

	requestBody, err := json.Marshal(&hoyoverse.LoginRequest{
		Account:   encryptedEmail,
		Password:  encryptedPassword,
		TokenType: hoyoverse.LoginTokenType,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, hoyoverse.LoginURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range hoyoverse.LoginHeaders {
		request.Header.Set(k, v)
	}

	if hoYoverseUUID != nil {
		r.hoYoverseUUID = *hoYoverseUUID
	} else {
		r.hoYoverseUUID = hoyoverse.GenerateUUID()
	}
	if miHoYoUUID != nil {
		r.miHoYoUUID = *miHoYoUUID
	} else {
		r.miHoYoUUID = hoyoverse.GenerateUUID()
	}
	r.lifecycleID = hoyoverse.GenerateUUID()

	r.httpClient.Jar.SetCookies(nil, []*http.Cookie{
		{
			Name:  hoyoverse.CookieLifecycleID,
			Value: "{%22value%22:%22" + r.lifecycleID + "%22}",
		},
		{
			Name:  hoyoverse.CookieHoYoverseUUID,
			Value: r.hoYoverseUUID,
		},
		{
			Name:  hoyoverse.CookieMiHoYoUUID,
			Value: r.miHoYoUUID,
		},
	})
	request.Header.Set(hoyoverse.HeaderXRpcDeviceID, r.hoYoverseUUID)
	request.Header.Set(hoyoverse.HeaderXRpcDeviceID, r.hoYoverseUUID)
	request.Header.Set(hoyoverse.HeaderXRpcLifecycleID, r.lifecycleID)

	response, err := r.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body")
	}

	var result hoyoverse.LoginResponse
	if err = json.Unmarshal(responseBody, &result); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return hoyoverse.MapAPIError(result.Code, result.Message)
}

func (r *HoYoverseRepository) listGameAccounts(ctx context.Context, gameBiz string, game entity.Game) ([]*hoyoverse.GameAccount, error) {
	u, err := url.Parse(hoyoverse.ListAccountsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	query := u.Query()
	query.Set("game_biz", gameBiz)
	u.RawQuery = query.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range hoyoverse.ListAccountsHeaders {
		request.Header.Set(k, v)
	}

	response, err := r.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %s", response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result hoyoverse.ListAccountsResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if err = hoyoverse.MapAPIError(result.Code, result.Message); err != nil {
		return nil, err
	}

	accounts := make([]*hoyoverse.GameAccount, 0, len(result.Data.List))
	for _, account := range result.Data.List {
		accounts = append(accounts, &hoyoverse.GameAccount{
			Game:     game,
			UID:      account.GameUID,
			Nickname: account.Nickname,
			Language: "ja", // TODO
			Region:   account.Region,
			GameBiz:  gameBiz,
		})
	}

	return accounts, nil
}
