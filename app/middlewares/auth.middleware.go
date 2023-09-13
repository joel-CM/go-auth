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
		// get authorization header
		authorizationHeader := c.GetHeader(authorizationHeaderKey)
		// verify that it exist
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header not privided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		// verify that the format is correct -> "Bearer access-token-string"
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("authorization header not privided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		// verify that the authorization type is correct -> Bearer
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type '%s'", authorizationType)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		// validate the token
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
