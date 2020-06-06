package binance

import (
	"context"
	"encoding/json"
)

// GetAccountService get account info
type GetPositionRiskService struct {
	c *Client
}

// Do send request
func (s *GetPositionRiskService) Do(ctx context.Context, opts ...RequestOption) (res *Account, err error) {
	r := &request{
		method:   "GET",
		endpoint: "GET /fapi/v1/positionRisk",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Account)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PositionRisk define position risk of your account

type PositionRisk struct {
	EntryPrice float64 `json:"entryPrice"`
    MarginType string `json:"marginType"`
    IsAutoAddMargin string `json:"isAutoAddMargin"`
    IsolatedMargin float64 `json:"isolatedMargin"`
    Leverage int64 `json:"leverage"`
    LiquidationPrice float64 `json:"liquidationPrice"`
    MarkPrice float64 `json:"markPrice"`
    MaxNotionalValue string `json:"maxNotionalValue"`
    PositionAmt float64 `json:"positionAmt"`
    Symbol string `json:"symbol"`
    UnRealizedProfit float64 `json:"unRealizedProfit"`
}