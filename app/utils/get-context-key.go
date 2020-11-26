package utils

import "github.com/gin-gonic/gin"

// GetContextKey : get key from context safely
func GetContextKey(context *gin.Context, key string, value interface{}) interface{} {
	if context.Keys != nil {
		return context.Keys[key]
	} else {
		return nil
	}
}
