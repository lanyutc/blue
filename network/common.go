package network

import (
	"github.com/lanyutc/blue"
)

var LOG = blue.GetLogger("net", 1)

const (
	PACKAGE_LESS = iota
	PACKAGE_FULL
	PACKAGE_ERROR
)
