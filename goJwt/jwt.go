package goJwt

import (
	"errors"
	"goTh/app/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*
*
jwt的思路
登录时创建token 并返回token，refresh_token给前端

用户访问时，带上token，通过中间件，解析验证token，验证通过，判断是否过期，验证是否在黑名单中（退出登录时存入到redis的accessToken）
如果token过期了，但是refresh_token没有过期，此时前端触发刷新token，如果refresh_token也过期了，则需要用户重新登录

用户退出时，需要accessToken存到redis的黑名单，没有redis的情况，只能等token主动过期，或者过期时间很短，用户每次访问都需要刷新token
*/
const (
	JWTSecret       = "ADjhjkdsfSFjhdskjf2432rf"
	AccessTokenTTL  = 15 * time.Minute   // 短期有效
	RefreshTokenTTL = 7 * 24 * time.Hour // 长期有效
)

type UserClaims struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

/*
*
创建token
*/
func CreateAccessToken(u *models.User) (string, error) {
	claims := UserClaims{
		UserId:   u.Id,
		UserName: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenTTL)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     //签发时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTSecret))
}

func ParseAccessToken(token string) (*UserClaims, error) {
	userClaims := &UserClaims{}
	tokenClaims, err := jwt.ParseWithClaims(token, userClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !tokenClaims.Valid {
		return nil, jwt.ErrInvalidType
	}
	return userClaims, nil
}

func GetTokenExp(tokenStr string) (*time.Time, error) {
	claims := jwt.MapClaims{}

	_, _, err := new(jwt.Parser).ParseUnverified(tokenStr, claims)
	if err != nil {
		return nil, err
	}

	if expFloat, ok := claims["exp"].(float64); ok {
		expTime := time.Unix(int64(expFloat), 0)
		return &expTime, nil
	}

	return nil, errors.New("exp claim not found")
}
