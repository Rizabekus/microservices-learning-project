package service

import "errors"

var ErrOrderNotFound = errors.New("order not found")
var ErrForbidden = errors.New("order does not belong to user")
var ErrOrderAlreadyCancelled = errors.New("order is already cancelled")
var ErrCannotCancelPaidOrder = errors.New("cannot cancel a paid order")
