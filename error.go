package z80

import "errors"

// ErrInvalidCodes is that detect invalid codes for operation.
var ErrInvalidCodes = errors.New("invalid codes")

// ErrNotImplemented is OPCode not implemented
var ErrNotImplemented = errors.New("not implemented")

// ErrTooShortIM0 is the interruption doesn't provide enough codes in IM 0.
var ErrTooShortIM0 = errors.New("too short data for IM 0")

// ErrBreakPoint shows PC is reached to one of break points.
var ErrBreakPoint = errors.New("break point reached")
