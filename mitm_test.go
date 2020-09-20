package main

import (
	"encoding/json"
	"testing"
)

func TestIP(t *testing.T) {
	err := json.Unmarshal([]byte(`{
		"api": 8888,
		"secret": "114514",
		"domains": [
			"netflix.com"
		],
		"allowedips": [
			"114.114.114.114",
			"114.114.115.115"
		],
		"address": "11.4.5.14",
		"upstream": "1.1.1.1:53"
	}`), &Data)
	if err != nil {
		t.Errorf("[TEST][json.Unmarshal] %v", err)
	}

	if !checkAllowIP("114.114.114.114") {
		t.Error("[TEST][IP] Not Allow")
	}
}

func TestDomain(t *testing.T) {
	err := json.Unmarshal([]byte(`{
		"api": 8888,
		"secret": "114514",
		"domains": [
			"netflix.com"
		],
		"allowedips": [
			"114.114.114.114",
			"114.114.115.115"
		],
		"address": "11.4.5.14",
		"upstream": "1.1.1.1:53"
	}`), &Data)
	if err != nil {
		t.Errorf("[TEST][json.Unmarshal] %v", err)
	}

	if !checkAllowDomain("abcd.netflix.com") {
		t.Error("[TEST][Domain] Not Allow")
	}
}
