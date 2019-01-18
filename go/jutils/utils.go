package jutils

import (
	"os"
)

func GetHostName() string {
	hn, _ := os.Hostname()
	return hn

}
