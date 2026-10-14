package shared

import "errors"

var ErrResourceNotFound = errors.New("resource not found")
var ErrRequestRejected = errors.New("request rejected")
