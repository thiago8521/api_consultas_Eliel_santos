package controllers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "github.com/saudemais/consultas/database"
    "github.com/saudemais/consultas/middleware"
    "github.com/saudemais/consultas/models"
)

type appointmentDTO struct {
    Datetime time.Time `json:"datetime" binding:"required"`
}

func CreateAppointment(c *gin.Context) {
    var body appointmentDTO
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": 400})
        return
    }

    if body.Datetime.Before(time.Now()) {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Consulta no passado não é permitida", "code": 400})
        return
    }

    patientID := middleware.PatientID(c)
    db := database.DB

    var exists models.Appointment
    if err := db.Where("patient_id = ? AND datetime = ?", patientID, body.Datetime).First(&exists).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Duas consultas no mesmo horário não são permitidas", "code": 400})
        return
    }

    appt := models.Appointment{
        PatientID: patientID,
        Datetime:  body.Datetime,
    }
    db.Create(&appt)

    c.JSON(http.StatusCreated, appt)
}

func ListAppointments(c *gin.Context) {
    patientID := middleware.PatientID(c)
    var appointments []models.Appointment
    db := database.DB.Where("patient_id = ?", patientID).Order("datetime").Find(&appointments)
    if db.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": db.Error.Error(), "code": 500})
        return
    }
    c.JSON(http.StatusOK, appointments)
}

func CancelAppointment(c *gin.Context) {
    patientID := middleware.PatientID(c)
    id := c.Param("id")

    db := database.DB
    var appt models.Appointment
    if err := db.Where("id = ? AND patient_id = ?", id, patientID).First(&appt).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusForbidden, gin.H{"message": "Consulta não encontrada ou acesso não autorizado", "code": 403})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "code": 500})
        return
    }

    db.Delete(&appt)
    c.Status(http.StatusNoContent)
}
