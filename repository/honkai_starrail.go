package repository

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/SlashNephy/auto-claimer/domain/entity"
	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
	"github.com/SlashNephy/auto-claimer/query"
	"github.com/SlashNephy/auto-claimer/workflow"
	"github.com/goccy/go-json"
	"github.com/samber/do/v2"
	"github.com/samber/lo"
)

type HonkaiStarrailRepository struct {
	httpClient *http.Client

	query.HonkaiStarrailCodeStore
	workflow.RedeemHonkaiStarrailCodeStore
}

func NewHonkaiStarrailRepository(i do.Injector) (*HonkaiStarrailRepository, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	return &HonkaiStarrailRepository{
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

func (r *HonkaiStarrailRepository) ListAvailableCodes(ctx context.Context) ([]*hoyoverse.Code, error) {
	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://api.ennead.cc/mihoyo/starrail/codes",
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
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

	var codes hoYoverseAvailableCodes
	if err = json.Unmarshal(body, &codes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return lo.Map(codes.Active, func(code *hoYoverseCode, _ int) *hoyoverse.Code {
		return &hoyoverse.Code{
			Game:    entity.GameHonkaiStarrail,
			Code:    code.Code,
			Rewards: code.Rewards,
		}
	}), nil
}

func (r *HonkaiStarrailRepository) Login(ctx context.Context, email, password string) error {
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

func (r *HonkaiStarrailRepository) ListGameAccounts(ctx context.Context) ([]*hoyoverse.GameAccount, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, hoyoverse.HonkaiStarrailListAccountsURL, nil)
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
			Game:     entity.GameHonkaiStarrail,
			UID:      account.GameUID,
			Nickname: account.Nickname,
			Language: "ja", // TODO
			Region:   account.Region,
			GameBiz:  "hkrpg_global",
		})
	}

	return accounts, nil
}

func (r *HonkaiStarrailRepository) RedeemCode(ctx context.Context, account *hoyoverse.GameAccount, code *hoyoverse.Code) error {
	requestBody, err := json.Marshal(&hoyoverse.RedemptionRequest{
		CDKey:    code.GetCode(),
		GameBiz:  account.GameBiz,
		Lang:     account.Language,
		Platform: "4",
		Region:   account.Region,
		Time:     time.Now().UnixMilli(),
		UID:      account.UID,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, hoyoverse.HonkaiStarrailRedemptionURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range hoyoverse.RedemptionHeaders {
		request.Header.Set(k, v)
	}

	response, err := r.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var result hoyoverse.RedemptionResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return hoyoverse.MapAPIError(result.Code, result.Message)
}
