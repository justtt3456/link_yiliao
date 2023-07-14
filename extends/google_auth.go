package extends

import (
	"bytes"
	"china-russia/common"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type GoogleAuthInterface interface {
	QrcodeUrl(account, secret string) string
	Secret() string
	Code(secret string) (string, error)
	VerifyCode(secret, code string) (bool, error)
}
type GoogleAuth struct {
	key string
}

func NewGoogleAuth() GoogleAuthInterface {
	return &GoogleAuth{
		key: "C5JDNW4HV3AIIKSM",
	}
}
func (this *GoogleAuth) un() int64 {
	return time.Now().UnixNano() / 1000 / 30
}
func (this *GoogleAuth) hmacSha1(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	if total := len(data); total > 0 {
		h.Write(data)
	}
	return h.Sum(nil)
}

func (this *GoogleAuth) base32encode(src []byte) string {
	return base32.StdEncoding.EncodeToString(src)
}

func (this *GoogleAuth) base32decode(s string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(s)
}

func (this *GoogleAuth) toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func (this *GoogleAuth) toUint32(bts []byte) uint32 {
	return (uint32(bts[0]) << 24) + (uint32(bts[1]) << 16) +
		(uint32(bts[2]) << 8) + uint32(bts[3])
}

func (this *GoogleAuth) oneTimePassword(key []byte, data []byte) uint32 {
	hash := this.hmacSha1(key, data)
	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := this.toUint32(hashParts)
	return number % 1000000
}

// 获取秘钥
func (this *GoogleAuth) Secret() string {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, this.un())
	return this.encode(strings.ToUpper(this.base32encode(this.hmacSha1(buf.Bytes(), nil))))
}

func (this *GoogleAuth) encode(s string) string {
	s = s + this.key
	res := base64.StdEncoding.EncodeToString([]byte(s))
	newStr := res[:5] + common.RandUpperString(1) + res[5:13] + common.RandUpperString(1) + res[13:]
	swap := []rune(newStr)
	for i, j := 0, len(swap)-1; i < j; i, j = i+1, j-1 {
		swap[i], swap[j] = swap[j], swap[i]
	}
	return string(swap)
}
func (this *GoogleAuth) decode(s string) string {
	swap := []rune(s)
	for i, j := 0, len(swap)-1; i < j; i, j = i+1, j-1 {
		swap[i], swap[j] = swap[j], swap[i]
	}
	stringSwap := string(swap)
	newStr := stringSwap[:5] + stringSwap[6:14] + stringSwap[15:]
	decodeByte, _ := base64.StdEncoding.DecodeString(newStr)
	return string(decodeByte)[:32]
}

// 获取动态码
func (this *GoogleAuth) Code(secret string) (string, error) {
	secretUpper := this.decode(secret)
	secretKey, err := this.base32decode(secretUpper)
	if err != nil {
		return "", err
	}
	number := this.oneTimePassword(secretKey, this.toBytes(time.Now().Unix()/30))
	return fmt.Sprintf("%06d", number), nil
}

// 获取动态码二维码内容
func (this *GoogleAuth) getQrcode(user, secret string) string {
	return fmt.Sprintf("otpauth://totp/%s?secret=%s", user, secret)
}

// 获取动态码二维码图片地址,这里是第三方二维码api
func (this *GoogleAuth) QrcodeUrl(user, secret string) string {
	secret = this.decode(secret)
	qrcode := this.getQrcode(user, secret)
	width := "200"
	height := "200"
	data := url.Values{}
	data.Set("data", qrcode)
	return "https://api.qrserver.com/v1/create-qr-code/?" + data.Encode() + "&size=" + width + "x" + height + "&ecc=M"
}

// 验证动态码
func (this *GoogleAuth) VerifyCode(secret, code string) (bool, error) {
	_code, err := this.Code(secret)
	if err != nil {
		return false, err
	}
	return _code == code, nil
}
