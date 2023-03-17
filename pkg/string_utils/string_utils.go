package stringutils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/spf13/viper"
)

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

func HashPassword(email, password string) string {
	secret := viper.GetString("authenticate.passwordHashSecret")
	hashMethod := sha256.New()
	hashMethod.Write([]byte(secret + email + password))
	hash := hashMethod.Sum(nil)
	result := strings.ToUpper(hex.EncodeToString(hash))
	return result
}

func GetHashMD5(phone string) string {
	hashMethod := md5.New()
	hashMethod.Write([]byte(phone))
	hash := hashMethod.Sum(nil)
	return hex.EncodeToString(hash)
}

func GetToken(id string) string {
	secret := os.Getenv("PASSWORD_SECRET")
	hashMethod := sha256.New()
	hashMethod.Write([]byte(secret + id))
	hash := hashMethod.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(hash))
}

func LowerInitial(fields []string) (results []string) {
	for _, str := range fields {
		result := ""
		for j, val := range str {
			result = string(unicode.ToLower(val)) + str[j+1:]
			break
		}
		results = append(results, result)
	}
	return results
}
