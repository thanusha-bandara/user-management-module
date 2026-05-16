package controllers
import (
    "context"
    "net/http"
    "time"
    "crm-backend/config"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

var JWTKey = []byte("CRM_SUPER_SECRET_KEY_2026")
type LoginInput struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

func LoginUser(c *gin.Context) {
    var input LoginInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Please Enter Correct Email and Password "})

        return
    }

    var userID int
    var username string
    var passwordHash string
    var roleID int
    var status string

    query := `SELECT user_id, username, password_hash, role_id, status FROM users WHERE email = $1`

    err := config.DB.QueryRow(
		context.Background(),
		query,
		input.Email
		).Scan(&userID,
		&username,
		&passwordHash,
		&roleID,
		&status)

    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "The entered email address is not in the system. Please check and try again."})
        return
    }

    if status != "active" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Your account has been temporarily deactivated. Please contact the administrator for more information."})
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(input.Password))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "The entered password is incorrect. Please check and try again."})
		return
	
    }
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := jwt.MapClaims{
        "user_id":  userID,
        "username": username,
        "role_id":  roleID,
        "exp":      expirationTime.Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(JWTKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate the JWT token."})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message":   "Login successful!",
        "token":     tokenString,
        "user_id":   userID,
        "username":  username,
        "role_id":   roleID,
    })
}
func LogoutUser(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out from the system!"})
}