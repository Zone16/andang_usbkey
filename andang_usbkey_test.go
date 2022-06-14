package andang_usbkey

import (
	"fmt"
	"testing"
)

var pri = `-----BEGIN PRIVATE KEY-----
MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBAL2s3eGmWByeefcd
yVA0naKsJ6eX1xwWBVlxYpMJu8WWmMaVDc96vZhdFGM90xMh5448ah1QtJsYs25N
hDqz/RqV8cTgxnz7Zexn/JHTULoKFpog87UinuTUZNj8npAs1ShcbN4jBv4lffPp
Q0mJ+XAjv5MNEsi5usHowqEL2YGpAgMBAAECgYAx1DUiL7IcKqH+Aow9juUOtE1i
oOg5D6vuGHR+DmOPAAY4vTW4mRUv7twJSlemQhTz7/kspGQeDrossx6W0WwfWOco
H1w3pTmTEJUG63dfsf62znPYX3E96JlEjh54Uc6bF/dahpLqLRcxe9d5YRi1UiTe
ImfNG9iwZxyh73bPUQJBAOlJr4/jOGMI3kdy9vtzovbz1pRRDqrDKXo6VpMMTdYz
m0vY+vJx1qjloi/1oJSSKGwCiFkFjS5UiYAXiN95NnUCQQDQJDWiL8EOpLCZPi6D
2WpC7wH7XsLKFvrM+2btveDaW1sYqOeqmy4PCNbnuZ7mDqtfQ9NMGnGKCMtoMSch
UT/lAkEApiG6a76FdklnefxOFK253SGyqva6ejL3g7qt0pRNjgA8VJxVwXf+RVMa
2AqU65jWPmzjSnogm2DKdrTL3VFFzQJBAMXI4r5PqqCDltzH3eOkgflArR10mpz2
4TRP4SshN73G+fWg6yOPemEHAVAJbxkl72sDRJIYvrNwT/meW5SpZJECQQDPzd86
Wcdtejs8rECL9ckv+LZoMCwg5qAJVJyHmIt7JcSyz1UezOVyK+kKdF3HAvUJSYm6
SBeM5NAWr4RVwIQz
-----END PRIVATE KEY-----
`
var pub = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC9rN3hplgcnnn3HclQNJ2irCen
l9ccFgVZcWKTCbvFlpjGlQ3Per2YXRRjPdMTIeeOPGodULSbGLNuTYQ6s/0alfHE
4MZ8+2XsZ/yR01C6ChaaIPO1Ip7k1GTY/J6QLNUoXGzeIwb+JX3z6UNJiflwI7+T
DRLIubrB6MKhC9mBqQIDAQAB`

func TestSign(t *testing.T) {
	res, err := Sign(pri, "123")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(res))
}

func TestVerify(t *testing.T) {
	err := Verify(pub, "xnvF0o6qQIbMteSfLZ/2a+aQ4bYL34akVtp+Snk+dIhfR7kufxUyWSDjV9DgVtHl+ZTMUlPhxcITeqiJlUIiCsTXVHUXWLR+RHkHaeV9Ue99owRDeFtXePebh1nyWJpTHgfJvbHjf0AIVPcE85iIUVxyJzoVuxl37qkA54xA9XE=", "123")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestVerifyAndVerify(t *testing.T) {
	sign := "123"
	res, err := Sign(pri, sign)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = Verify(pub, res, sign)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ok")
}
