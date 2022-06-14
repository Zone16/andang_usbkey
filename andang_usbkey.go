package andang_usbkey

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/farmerx/gorsa"
	mRand "math/rand"
)

// Sign 签名
// privateKey string 私钥（pkcs#8格式）
// origData string 原始随机数
func Sign(privateKey, origData string) (string, error) {
	h := sha256.New()
	h.Write([]byte(origData))

	err := gorsa.RSA.SetPrivateKey(privateKey)
	if err != nil {
		return "", err
	}

	sign, err := gorsa.RSA.PriKeyENCTYPT(h.Sum(nil))
	if err != nil {
		return "", err
	}

	signBase64 := base64.StdEncoding.EncodeToString(sign)

	return signBase64, nil
}

// Verify 验签
// publicKey string 公钥（pkcs#8格式）
// sign string 私钥签名
// origData string 原始随机数
func Verify(publicKey, sign, origData string) error {
	pub := `
-----BEGIN PUBLIC KEY-----
` + publicKey + `
-----END PUBLIC KEY-----`

	return verify(pub, sign, origData)
}

// Verify2 验签(完整公钥模式)
// publicKey string 公钥（pkcs#8格式）
// sign string 私钥签名
// origData string 原始随机数
func Verify2(publicKey, sign, origData string) error {
	return verify(publicKey, sign, origData)
}

// 验签
func verify(publicKey, sign, origData string) error {
	signByte, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}

	h := sha256.New()
	h.Write([]byte(origData))

	err = gorsa.RSA.SetPublicKey(publicKey)
	if err != nil {
		return err
	}

	deSign, err := gorsa.RSA.PubKeyDECRYPT(signByte)
	if err != nil {
		return err
	}

	if hex.EncodeToString(h.Sum(nil)) != hex.EncodeToString(deSign) {
		return errors.New("验签失败")
	}

	return nil
}

// RandSeq 获取随机数
func RandSeq(n int) string {
	letters := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[mRand.Intn(len(letters))]
	}
	return string(b)
}
