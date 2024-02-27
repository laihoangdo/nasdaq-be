package utils

import "nasdaqvfs/config"

func AppendCdnUrl(cfg *config.Config, value string) string {
	if cfg == nil {

		return ""
	}
	return cfg.CDN.CDN_URL + value
}

func AppendCdnUrlString(cdn, value string) string {
	if cdn == "" || value == "" {
		return ""
	}
	return cdn + value
}
