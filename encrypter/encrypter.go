package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"go-base/consoleColors"
	"io"
	"os"
)

type Encrypter struct {
	key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic(consoleColors.Colors().RedBold("Ключ для шифрования не найден"))
	}
	return &Encrypter{
		key: key,
	}
}

func (enc *Encrypter) Encript(plainString []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.key))
	enc.encrypterError(err)

	aesGSM, err := cipher.NewGCM(block)
	enc.encrypterError(err)

	nonce := make([]byte, aesGSM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	enc.encrypterError(err)
	return aesGSM.Seal(nonce, nonce, plainString, nil)
}

func (enc *Encrypter) Decript(plainString []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.key))
	enc.encrypterError(err)

	aesGSM, err := cipher.NewGCM(block)
	enc.encrypterError(err)

	nonceSize := aesGSM.NonceSize()
	nonce := plainString[:nonceSize]
	cipherText := plainString[nonceSize:]
	text, err := aesGSM.Open(nil, nonce, cipherText, nil)
	enc.encrypterError(err)

	return text
}

func (enc *Encrypter) encrypterError(err error) {
	if err != nil {
		panic(consoleColors.Colors().RedBold("Ошибка шифрования"))
	}
}
