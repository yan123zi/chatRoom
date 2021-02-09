package middlewares

import (
	"chatRoom/server/common"
	"chatRoom/server/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/sirupsen/logrus"
)

const (
	secretKey = "iowefiuzxfj/><?.,"
)

type Claims struct {
	User *model.User
	jwt.StandardClaims
}

func Auth() iris.Handler {
	return func(c *context.Context) {
		tokenStr := c.GetHeader("token-jwt")
		token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			logrus.Errorf("Auth err:%s", err.Error())
			c.StopWithJSON(401, common.NewError(-1, "unauthorized"))
			return
		}
		//判断token是否有效
		_, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			logrus.Error("Auth err:%s", err.Error())
			c.StopWithJSON(401, common.NewError(-1, "unauthorized"))
			return
		}
		c.Next()
	}
}

func MakeToken(user *model.User) (string, error) {
	claims := &Claims{
		User: user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)
	tokenStr, err := token.SignedString(secretKey)
	return tokenStr, err
}
