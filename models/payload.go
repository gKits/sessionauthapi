package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var ErrExpiredToken = errors.New("token expired")

type Payload struct {
    ID           uuid.UUID
    Username     string
    IssuedAt     time.Time
    ExpiredAt    time.Time
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
    tokenID, err := uuid.NewRandom()
    if err != nil {
        return nil, err
    }

    payload := &Payload{
        ID:         tokenID,
        Username:   username,
        IssuedAt:   time.Now(),
        ExpiredAt:  time.Now().Add(duration),
    }
    return payload, nil
}

func (payload *Payload) Valid() error {
    if time.Now().After(payload.ExpiredAt) {
        return ErrExpiredToken
    }
    return nil
}
