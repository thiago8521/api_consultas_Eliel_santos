package controllers

import (
    "net/http"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"

    "github.com/saudemais/consultas/database"
    "github.com/saudemais/consultas/models"
)

type registerDTO struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

func Register(c *gin.Context) {
    var body registerDTO
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": 400})
        return
    }

    db := database.DB
    var existing models.Patient
    if err := db.Where("email = ?", body.Email).First(&existing).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{"message": "Paciente já cadastrado", "code": 409})
        return
    }

    hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 12)

    patient := models.Patient{
        Name:         body.Name,
        Email:        body.Email,
        PasswordHash: string(hash),
    }
    db.Create(&patient)

    c.JSON(http.StatusCreated, gin.H{"message": "Paciente cadastrado com sucesso"})
}

type loginDTO struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
    var body loginDTO
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": 400})
        return
    }

    db := database.DB
    var patient models.Patient
    if err := db.Where("email = ?", body.Email).First(&patient).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Credenciais inválidas", "code": 401})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "code": 500})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(patient.PasswordHash), []byte(body.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Credenciais inválidas", "code": 401})
        return
    }

    secret := []byte(os.Getenv("JWT_SECRET"))
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":  patient.ID,
        "exp": time.Now().Add(24 * time.Hour).Unix(),
    })
    s, _ := token.SignedString(secret)

    c.JSON(http.StatusOK, gin.H{"token": s})
}
