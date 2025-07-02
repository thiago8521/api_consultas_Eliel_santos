package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/saudemais/consultas/controllers"
    "github.com/saudemais/consultas/middleware"
)

func Register(r *gin.Engine) {
    // Public
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    // Protected
    auth := r.Group("/")
    auth.Use(middleware.Auth())
    {
        auth.POST("/appointments", controllers.CreateAppointment)
        auth.GET("/appointments", controllers.ListAppointments)
        auth.DELETE("/appointments/:id", controllers.CancelAppointment)
    }
}
