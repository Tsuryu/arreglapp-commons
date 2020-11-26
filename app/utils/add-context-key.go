package utils

import "github.com/gin-gonic/gin"

// AddContextKey : add key to context safely
func AddContextKey(context *gin.Context, key string, value interface{}) {
	if context.Keys != nil {
		context.Keys[key] = value
	} else {
		keys := make(map[string]interface{})
		keys[key] = value
		context.Keys = keys
	}
}
