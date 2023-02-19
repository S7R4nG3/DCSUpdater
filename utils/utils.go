package utils

import (
	"github.com/sirupsen/logrus"
)

func Check(e error) {
	if e != nil {
		logrus.Error(e)
	}
}
