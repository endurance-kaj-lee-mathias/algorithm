package middleware

import "errors"

var MissingAPIKey = errors.New("X-API-Key header is missing")
var InvalidAPIKey = errors.New("X-API-Key is invalid")
