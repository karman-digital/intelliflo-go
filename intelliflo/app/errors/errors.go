package intellifloerrors

import "errors"

var ErrAccessTokenNotSet = errors.New("access token not set")

var ErrAccessTokenExpired = errors.New("access token expired")
