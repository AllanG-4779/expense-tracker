package util

import "strings"

func EncryptPassword(password string) string {
	return strings.Split(password, "")[0] + "********" + strings.Split(password, "")[len(password)-1]
}