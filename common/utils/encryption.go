package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"os"
)

// 生成密钥
func generateKey(password, salt []byte, keyLen int) []byte {
	return pbkdf2.Key(password, salt, 4096, keyLen, sha256.New)
}

// 加密密码
func encryptPassword(password, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	ciphertext := gcm.Seal(nonce, nonce, password, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// 解密密码
func decryptPassword(ciphertext string, key []byte) ([]byte, error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertextBytes := decodedCiphertext[:nonceSize], decodedCiphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// 加密密码

func Encryption(pass string) string {
	// 生成密钥
	password := []byte("yourSecretPassword")
	salt := []byte("yourSalt")
	key := generateKey(password, salt, 32)

	// 要加密的原始密码
	originalPassword := []byte(pass)

	// 加密密码
	encryptedPassword, err := encryptPassword(originalPassword, key)
	if err != nil {
		fmt.Println("加密错误:", err)
		os.Exit(1)
	}
	return encryptedPassword
}

// 解密密码

func Decrypt(pass string) string {
	// 生成密钥
	password := []byte("yourSecretPassword")
	salt := []byte("yourSalt")
	key := generateKey(password, salt, 32)

	// 解密密码
	decryptedPassword, err := decryptPassword(pass, key)
	if err != nil {
		fmt.Println("解密错误:", err)
		os.Exit(1)
	}
	return string(decryptedPassword)
}
