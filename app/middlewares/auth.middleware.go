package middlewares

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joel-CM/go-auth/app/services"
)

var (
	authorizationHeaderKey  string = "Authorization"
	authorizationTypeBearer string = "bearer"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header not privided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("authorization header not privided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type '%s'", authorizationType)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		accessTokenString := fields[1]
		ok, err := services.VerifyToken(accessTokenString)
		if err != nil {
			log.Print(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		if !ok {
			err := errors.New("invalid access token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}
		c.Next()
	}
}

func errResponse(err error) map[string]any {
	return gin.H{"message": err.Error()}
}
