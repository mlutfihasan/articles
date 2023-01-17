package auth

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/exception"
	"gitlab.com/VNEU/mf-micro-service-cogs-data/helper"
)

type CreateAuthFunc func(userID uint, tokenDetails *TokenDetails)

type AccessDetails struct {
	UserID uint
	Role   string
}

type TokenDetails struct {
	AccessToken         string
	RefreshToken        string
	AccessUUID          string
	RefreshUUID         string
	UserAgent           string
	RemoteAddress       string
	AtExpired           int64
	RefreshTokenExpired int64
}

func Auth(next func(c *gin.Context, auth *AccessDetails), roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check JWT Token
		tokenAuth, err := ExtractTokenMetadata(ExtractToken(c.Request))
		if err != nil {
			helper.PanicIfError(exception.ErrUnauthorized)
		}

		// Check Permission User
		// if !helper.Contains(roles, tokenAuth.Role) {
		// 	helper.PanicIfError(exception.ErrPermissionDenied)
		// }

		next(c, tokenAuth)
	}
}

func CreateToken(level string, userAgent, remoteAddress *string, createAuthFunc CreateAuthFunc) (*TokenDetails, error) {
	td := &TokenDetails{
		AtExpired:           time.Now().Add(time.Hour * 24 * 3).Unix(),
		AccessUUID:          uuid.NewV4().String(),
		RefreshTokenExpired: time.Now().Add(time.Hour * 24 * 7).Unix(),
		RefreshUUID:         uuid.NewV4().String(),
		UserAgent:           *userAgent,
		RemoteAddress:       *remoteAddress,
	}

	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["exp"] = td.AtExpired
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["exp"] = td.RefreshTokenExpired
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func ExtractTokenMetadata(stringToken string) (*AccessDetails, error) {
	token, err := VerifyToken(stringToken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			UserID: uint(userID),
			Role:   claims["role"].(string),
		}, nil
	}
	return nil, err
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	// if err != nil {
	// 	return nil, err
	// }
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
