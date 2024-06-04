package demo2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("x14nKey")

var users = map[string]string{
	"user1": "3953",
	"user2": "7980",
}

type Credentials struct {
	Name   string
	Passwd string
}

type Claims struct {
	Username string
	jwt.StandardClaims
}

func Sigin(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//密码校验
	expectedPasswd, ok := users[credentials.Name]
	if !ok || expectedPasswd != credentials.Passwd {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//设置过期时间
	expriationTime := time.Now().Add(5 * time.Minute)

	//生成payload
	claims := &Claims{
		Username: credentials.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expriationTime.Unix(),
		},
	}

	//根据header和payload生成令牌
	//根据令牌跟秘钥生成最终的token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//将token设置到cookies中
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expriationTime,
	})

	responseData := map[string]string{"token": tokenString}
	responseJson, err := json.Marshal(responseData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	// cookie, err := r.Cookie("token")

	// if err == http.ErrNoCookie {
	// 	w.WriteHeader(http.StatusUnauthorized)

	// 	if err != nil {
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	}
	// 	return
	// }

	// tknStr := cookie.Value

	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tknStr := strings.Split(tokenHeader, " ")[1]

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) { return jwtKey, nil })
	if err != nil {
		if err == jwt.ErrInvalidKey {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("welcome:%s", claims.Username)))
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tknStr := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}
