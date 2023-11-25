package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// handleHTTPError handles the HTTP error by returning a JSON response with the error message.
// If the error is not nil, it returns true, otherwise false.
func HandleHTTPError(c *gin.Context, err error, httpStatus int, message string) bool {
	if err != nil {
		c.JSON(httpStatus, gin.H{"message": fmt.Sprintf("%s: %v", message, err)})
		c.Abort()
		return true
	}
	return false
}

// handleObjectError handles the object error by returning a JSON response with the error message.
// If the object is nil, it returns true, otherwise false.
func HandleObjectError(c *gin.Context, obj interface{}, httpStatus int, message string) {
	if obj == nil {
		c.JSON(httpStatus, gin.H{"message": message})
		c.Abort()
	}
}
