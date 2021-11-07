package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

//SHA256生成哈希值
func GetSHA256HashCode(message []byte) string {
	//计算哈希值，返回一个长度为32的数组
	bytes2 := sha256.Sum256(message)
	//将数组转换成切片，转换成16进制，返回字符串
	hashcode2 := hex.EncodeToString(bytes2[:])
	return strings.ToUpper(hashcode2)
}
