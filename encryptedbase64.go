package encryptedbase64
import (
	"encoding/base64"
)

func Encrypt(inputString string, encryptionKey string) string {
	inputBytes := []byte(inputString)
	encryptionKeyBytes := []byte(encryptionKey)
	for characterIndex := range inputBytes {
		inputBytes[characterIndex] += encryptionKeyBytes[(characterIndex + len(encryptionKeyBytes) - 1) % len(encryptionKeyBytes)]
	}
	return base64.StdEncoding.EncodeToString(inputBytes)
}

func Decrypt(inputString string, encryptionKey string) (string, error) {
	decodedBytes, errorObject := base64.StdEncoding.DecodeString(inputString)
	if errorObject != nil {
		return "", errorObject
	} else {
		encryptionKeyBytes := []byte(encryptionKey)
		for characterIndex := range decodedBytes {
			decodedBytes[characterIndex] -= encryptionKeyBytes[(characterIndex + len(encryptionKeyBytes) - 1) % len(encryptionKeyBytes)]
		}
		return string(decodedBytes), nil
	}
}