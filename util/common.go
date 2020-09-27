package util

import "github.com/schigh/str"

func EncryptPassword(password string) string {
	return str.SHA256(password)
}
