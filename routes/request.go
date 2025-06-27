package routes

import (
    "github.com/labstack/echo/v4"

    "keycurl/handlers"
)

func RequestRoutes() *echo.Echo {

    e.GET("/api/request/all/:email", handlers.GetAllRequests)
    e.GET("/api/request/favorites/:email", handlers.GetAllFavoriteRequests)
    e.POST("/api/request/new/:email", handlers.CreateRequest)
    e.DELETE("/api/request/delete/:email/:reqID", handlers.HideRequest)

    return e
}
