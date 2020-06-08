// Copyright 2019 Tokenomy Technologies Ltd. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package tokenomy

import "github.com/shuLhan/share/lib/math/big"

//
// TradePrice contains the information about completed trade.
//
type TradePrice struct {
	ID         int64    `json:"id"`
	Pair       string   `json:"pair,omitempty"`
	TradeTime  int64    `json:"trade_time"`
	BaseAmount *big.Rat `json:"base_amount"`
	CoinAmount *big.Rat `json:"coin_amount"`
	Price      *big.Rat `json:"price"`
}
