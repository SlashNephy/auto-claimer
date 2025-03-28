package repository

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/SlashNephy/auto-claimer/domain/entity"
	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
	"github.com/goccy/go-json"
)

func (r *HoYoverseRepository) ListGenshinImpactGameAccounts(ctx context.Context) ([]*hoyoverse.GameAccount, error) {
	return r.listGameAccounts(ctx, "nap_global", entity.GameGenshinImpact)
}

func (r *HoYoverseRepository) RedeemGenshinImpactCode(ctx context.Context, account *hoyoverse.GameAccount, code *hoyoverse.Code) error {
	u, err := url.Parse(hoyoverse.GenshinImpactRedemptionURL)
	if err != nil {
		return fmt.Errorf("failed to parse url: %w", err)
	}

	query := u.Query()
	query.Set("uid", account.UID)
	query.Set("region", account.Region)
	query.Set("lang", account.Language)
	query.Set("cdkey", code.GetCode())
	query.Set("game_biz", account.GameBiz)
	query.Set("sLangKey", "en-us")
	u.RawQuery = query.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range hoyoverse.RedemptionHeaders {
		request.Header.Set(k, v)
	}

	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Origin", hoyoverse.GenshinImpactRedemptionOrigin)
	request.Header.Set("Referer", hoyoverse.GenshinImpactRedemptionOrigin+"/")

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
