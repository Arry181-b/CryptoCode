package utils

import "crypto/md5"

func Md5Hash(data []byte) ([]byte) {
	Md5Hash := md5.New()
	Md5Hash.Write(data)
	return 	Md5Hash.Sum(nil)
}
