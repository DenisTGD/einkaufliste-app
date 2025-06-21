package auth

import (
    "golang.org/x/crypto/bcrypt"
    "github.com/gin-gonic/gin"
)

type RegisterInput struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

func RegisterHandler(c *gin.Context) {
    var input RegisterInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    _, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(500, gin.H{"error": "konnte Passwort nicht verschlüsseln"})
        return
    }

    c.JSON(200, gin.H{
        "message": "Registrierung erfolgreich",
        "email":   input.Email,
    })
}
