package http_server

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		corsMiddleware := cors.New(cors.Options{
			AllowedOrigins:   viper.GetStringSlice("cors.origins"),
			AllowedMethods:   viper.GetStringSlice("cors.methods"),
			AllowedHeaders:   viper.GetStringSlice("cors.headers"),
			AllowCredentials: true,
		})
		corsMiddleware.HandlerFunc(c.Writer, c.Request)
		c.Next()
	}
}
