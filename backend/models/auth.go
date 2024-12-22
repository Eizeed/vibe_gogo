package models

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
    UserUUID        uuid.UUID
    jwt.RegisteredClaims
}

type JWT struct {};

func (t *JWT) GenToken(uuid uuid.UUID) (string, error) {
    secret := os.Getenv("JWT_SECRET");
    if secret == "" {
        return "", errors.New("No jwt secret provided");
    }

    claims := Claims {
        uuid,
        jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    };
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims);

    tokenString, err := token.SignedString([]byte(secret));
    if err != nil {
        return "", err;
    }

    return tokenString, nil;
}

func (t *JWT) DecodeToken(tokenString string) (Claims, error) {
    secret := os.Getenv("JWT_SECRET");
    if secret == "" {
        return Claims{}, errors.New("No jwt secret provided");
    }

    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })
    if err != nil {
        log.Fatal(err)
        return Claims{}, err;
    } else if claims, ok := token.Claims.(*Claims); ok {
        return *claims, nil;
    } else {
        log.Fatal("unknown claims type, cannot proceed")
        return Claims{}, err;
    }
}
















