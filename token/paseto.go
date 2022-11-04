package token

import (
	"os"
	"strconv"
	"time"

	"github.com/gKits/sessionauthapi/models"
    "github.com/o1egl/paseto"
)

func CreatePasetoToken(username string) (string, error) {
    duration, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION"))
    if err != nil {
        return "", err
    }
    payload, err := models.NewPayload(username, time.Duration(duration))
    if err != nil {
        return "", err
    }

    p := paseto.NewV2()
    return p.Encrypt([]byte(os.Getenv("TOKEN_SECRET")), payload, nil)
}

func ValidatePasetoToken(tokenString string) (*models.Payload, error) {
    payload := &models.Payload{}

    p := paseto.NewV2()
    err := p.Decrypt(tokenString, []byte(os.Getenv("TOKEN_SECRET")), payload, nil)
    if err != nil {
        return nil, err
    }

    err = payload.Valid()
    if err != nil {
        return nil, err
    }

    return payload, nil
}
