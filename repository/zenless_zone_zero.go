package repository

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/SlashNephy/auto-claimer/domain/entity"
	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
	"github.com/goccy/go-json"
)

func (r *HoYoverseRepository) ListZenlessZoneZeroGameAccounts(ctx context.Context) ([]*hoyoverse.GameAccount, error) {
	return r.listGameAccounts(ctx, "nap_global", entity.GameZenlessZoneZero)
}

func (r *HoYoverseRepository) RedeemZenlessZoneZeroCode(ctx context.Context, account *hoyoverse.GameAccount, code *hoyoverse.Code) error {
	requestBody, err := json.Marshal(&hoyoverse.RedemptionRequest{
		CDKey:      code.GetCode(),
		DeviceUUID: r.miHoYoUUID,
		GameBiz:    account.GameBiz,
		Lang:       account.Language,
		Platform:   "4",
		Region:     account.Region,
		Time:       time.Now().UnixMilli(),
		UID:        account.UID,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, hoyoverse.ZenlessZoneZeroRedemptionURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range hoyoverse.RedemptionHeaders {
		request.Header.Set(k, v)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Origin", hoyoverse.ZenlessZoneZeroRedemptionOrigin)
	request.Header.Set("Referer", hoyoverse.ZenlessZoneZeroRedemptionOrigin+"/")

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
