package smsgateway

import "errors"

type ProcessState string

var ErrConflictFields = errors.New("conflict fields")
