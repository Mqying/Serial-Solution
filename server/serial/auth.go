package serial

import (
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/silverswords/embeded-auth/crypto"
	"github.com/silverswords/embeded-auth/sign"
)

const (
	TableName      = "user"
	IdFilePath     = "config/id.pem"
	PublicFilePath = "config/public.pem"
)

var (
	errIdDecodeFailed     = errors.New("invalid id.pem")
	errPublicDecodeFailed = errors.New("invalid public.pem")
)

//Login the administrative userAuth logins
func Login() error {
	isVerified, err := Verify()
	if !isVerified {
		return err
	}

	return nil
}

//verify certificate
func Verify() (bool, error) {
	code := fmt.Sprintf("%d-%d", ID[0], ID[1])
	b := []byte(code)

	publicKeyBytes, _, err := decode(PublicFilePath)
	if publicKeyBytes == nil {
		return false, err
	}

	crypherBytes, rest, err := decode(IdFilePath)
	if crypherBytes == nil {
		return false, err
	}

	signBlock, _ := pem.Decode(rest)
	if signBlock == nil {
		return false, errIdDecodeFailed
	}

	isCrypherValid, err := crypto.Verify(b, crypherBytes, publicKeyBytes)
	if !isCrypherValid {
		return false, err
	}

	isSignValid, err := sign.VerifySign(b, signBlock.Bytes, publicKeyBytes)
	if !isSignValid {
		return false, err
	}

	return true, nil
}

//decode the .pem files
func decode(filePath string) ([]byte, []byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	pemBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, nil, err
	}

	block, rest := pem.Decode(pemBytes)
	if block == nil {
		return nil, nil, errPublicDecodeFailed
	}

	return block.Bytes, rest, nil
}
