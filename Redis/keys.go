package redis

import (
	"strings"
)

type Key string

func (k Key) String() string {
	return string(k)
}

type SettingsKey string

const (
	prefixSettings     = "settings"
	prefixSessionState = "session_state"
	prefixExchangeRate = "exchange_rate"
)

func Settings(settingsKey SettingsKey) Key {
	return Key(prefixSettings + ":" + settingsKey)
}

func SessionState(sessionID string) Key {
	return Key(prefixSessionState + ":" + sessionID)
}

func ExchangeRate(pair string) Key {
	return Key(prefixExchangeRate + ":" + strings.ToLower(pair))
}
