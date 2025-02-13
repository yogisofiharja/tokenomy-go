// Copyright 2025 CAMP Investment Technologies Ltd. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package camp

import "github.com/shuLhan/share/lib/math/big"

// Trade contains information about trade bid or ask, either open or
// closed.
type Trade struct {
	Price *big.Rat `json:"price,omitempty"`

	BaseAmount *big.Rat `json:"base_amount,omitempty"`
	BaseFilled *big.Rat `json:"base_filled,omitempty"`
	BaseRemain *big.Rat `json:"base_remain,omitempty"`

	CoinAmount *big.Rat `json:"coin_amount,omitempty"`
	CoinFilled *big.Rat `json:"coin_filled,omitempty"`
	CoinRemain *big.Rat `json:"coin_remain,omitempty"`

	Pair   string `json:"pair,omitempty"`
	Type   string `json:"type,omitempty"`   // Its either "sell" or "buy".
	Method string `json:"method,omitempty"` // Its either "limit" or "market".
	Status string `json:"status,omitempty"` // Status for closed trade, its either "cancelled" or "filled".

	BaseAsset string `json:"base_asset,omitempty"`
	CoinAsset string `json:"coin_asset,omitempty"`

	ID         int64 `json:"id,omitempty"`
	SubmitTime int64 `json:"submit_time,omitempty"`
	FinishTime int64 `json:"finish_time,omitempty"`
}
