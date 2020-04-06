// Copyright 2019 Tokenomy Technologies Ltd. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

//
// Package v2 is the official Go library for Tokenomy API v2.
// This package provide client for REST and Websocket API.
//
package v2

import (
	"strconv"
	"time"
)

const (
	DefaultAddress = "https://api.tokenomy.com"

	apiMarketDepths     = "/v2/market/depths"
	apiMarketInfo       = "/v2/market/info"
	apiMarketTradesOpen = "/v2/market/trades/open"
	apiMarketPrices     = "/v2/market/prices"
	apiMarketTicker     = "/v2/market/ticker"
	apiMarketTrades     = "/v2/market/trades"
	apiMarketSummaries  = "/v2/market/summaries"

	apiUserInfo         = "/v2/user/info"
	apiUserTrades       = "/v2/user/trades"
	apiUserTradesClosed = "/v2/user/trades/closed"
	apiUserTradesOpen   = "/v2/user/trades/open"
	apiUserTrade        = "/v2/user/trade"
	apiUserTransactions = "/v2/user/transactions"

	apiTradeAsk       = "/v2/trade/ask"
	apiTradeBid       = "/v2/trade/bid"
	apiTradeCancelAsk = "/v2/trade/cancel/ask"
	apiTradeCancelBid = "/v2/trade/cancel/bid"
)

func timestamp() int64 {
	return time.Now().Unix()
}

func timestampAsString() string {
	return strconv.FormatInt(timestamp(), 10)
}
