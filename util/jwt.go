package util

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("keyboard cat")
var exp_login = time.Hour * 24

type JwtMapClaims struct {
	jwt.RegisteredClaims
	UserId uint
	Exp    *jwt.NumericDate
}

func CreateToken(userId uint) (token string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtMapClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   fmt.Sprint(userId),
		},
		Exp: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})

	token, err = jwtToken.SignedString(secret)
	return
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtMapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// jwtMapClaims := token.Claims.(*JwtMapClaims)
	return token, nil
}

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		accessToken := strings.Split(authorization, "Bearer ")[1]
		token, err := VerifyToken(accessToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		jwtMapClaims := token.Claims.(*JwtMapClaims)
		userId := jwtMapClaims.UserId
		// username := jwtMapClaims.Username
		log.Println("AuthorizationMiddleware:", userId)

		ctx := context.WithValue(r.Context(), "user_id", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetValueFromContext(ctx context.Context) (userId uint, err error) {
	userId, ok := ctx.Value("user_id").(uint)
	if !ok {
		return 0, fmt.Errorf("Unauthorized")
	}
	return userId, nil
}

/*
func GetAuthorization(w http.ResponseWriter, r *http.Request) (userId uint, username string, err error) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		err = fmt.Errorf("Unauthorized")
		return
	}

	accessToken := strings.Split(authorization, "Bearer ")[1]
	token, err := VerifyToken(accessToken)
	if err != nil {
		return
	}

	jwtMapClaims := token.Claims.(*JwtMapClaims)
	userId = jwtMapClaims.UserId

	return

	// fmt.Printf("%+v\n", jwtMapClaims.Username)

	// w.WriteHeader(http.StatusUnauthorized)
	// w.Write([]byte("Unauthorized"))

	// var user models.User
	// if err := db.First(&user, "email = ?", jwtMapClaims.Username).Error; err != nil {
	// 	return
	// }

	// var userProfile = UserProfile{
	// 	ID:       user.ID,
	// 	Username: user.Email,
	// }
	// json.NewEncoder(w).Encode(userProfile)
}
*/
