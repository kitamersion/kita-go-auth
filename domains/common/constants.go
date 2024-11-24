package common

import "os"

const (
	REFRESH_TOKEN_EXPIRY = 3600 * 24 * 30 // 30 days for refresh token
	ACCESS_TOKEN_EXPIRY  = 3600 * 24      // 1 day for access token
)

var IsProduction = os.Getenv("ENV") == "production"
