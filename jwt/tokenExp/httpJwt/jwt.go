package httpjwt

import (
	"errors"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWT struct {
	SiginKey []byte
}

var (
	// TokenExpired     error  = errors.New("token is expired")
	// TokenNotValidYet error  = errors.New("token not active yet")
	// TokenMalformed   error  = errors.New("That 's not even a token:")
	// TokenInvalid     error  = errors.New("Could't handle this token:")
	// SignKey          string = "x14n"

	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "x14n"
)

type CustomClaims struct {
	ID       int    `json:"user ID"`
	Name     string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func JWTAUTH() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			log.Println("没有token")
			ResponseError(ctx, -1, errors.New("没有token"))
			ctx.Abort()
			return
		}
		log.Println("GET TOKEN:", token)
		j := NewJWT()

		claims, err := j.ParseToken(token)
		fmt.Println("claims", claims)
		if err != nil {
			if err == TokenExpired {
				ResponseError(ctx, -1, err)
				ctx.Abort()
				return
			}
			ResponseError(ctx, -1, err)
			ctx.Abort()
			return
		}
		ctx.Set("cliams", claims)

	}
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

func GetSignKey() string {
	return SignKey
}

func (j *JWT) CreateToken(clamis *CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clamis)
	return token.SignedString(j.SiginKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.SiginKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}

	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, TokenInvalid

}
