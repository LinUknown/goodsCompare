package util

import (
	"crypto/md5"
	"fmt"
)

const TIME_TEMPLATE = "2006-01-02 15:04:05"
const NAME_SALT = "this_is_salt"
const REDIS_USER_PREFIX  = "redis_user_prefix"

func GetGoodPrex(id int64 , eid int)string{
	return fmt.Sprintf("good_key_%v_%d",id,eid)
}

func GetToken(str string) string {
	data := []byte(str+NAME_SALT)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}