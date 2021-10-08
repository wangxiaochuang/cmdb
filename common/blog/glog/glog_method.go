package glog

import (
	"strconv"
)

func GetV() Level {
	return logging.verbosity
}

func SetV(level Level) {
	logging.verbosity.Set(strconv.Itoa(int(level)))
}
