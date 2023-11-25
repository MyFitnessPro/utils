package utils

import (
	"fmt"
	"net/http"

	goFirebase "github.com/MyFitnessPro/firebase"
	"github.com/gin-gonic/gin"
)

// extractAndValidateQueryParams extracts and validates the query parameters.
// If uid or role are empty, it returns an error.
func extractAndValidateQueryParams(c *gin.Context) (string, string, error) {
	uid := c.Query("uid")
	role := c.Query("role")

	if uid == "" || role == "" {
		return "", "", fmt.Errorf("Missing uid or role")
	}
	return uid, role, nil
}

// handleHTTPError handles the HTTP error by returning a JSON response with the error message.
// If the error is not nil, it returns true, otherwise false.
func handleHTTPError(c *gin.Context, err error, httpStatus int, message string) bool {
	if err != nil {
		c.JSON(httpStatus, gin.H{"message": fmt.Sprintf("%s: %v", message, err)})
		c.Abort()
		return true
	}
	return false
}

// handleObjectError handles the object error by returning a JSON response with the error message.
// If the object is nil, it returns true, otherwise false.
func handleObjectError(c *gin.Context, obj interface{}, httpStatus int, message string) {
	if obj == nil {
		c.JSON(httpStatus, gin.H{"message": message})
		c.Abort()
	}
}

// processRequest processes the HTTP request by extracting and validating query parameters,
// binding the request body to a map, and returning the extracted data.
// If any error occurs during the process, it returns the error.
func processRequest(c *gin.Context, client *goFirebase.FirebaseClient) (string, string, map[string]interface{}, error) {
	uid, role, err := extractAndValidateQueryParams(c)
	if handleHTTPError(c, err, http.StatusBadRequest, "Invalid query parameters") {
		return "", "", nil, err
	}

	var userData map[string]interface{}
	if c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPost {
		if err := c.BindJSON(&userData); err != nil {
			handleHTTPError(c, err, http.StatusBadRequest, "Invalid request body")
			return "", "", nil, err
		}
	}

	return uid, role, userData, nil
}
