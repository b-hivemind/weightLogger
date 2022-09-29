package api

import (
	"fmt"
	"time"

	"bhavdeep.me/weight_logger/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	JWT_SECRET_KEY = []byte("SuperSecretKey")
)

type Claims struct {
	UUID string `json:"id"`
	jwt.StandardClaims
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"msg": message})
}

func getClaimsFromToken(c *gin.Context) (*Claims, error) {
	claims := &Claims{}
	tokenStr := c.Request.Header.Get("bearer")
	token, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return JWT_SECRET_KEY, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			respondWithError(c, 401, "Unauthorized")
			return claims, err
		}
	}
	if !token.Valid {
		respondWithError(c, 401, "Unauthorized")
		return claims, err
	}
	return claims, nil
}

func jwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("bearer")
		fmt.Println(c.Request.Header)
		if token == "" {
			respondWithError(c, 401, "Unauthorized")
			return
		}
		// Will kill the reqauest on failure
		getClaimsFromToken(c)
		c.Next()
	}
}

func getJWTKey(user db.User) (string, error) {
	expirationTime := time.Now().Add(time.Minute * 10)
	claims := &Claims{
		UUID: user.UUID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func handleLogin(c *gin.Context) {
	validator := authQuery{}
	if err := c.ShouldBindJSON(&validator); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	tempUser := db.User{
		Username: validator.Username,
		Password: validator.Password,
	}
	if userExists, err := db.FindUser(tempUser); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	} else if userExists.Password == "" {
		c.JSON(400, gin.H{"msg": "User does not exist"})
		return
	} else if userExists.Password == "-1" {
		c.JSON(401, gin.H{"msg": "Unauthorized"})
		return
	} else {
		// TODO Do JWT stuff here
		jwtToken, err := getJWTKey(userExists)
		if err != nil {
			c.JSON(500, gin.H{"err": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"profile": userExists,
			"token":   jwtToken,
		})
	}

}

func handleRegister(c *gin.Context) {
	validator := authQuery{}
	if err := c.ShouldBindJSON(&validator); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	tempUser := db.User{
		Username: validator.Username,
		Password: validator.Password,
	}
	if userExists, err := db.FindUser(tempUser); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	} else if userExists.Password != "" {
		c.JSON(400, gin.H{"msg": "User already exists"})
		return
	}
	if user, err := db.RegisterUser(tempUser); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	} else {
		jwtToken, err := getJWTKey(user)
		if err != nil {
			c.JSON(500, gin.H{"err": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"profile": user,
			"token":   jwtToken,
		})
	}
}
