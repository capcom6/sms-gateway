package smsgateway

import "errors"

type ProcessingState string

var ErrConflictFields = errors.New("conflict fields")
