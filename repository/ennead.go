package repository

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/SlashNephy/auto-claimer/domain/entity"
	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
	"github.com/goccy/go-json"
	"github.com/samber/do/v2"
	"github.com/samber/lo"
)

type EnneadRepository struct {
	httpClient *http.Client
}

func NewEnneadRepository(do.Injector) (*EnneadRepository, error) {
	return &EnneadRepository{
		httpClient: &http.Client{},
	}, nil
}

func (r *EnneadRepository) ListAvailableHonkaiStarrailCodes(ctx context.Context) ([]*hoyoverse.Code, error) {
	return r.listAvailableCodes(ctx, "starrail")
}

func (r *EnneadRepository) ListAvailableGenshinImpactCodes(ctx context.Context) ([]*hoyoverse.Code, error) {
	return r.listAvailableCodes(ctx, "genshin")
}

func (r *EnneadRepository) ListAvailableZenlessZoneZeroCodes(ctx context.Context) ([]*hoyoverse.Code, error) {
	return r.listAvailableCodes(ctx, "zenless")
}

func (r *EnneadRepository) listAvailableCodes(ctx context.Context, game string) ([]*hoyoverse.Code, error) {
	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://api.ennead.cc/mihoyo/%s/codes", game),
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
